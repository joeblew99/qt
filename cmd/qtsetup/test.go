package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func test(buildTarget string) {
	utils.Log.Infof("running setup/test %v (~10min)", buildTarget)

	//TODO: cleanup
	for _, example := range []string{filepath.Join("widgets", "line_edits"), filepath.Join("widgets", "video_player"), filepath.Join("widgets", "graphicsscene"), filepath.Join("widgets", "dropsite"), filepath.Join("widgets", "table"),
		filepath.Join("widgets", "treeview", "treeview_dual"), filepath.Join("widgets", "treeview", "treeview_filelist"), filepath.Join("widgets", "bridge2"), filepath.Join("widgets", "systray"), filepath.Join("widgets", "renderer"),
		filepath.Join("widgets", "textedit"),
		filepath.Join("quick", "bridge"), filepath.Join("quick", "bridge2"), filepath.Join("quick", "calc"), filepath.Join("quick", "dialog"), filepath.Join("quick", "sailfish"), filepath.Join("quick", "translate"), filepath.Join("quick", "view"),
		filepath.Join("quick", "tableview"), filepath.Join("quick", "dynamic"),
		filepath.Join("qml", "application"), filepath.Join("qml", "material"), filepath.Join("qml", "prop"),
		filepath.Join("uitools", "calculator")} {

		if (buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && (!strings.Contains(example, "quick") || (example == filepath.Join("quick", "dynamic") || example == filepath.Join("quick", "tableview") || example == filepath.Join("quick", "bridge") || example == filepath.Join("quick", "dialog"))) {
		} else if !(buildTarget == "sailfish" || buildTarget == "sailfish-emulator") && example == filepath.Join("quick", "sailfish") {
		} else if buildTarget != "desktop" && example == filepath.Join("widgets", "textedit") {
		} else {

			utils.Log.Infoln("testing", example)

			utils.RunCmd(exec.Command(filepath.Join(utils.MustGoBin(), "qtdeploy"),

				"test",

				strings.TrimSuffix(buildTarget, "-docker"),

				filepath.Join(utils.GoQtPkgPath("internal", "examples"), example),

				func() string {
					if strings.HasSuffix(buildTarget, "-docker") {
						return "docker"
					}
					return ""
				}()),

				fmt.Sprintf("test %v", example))
		}
	}
}
