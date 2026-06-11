package cmd

import (
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "System monitoring commands",
	Long:  "Monitor CPU, RAM and Disk usage.",
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
