package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"go.nlx.io/nlx/config-api/configapi"
)

var announceOptions struct {
	componentType string
	signature     string
}

var acknowledgeOptions struct {
	Identifier string
}

func init() {
	rootCmd.AddCommand(announceCommand)
	announceCommand.AddCommand(announceCreateCommand)
	announceCreateCommand.Flags().StringVarP(&announceOptions.componentType, "type", "t", "", "component type to announce")
	announceCreateCommand.Flags().StringVarP(&announceOptions.componentType, "signature", "s", "", "signature of component")
	announceCreateCommand.MarkFlagRequired("type")
	announceCommand.AddCommand(announceListCommand)
}

var announceListCommand = &cobra.Command{
	Use:   "list",
	Short: "list announcements",
	Long:  `list announcements available on the config-api`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().ListAnnouncements(context.Background(), &configapi.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		for _, announcement := range response.Announcements {
			log.Println(announcement)
		}

	},
}

var announceCreateCommand = &cobra.Command{
	Use:   "create",
	Short: "announce yourself to the config api",
	Long:  `announce yourself to the config api`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().Announce(context.Background(), &configapi.AnnounceRequest{
			ComponentName: announceOptions.componentType,
			Signature:     announceOptions.signature,
		})

		if err != nil {
			log.Fatal(err)
		}

		if len(response.Error) > 0 {
			log.Fatalf("error anouncing component %s", response.Error)
		}

		log.Println("component announched")
	},
}

var announceCommand = &cobra.Command{
	Use:   "announce",
	Short: "manage announcements",
	Long:  `manage announcements`,
}
