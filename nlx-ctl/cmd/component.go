package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.nlx.io/nlx/config-api/configapi"
)

func init() {
	rootCmd.AddCommand(componentCommand)
	componentCommand.AddCommand(componentListCommand)
}

var componentListCommand = &cobra.Command{
	Use:   "list",
	Short: "list components",
	Long:  `list components available on the config-api`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().ListComponents(context.Background(), &configapi.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		for _, component := range response.Components {
			log.Println(fmt.Sprintf("kind: %s, name: %s", component.Kind, component.Name))
		}

	},
}

var componentCommand = &cobra.Command{
	Use:   "component",
	Short: "manage components",
	Long:  `manage components`,
}
