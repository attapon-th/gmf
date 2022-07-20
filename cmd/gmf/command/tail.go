package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/hpcloud/tail"
	"github.com/spf13/cobra"
)

var (
	tailFollow *bool
	tailLines  *int
)

func init() {
	rootCmd.AddCommand(tailCmd)

	tailFollow = tailCmd.Flags().BoolP("follow", "f", false, "tail defaults to following the file descriptor")
	tailLines = tailCmd.Flags().IntP("lines", "n", 0, "output the last NUM lines")
}

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Print the last number of lines.",
	Long:  `Print tail file `,
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if _, err := os.Stat(args[0]); !errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("filename: %s not exist.", args[0])
	},
	Run: func(_ *cobra.Command, args []string) {
		t, _ := tail.TailFile(args[0], tail.Config{
			Follow: *tailFollow,
		})

		// * tail -n
		if !t.Follow && *tailLines > 0 {
			list := arraylist.New()
			for line := range t.Lines {
				list.Add(line.Text)
				if list.Size() > *tailLines {
					list.Remove(0)
				}
			}
			list.Each(func(_ int, v interface{}) {
				fmt.Println(v)
			})
		} else {
			// * tail -f
			for line := range t.Lines {
				fmt.Println(line.Text)
			}
		}

	},
}
