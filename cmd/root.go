package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devctl",
	Short: "Production Devops cli tool",
	Long: ` Devctl is a production grade devops cli tool built using go and cobra 

	features:
	-Docker container managment
	-Backend/api health monitoring 
	-PostgresSQL health check 
	-Deployment automation 
	-System monitring
	-Devops workflow automation

	Example Command:
	devctl docker ps
	devctl docker logs  backend 
	devctl server health 
	devctl db postgres 
	devctl deploy dev 
	devctl monitor cpu  
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP(
		"toggle",
		"t",
		false,
		"help message for toggle",
	)
}
