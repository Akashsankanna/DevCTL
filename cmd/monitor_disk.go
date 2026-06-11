package cmd

import (
	"github.com/spf13/cobra"

	"github.com/akash/devctl/internal/monitor"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Show disk usage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return monitor.DiskUsage()
	},
}

func init() {
	monitorCmd.AddCommand(diskCmd)
}
