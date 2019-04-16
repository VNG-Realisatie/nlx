package cmd

import (
	"context"
	"io/ioutil"

	"github.com/spf13/cobra"
	"go.nlx.io/nlx/config-api/configapi"
)

func init() {
	rootCmd.AddCommand(configCommand)
	configCommand.AddCommand(configCreateCommand)
	configCreateCommand.Flags().StringVarP(&configCreateOptions.configPath, "config-file", "c", "", "path to config file")
	configCreateCommand.Flags().StringVarP(&configCreateOptions.id, "id", "i", "", "id of config")
	configCreateCommand.MarkFlagRequired("config-file")
	configCreateCommand.MarkFlagRequired("id")
}

var configCreateOptions struct {
	configPath string
	id         string
}

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "manage component configuration",
	Long:  `manage component configuration`,
}

var configCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a configuration",
	Run: func(cmd *cobra.Command, arg []string) {
		//	ctx := context.Background()

		b, err := ioutil.ReadFile(configCreateOptions.configPath)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		resp, err := getConfigClient().SetConfig(ctx, &configapi.SetConfigRequest{
			Config: &configapi.Config{
				Kind:   "inway",
				Config: string(b),
			},
			ComponentName: configCreateOptions.id,
		})
		if err != nil {
			panic(err)
		}
		if len(resp.Error) != 0 {
			panic(resp.Error)
		}
	},
}
