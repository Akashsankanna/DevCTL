package cmd

import (
	"fmt"

	dockercmd "github.com/akash/devctl/internal/docker"

	"github.com/spf13/cobra"
)

var dockerPsCmd = &cobra.Command{
	Use:   "ps",
	Short: "List docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		if err := dockercmd.ListContainers(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	dockerCmd.AddCommand(dockerPsCmd)
}
