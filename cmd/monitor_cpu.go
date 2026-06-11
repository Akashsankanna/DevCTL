package cmd

import (
	monitorpkg "github.com/akash/devctl/internal/monitor"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Show CPU usage information",
	RunE: func(cmd *cobra.Command, args []string) error {
		return monitorpkg.ShowCPU()
	},
}

func init() {
	monitorCmd.AddCommand(cpuCmd)
}
