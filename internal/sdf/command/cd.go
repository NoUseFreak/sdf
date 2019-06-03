package command

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/NoUseFreak/sdf/internal/pkg/navigate"
	"github.com/NoUseFreak/sdf/internal/sdf/cd"
)

func init() {
	rootCmd.AddCommand(cdCmd)
}

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "cd dir",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Must provide a name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, "")
		input = strings.ReplaceAll(input, "/", "")
		targetDir, err := cd.Find(input)

		if err != nil {
			logrus.Error(err)
			return
		}

		if targetDir != "" {
			navigate.Chdir(targetDir)
		}
	},
}
