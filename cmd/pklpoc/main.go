package main

import (
	"fmt"
	"os"

	configserviceCmd "bitbucket.org/zetaactions/pklpoc/configservice/cmd"
	exampleserviceCmd "bitbucket.org/zetaactions/pklpoc/example/cmd"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "pklpoc",
	Short: "Kosmos + Zuri + Zasper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Trinity. Use -h to see more commands")
	},
}

// You will additionally define flags and handle configuration in your init() function.
func init() {
	rootCmd.AddCommand(configserviceCmd.RootCmd)  // add all cosmos commands
	rootCmd.AddCommand(exampleserviceCmd.RootCmd) // add all cosmos commands
}
