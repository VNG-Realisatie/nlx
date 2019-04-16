package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initOptions struct {
	address string
	cert    string
	key     string
	ca      string
}

func init() {
	rootCmd.AddCommand(initCommand)
	initCommand.Flags().StringVarP(&initOptions.address, "address", "a", "", "address of the config-api")
	initCommand.Flags().StringVarP(&initOptions.cert, "cert", "c", "", "path to certificate")
	initCommand.Flags().StringVarP(&initOptions.key, "key", "k", "", "path to private key")
	initCommand.Flags().StringVarP(&initOptions.ca, "ca", "", "", "path to CA used to verify the connection to the config-api")

	initCommand.MarkFlagRequired("key")
	initCommand.MarkFlagRequired("cert")
	initCommand.MarkFlagRequired("address")
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize nlx-ctl",
	Long:  `use init to initialize the nlx-ctl with the address of the config-api and cert key pair`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.Set("api-address", initOptions.address)
		viper.Set("key-path", initOptions.key)
		viper.Set("cert-path", initOptions.cert)
		viper.Set("ca-path", initOptions.ca)
		err := viper.WriteConfig()
		if err != nil {
			log.Panic(err)
		}
	},
}
