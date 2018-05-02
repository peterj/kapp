package cmd

import (
	"fmt"
	"runtime"

	"github.com/peterj/kapp/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`version     : %s
git hash    : %s
go version  : %s
go compiler : %s
platform    : %s/%s`, version.VERSION, version.GITCOMMIT,
			runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
	},
}
