package templater

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func goFunction(function *parser.Function) string {
	var output = fmt.Sprintf("%v{\n%v\n}", goFunctionHeader(function), goFunctionBody(function))
	if functionIsSupported(parser.ClassMap[function.Class()], function) {
		if UseStub() {
			if function.SignalMode != parser.CALLBACK {
				return output
			}
		} else {
			return output
		}
	}
	return ""
}

func goFunctionHeader(function *parser.Function) string {
	return fmt.Sprintf("func %v %v(%v)%v",
		func() string {
			if function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK {
				return ""
			}
			return fmt.Sprintf("(ptr *%v)", function.Class())
		}(),

		converter.GoHeaderName(function),

		converter.GoHeaderInput(function),

		converter.GoHeaderOutput(function),
	)
}

func goFunctionBody(function *parser.Function) string {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if strings.ToLower(os.Getenv("QT_DEBUG")) == "true" {
		fmt.Fprintf(bb, "defer qt.Recover(\"\t%v%v%v(%v) %v\")\n",
			function.Class(),

			func() string {
				return strings.Repeat(" ", 45-len(function.Class()))
			}(),

			converter.GoHeaderName(function), converter.GoHeaderInput(function), converter.GoHeaderOutput(function))

		fmt.Fprintf(bb, "qt.Debug(\"\t%v%v%v(%v) %v\")\n",

			function.Class(),

			func() string {
				return strings.Repeat(" ", 45-len(function.Class()))
			}(),

			converter.GoHeaderName(function), converter.GoHeaderInput(function), converter.GoHeaderOutput(function))
	}

	if parser.ClassMap[function.Class()].Stub {
		if converter.GoHeaderOutput(function) != "" {
			return fmt.Sprintf("\nreturn %v", converter.GoOutputParametersFromCFailed(function))
		}
		return ""
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK) {
		fmt.Fprintf(bb, "if ptr.Pointer() != nil {\n")
	}

	for _, parameter := range function.Parameters {
		if parameter.Value == "..." || (parameter.Value == "T" && parser.ClassMap[function.Class()].Module == "QtAndroidExtras" && function.TemplateModeJNI == "") {
			for i := 0; i < 10; i++ {
				if parameter.Value == "T" {
					fmt.Fprintf(bb, "var p%v, d%v = assertion(%v, %v)\n", i, i, i, parameter.Name)
				} else {
					fmt.Fprintf(bb, "var p%v, d%v = assertion(%v, v...)\n", i, i, i)
				}
				fmt.Fprintf(bb, "if d%v != nil {\ndefer d%v()\n}\n", i, i)

				if parameter.Value == "T" {
					break
				}
			}
		}
	}

	if ((function.Meta == parser.PLAIN && function.SignalMode == "") ||
		(function.Meta == parser.SLOT && function.SignalMode == "") ||
		(function.Meta == parser.CONSTRUCTOR || function.Meta == parser.DESTRUCTOR) && function.SignalMode == "") ||
		(function.Meta == parser.SIGNAL && (function.SignalMode == "" || function.SignalMode == parser.CONNECT || function.SignalMode == parser.DISCONNECT)) ||
		(function.Meta == parser.GETTER || function.Meta == parser.SETTER) {

		//TODO:
		if functionIsSupported(parser.ClassMap[function.Class()], function) {
			cppFunction(function)
			if functionIsSupported(parser.ClassMap[function.Class()], function) {

				for _, alloc := range converter.GoInputParametersForCAlloc(function) {
					fmt.Fprint(bb, alloc)
				}

				var body = converter.GoOutputParametersFromC(function, fmt.Sprintf("C.%v(%v)", converter.CppHeaderName(function), converter.GoInputParametersForC(function)))
				fmt.Fprint(bb, func() string {
					if converter.GoHeaderOutput(function) != "" {
						switch {
						case function.NeedsFinalizer && classIsSupported(parser.ClassMap[converter.CleanValue(function.Output)]) || function.Meta == parser.CONSTRUCTOR && !(classNeedsCallbackFunctions(parser.ClassMap[function.Name]) || parser.ClassMap[function.Name].IsQObjectSubClass()):
							{
								return fmt.Sprintf("var tmpValue = %v\nruntime.SetFinalizer(tmpValue, (%v).Destroy%v)\nreturn tmpValue%v",

									body,

									func() string {
										if function.TemplateModeJNI != "" {
											return fmt.Sprintf("*%v", converter.CleanValue(function.Output))
										}
										return converter.GoHeaderOutput(function)
									}(),

									func() string {
										if function.Meta == parser.CONSTRUCTOR {
											return function.Name
										}
										return converter.CleanValue(function.Output)
									}(),

									func() string {
										if function.TemplateModeJNI == "String" {
											return ".ToString()"
										}
										return ""
									}())
							}

						case parser.ClassMap[converter.CleanValue(function.Output)].IsQObjectSubClass() && converter.GoHeaderOutput(function) != "unsafe.Pointer" || function.Meta == parser.CONSTRUCTOR && parser.ClassMap[converter.CleanValue(function.Name)].IsQObjectSubClass():
							{
								return fmt.Sprintf("var tmpValue = %v\nif !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), \"QObject::destroyed\") {\ntmpValue.ConnectDestroyed(func(%v){ tmpValue.SetPointer(nil) })\n}\nreturn tmpValue",

									body,

									func() string {
										if parser.ClassMap[function.Class()].Module == "QtCore" {
											return "*QObject"
										}
										return "*core.QObject"
									}())
							}

						default:
							{
								return fmt.Sprintf("return %v", body)
							}
						}
					}
					return body
				}())

			}
			function.Access = "public"
		}

	}

	switch function.SignalMode {
	case parser.CALLBACK:
		{
			fmt.Fprintf(bb, "%vif signal := qt.GetSignal(fmt.Sprint(ptr), \"%v::%v%v\"); signal != nil {\n",
				func() string {
					if function.Meta != parser.SLOT {
						return "\n"
					}
					return ""
				}(), function.Class(), function.Name, function.OverloadNumber)

			if converter.GoHeaderOutput(function) == "" {
				fmt.Fprintf(bb, "signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function))
			} else {
				fmt.Fprintf(bb, "return %v", converter.GoInput(fmt.Sprintf("signal.(%v)(%v)", converter.GoHeaderInputSignalFunction(function), converter.GoInputParametersForCallback(function)), function.Output, function))
			}

			fmt.Fprintf(bb, "\n}%v\n",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if function.Virtual == parser.IMPURE {
							return "else{"
						}
					}
					return ""
				}(),
			)

			if converter.GoHeaderOutput(function) == "" {
				if function.Virtual == parser.IMPURE && functionIsSupportedDefault(function) {
					fmt.Fprintf(bb, "New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.Class()), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function))
				}
			} else {
				if function.Virtual == parser.IMPURE && functionIsSupportedDefault(function) {
					fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(fmt.Sprintf("New%vFromPointer(ptr).%v%vDefault(%v)", strings.Title(function.Class()), strings.Replace(strings.Title(function.Name), parser.TILDE, "Destroy", -1), function.OverloadNumber, converter.GoInputParametersForCallback(function)), function.Output, function))
				} else {
					fmt.Fprintf(bb, "\nreturn %v", converter.GoInput(converter.GoOutputParametersFromCFailed(function), function.Output, function))
				}
			}

			fmt.Fprintf(bb, "%v",
				func() string {
					if converter.GoHeaderOutput(function) == "" {
						if function.Virtual == parser.IMPURE {
							return "\n}"
						}
					}
					return ""
				}(),
			)

		}

	case parser.CONNECT, parser.DISCONNECT:
		{
			fmt.Fprintf(bb, "\nqt.%vSignal(fmt.Sprint(ptr.Pointer()), \"%v::%v%v\"%v)",
				function.SignalMode,

				function.Class(),

				function.Name,

				function.OverloadNumber,

				func() string {
					if function.SignalMode == parser.CONNECT {
						return ", f"
					}
					return ""
				}(),
			)
		}
	}

	if (function.Meta == parser.DESTRUCTOR || strings.Contains(function.Name, "deleteLater") || strings.HasPrefix(function.Name, parser.TILDE)) && function.SignalMode == "" {
		if classNeedsCallbackFunctions(parser.ClassMap[function.Class()]) || parser.ClassMap[function.Class()].IsQObjectSubClass() {
			fmt.Fprint(bb, "\nqt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))")
		}
		fmt.Fprint(bb, "\nptr.SetPointer(nil)")
	}

	if !(function.Static || function.Meta == parser.CONSTRUCTOR || function.SignalMode == parser.CALLBACK) {
		fmt.Fprint(bb, "\n}")
		if converter.GoHeaderOutput(function) != "" {
			fmt.Fprintf(bb, "\nreturn %v", converter.GoOutputParametersFromCFailed(function))
		}
	}

	return bb.String()
}
