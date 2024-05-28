package cmd

import (
	"fmt"

	"bitbucket.org/zetaactions/pklpoc/configservice"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "configservice",
	Short: "Run the config service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Config Service. Use -h to see more commands")
	},
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Run the config service",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", "Welcome to Trinity. Use -h to see more commands")

			cfgService := configservice.NewConfigService()
			router := cfgService.BuildRouter()
			router.Run(":8080")
		},
	})
}
