package command

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/NoUseFreak/sdf/internal/pkg/navigate"
	"github.com/NoUseFreak/sdf/internal/sdf/clone"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repo in the structured setup",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Must provide a repo")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		repoInput := args[0]
		targetDir, err := clone.CloneRepo(repoInput)

		if err != nil {
			logrus.Error(err)
			return
		}

		navigate.Chdir(targetDir)
	},
}
