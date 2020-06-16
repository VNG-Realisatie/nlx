// nolint:dupl
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/pkg/configapi"
)

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(insightCommand)
	insightCommand.AddCommand(putInsightCommand)
	insightCommand.AddCommand(getInsightCommand)

	putInsightCommand.Flags().StringVarP(&insightOptions.insightAPIURL, "insight-api-url", "i", "insight", "URL of the insight api")
	putInsightCommand.Flags().StringVarP(&insightOptions.irmaServerURL, "irma-server-url", "r", "irma", "URL of the irma server")
	err := putInsightCommand.MarkFlagRequired("insight-api-url")
	if err != nil {
		panic(err)
	}
	err = putInsightCommand.MarkFlagRequired("irma-server-url")
	if err != nil {
		panic(err)
	}
}

var insightOptions struct {
	insightAPIURL string
	irmaServerURL string
}

var insightCommand = &cobra.Command{
	Use:   "insight",
	Short: "Manage the insight configuration",
}

var putInsightCommand = &cobra.Command{
	Use:   "put",
	Short: "Set the insight API url and the Irma server url",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().PutInsightConfiguration(context.Background(), &configapi.InsightConfiguration{
			InsightAPIURL: insightOptions.insightAPIURL,
			IrmaServerURL: insightOptions.irmaServerURL,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n\n", response)

	},
}

var getInsightCommand = &cobra.Command{
	Use:   "get",
	Short: "Returns the current insight configuration",
	Run: func(cmd *cobra.Command, arg []string) {

		response, err := getConfigClient().GetInsightConfiguration(context.Background(), &configapi.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n\n", response)
	},
}
