package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose indicates if app output should be verbose or not
var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "show verbose output")
	flag.Parse()
}

var rootCmd = &cobra.Command{
	Use:   "create-k8s-app",
	Short: "Create Kubernetes App is a tool for creating simple services that run on Kubernetes",
	Long: `Create Kubernetes App helps you jump start your Kubernetes services by creating all
necessary files that are need for getting your service up and running in Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
