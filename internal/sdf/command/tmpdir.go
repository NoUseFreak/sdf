package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tmpDirCmd)
}

var tmpDirCmd = &cobra.Command{
	Use:    "_tmpfile",
	Hidden: true,
	Args:   cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := ioutil.TempFile(os.TempDir(), "*")
		fmt.Println(file.Name())
	},
}
