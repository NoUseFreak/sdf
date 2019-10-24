package command

import (
	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/env"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.AddCommand(envCatCmd)
	envCmd.AddCommand(envEditCmd)
	envCmd.AddCommand(envListCmd)
	envCmd.AddCommand(envNewCmd)
	envCmd.AddCommand(envRemoveCmd)
	envCmd.AddCommand(envSaveCmd)
	envCmd.AddCommand(envUseCmd)
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "env",
}

var envCatCmd = &cobra.Command{
	Use:     "cat",
	Aliases: []string{"view"},
	Short:   "Show and env set",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Cat(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}
var envEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an env set",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Edit(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available env sets",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range env.List() {
			output.Println(m.Name)
		}
	},
}
var envNewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"add", "create"},
	Short:   "Create a new env set",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.New(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}

var envRemoveCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "del", "delete"},
	Short:   "Remove an env set",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Remove(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}
var envSaveCmd = &cobra.Command{
	Use:    "_savetmpfile",
	Hidden: true,
	Args:   cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Save(args[0], args[1]); err != nil {
			output.Error(err.Error())
		}
	},
}
var envUseCmd = &cobra.Command{
	Use:   "use",
	Short: "Use and env set",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.Use(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}
