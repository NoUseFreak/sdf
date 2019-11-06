package command

import (
	"fmt"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/wdid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wdidCmd)
}

var wdidCmd = &cobra.Command{
	Use:   "wdid [amount] [unit]",
	Short: "What did I do?",
	Run: func(cmd *cobra.Command, args []string) {
		output.Println("Generating WhatDidIDo report")

		amount := "1"
		unit := "day"

		if len(args) >= 1 {
			amount = args[0]
		}
		if len(args) >= 2 {
			unit = args[1]
		}

		wdid.PrintReport(fmt.Sprintf("%s %s ago", amount, unit))
	},
}
