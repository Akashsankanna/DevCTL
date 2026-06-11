package cmd

import (
	dockerpkg "github.com/akash/devctl/internal/docker"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [container]",
	Short: "Restart a Docker container",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return dockerpkg.RestartContainer(args[0])
	},
}

func init() {
	dockerCmd.AddCommand(restartCmd)
}
