package cmd

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/zetaactions/pklpoc/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "exampleservice",
	Short: "Run the example service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", "Welcome to Example Service. Use -h to see more commands")
	},
}

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the config service",
		// flags
		// args

		// run
		Run: func(cmd *cobra.Command, args []string) {
			environment := cmd.Flag("environment").Value.String()
			formation := cmd.Flag("formation").Value.String()
			cfgUrl := fmt.Sprintf("http://localhost:8080/static/environment/%s/%s/main.pkl", environment, formation)

			cfg, err := config.LoadFromHTTP(cmd.Context(), cfgUrl)
			if err != nil {
				panic(err)
			}

			// TODO: bootstrap my app with the config on cfg

			encoded, err := json.Marshal(cfg)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(encoded))
		},
	}
	runCmd.PersistentFlags().StringP("environment", "e", "local", "environment name dev/local/prod")
	runCmd.PersistentFlags().StringP("formation", "f", "default", "formation name")
	RootCmd.AddCommand(runCmd)
}
