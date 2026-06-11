package cmd

import (
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker container management commands",
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
