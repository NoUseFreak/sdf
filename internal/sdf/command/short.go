package command

import (
	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/short"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shortCmd)

	shortCmd.AddCommand(shortListCmd)
	shortCmd.AddCommand(shortNewCmd)
	shortCmd.AddCommand(shortRemoveCmd)
}

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "short",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		short.RunShort(args[0])
	},
}

var shortListCmd = &cobra.Command{
	Use:   "list",
	Short: "list available shorts",
	Run: func(cmd *cobra.Command, args []string) {
		output.Println("Available short commands:")
		for _, s := range short.List() {
			output.Println(" - %-10v (%s)", s.Name, s.Cmd)
		}
	},
}

var shortNewCmd = &cobra.Command{
	Use:     "new <name> <command>",
	Aliases: []string{"add", "create"},
	Short:   "Create a new short",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		output.Println("Saving command")
		short.Add(args[0], args[1])
	},
}

var shortRemoveCmd = &cobra.Command{
	Use:     "rm <name>",
	Aliases: []string{"remove", "del", "delete"},
	Short:   "Remove a short",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := short.Remove(args[0]); err != nil {
			output.Error(err.Error())
		}
	},
}
