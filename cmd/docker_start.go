package cmd

import (
	dockerpkg "github.com/akash/devctl/internal/docker"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [container]",
	Short: "Start a Docker container",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return dockerpkg.StartContainer(args[0])
	},
}

func init() {
	dockerCmd.AddCommand(startCmd)
}
