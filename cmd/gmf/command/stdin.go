package command

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/spf13/cobra"
)

var (
	stdinLines int
	stdinTrim  bool
	stdinRegex bool
)

func init() {
	rootCmd.AddCommand(stdinCmd)
	stdinCmd.Flags().IntVarP(&stdinLines, "number", "n", 0, "Get with line number, start line at 1.")
	stdinCmd.Flags().BoolVarP(&stdinTrim, "space", "s", false, "strings.TrimSpace")
	stdinCmd.Flags().BoolVarP(&stdinRegex, "regex", "r", false, "regexp.MustCompile(args[0]).ReplaceAllString(input, args[1])")
}

var stdinCmd = &cobra.Command{
	Use:   "stdin",
	Short: "STDIN Processing",
	Long:  `STDIN Processing`,
	Run: func(_ *cobra.Command, args []string) {
		var rx *regexp.Regexp
		sc := bufio.NewScanner(os.Stdin)
		l := arraylist.New()
		if stdinRegex && len(args) == 2 {
			VPrintf("Regular Expression\nRegex: %q, Replace: %q\n", args[0], args[1])
			rx = regexp.MustCompile(args[0])
		}
		for sc.Scan() {
			s := sc.Text()
			if rx != nil {
				s = rx.ReplaceAllString(s, args[1])
			}
			l.Add(s)
		}
		if stdinLines > 0 {
			if v, ok := l.Get(stdinLines - 1); ok {
				s := fmt.Sprint(v)
				if stdinTrim {
					s = strings.TrimSpace(s)
				}
				fmt.Println(s)
			}
		} else {
			printAll(l)
		}

	},
}

func printAll(l *arraylist.List) {
	l.Each(func(_ int, v interface{}) {
		fmt.Println(v)
	})
}
