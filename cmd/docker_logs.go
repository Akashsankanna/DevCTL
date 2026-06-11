package cmd

import (
	dockerpkg "github.com/akash/devctl/internal/docker"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [container]",
	Short: "Show logs of a Docker container",
	Long:  "Display logs from a running or stopped Docker container.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return dockerpkg.ShowLogs(args[0])
	},
}

func init() {
	dockerCmd.AddCommand(logsCmd)
}
