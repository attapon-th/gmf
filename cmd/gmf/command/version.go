package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	versionInc *bool
)

func init() {
	rootCmd.AddCommand(versionCmd)
	versionInc = versionCmd.Flags().BoolP("increment", "i", false, "Increment version in the revision do count.")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the Version.",
	Long:  `Show the Version `,
	Run: func(_ *cobra.Command, _ []string) {
		// fmt.Printf("Current Version: %s n", viper.GetString("vars.app_version"))
		taskfile := getTaskfile()
		lt := taskfile.Iterator()
		version := "0.0.0"
		versionLines := []string{}
		versionIndex := -1
		for lt.Next() {
			s := fmt.Sprint(lt.Value())
			if strings.Contains(s, "APP_VERSION:") {
				versionLines = strings.SplitN(s, ": ", 2)
				if len(versionLines) < 2 {

					break
				}
				versionIndex = lt.Index()
				s = strings.TrimSpace(versionLines[1])
				if strings.Count(s, ".") == 2 {
					version = s
				}
				break
			}
		}

		if versionIndex < 1 {
			fmt.Println("Not exist variable\n---\nvars: \n    APP_VERSION: x.x.x")
			os.Exit(1)
		}
		fmt.Printf("Current Version: %v\n", version)

		if *versionInc {
			s := version
			sp := strings.SplitN(s, ".", 3)
			if len(sp) < 3 {
				sp = []string{"0", "0", "0"}
			}
			v, _ := strconv.Atoi(sp[2])
			v++
			sp[2] = strconv.Itoa(v)
			version = strings.Join(sp, ".")
			versionLines[1] = version
			taskfile.Set(versionIndex, strings.Join(versionLines, ": "))
			fmt.Printf("New Version: %s\n", version)
			saveTaskfile(taskfile)
		}
	},
}

// func getVersion() {
// 	for _, l := range
// }
