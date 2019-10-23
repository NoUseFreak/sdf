package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/setup"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set it up",
	Run: func(cmd *cobra.Command, args []string) {
		homeCfgDir := setup.CreateHomeDir()
		setup.CreateConfig(homeCfgDir)
		setup.SetupShellFunc()

		output.Print("Completed setting up your %s profile\n", viper.GetString("profile"))

	},
}
