package cmd

import (
	"github.com/akash/devctl/internal/monitor"
	"github.com/spf13/cobra"
)

var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Show RAM usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return monitor.RAMUsage()
	},
}

func init() {
	monitorCmd.AddCommand(ramCmd)
}
