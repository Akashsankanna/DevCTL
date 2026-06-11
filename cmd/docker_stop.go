package cmd

import (
	dockerpkg "github.com/akash/devctl/internal/docker"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [container]",
	Short: "Stop a Docker container",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return dockerpkg.StopContainer(args[0])
	},
}

func init() {
	dockerCmd.AddCommand(stopCmd)
}
