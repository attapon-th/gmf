package command

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/hpcloud/tail"
	"github.com/spf13/cobra"
)

var (
	AppName   string
	Version   string
	Build     string
	DateBuild string
	cfgFile   string
	verbose   bool
)

var rootCmd = &cobra.Command{
	Use:   "gmf",
	Short: "Go Module Formater",
	Long:  `Tools management for golang project`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {

	rootCmd.Long = fmt.Sprintf(`
	Tools management for golang project
	-----------------------------------
	Version: %s
	Build:   %s
	Date:    %s
	-----------------------------------
		`, Version, Build, DateBuild)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "taskfile", "t", "./Taskfile.yaml", "Taskfile path for control project.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print processing")
	// cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func getTaskfile() *arraylist.List {
	var lineConfigs *arraylist.List = arraylist.New()
	t, err := tail.TailFile(cfgFile, tail.Config{Follow: false})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for line := range t.Lines {
		lineConfigs.Add(line.Text)
	}
	return lineConfigs
}

func saveTaskfile(ls *arraylist.List) {
	var buf bytes.Buffer
	ls.Each(func(_ int, v interface{}) {
		_, _ = buf.WriteString(fmt.Sprintln(v))
	})
	if err := ioutil.WriteFile(cfgFile, buf.Bytes(), fs.ModePerm); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("save to file: ", cfgFile)
}

func VPrintf(format string, a ...any) {
	if verbose {
		fmt.Printf(format, a...)
	}

}

func VPrintln(s ...any) {
	if verbose {
		fmt.Println(s...)
	}
}
