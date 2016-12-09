package templater

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func cppEnum(enum *parser.Enum, value *parser.Value) string {
	return fmt.Sprintf("%v\n{\n\t%v\n}", cppEnumHeader(enum, value), cppEnumBody(enum, value))
}

func cppEnumHeader(enum *parser.Enum, value *parser.Value) string {
	return fmt.Sprintf("int %v_%v_Type()", enum.Class(), value.Name)
}

func cppEnumBody(enum *parser.Enum, value *parser.Value) string {
	//needed for sailfish with 5.7 docs
	if strings.HasPrefix(value.Name, "MV_") || strings.HasPrefix(value.Name, "PM_") || strings.HasPrefix(value.Name, "SH_") {
		return fmt.Sprintf(`#if QT_VERSION >= 0x056000
		return %v::%v;
	#else
		return 0;
	#endif`, enum.Class(), value.Name)
	}

	//needed for msys2 with 5.7 docs
	if strings.HasPrefix(value.Name, "PE_") || strings.HasPrefix(value.Name, "SE_") {
		return fmt.Sprintf(`#if QT_VERSION >= 0x057000
		return %v::%v;
	#else
		return 0;
	#endif`, enum.Class(), value.Name)
	}

	return fmt.Sprintf("return %v::%v;", enum.Class(), value.Name)
}
