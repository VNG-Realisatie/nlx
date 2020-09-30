package cmd

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initOptions struct {
	address string
	cert    string
	key     string
	ca      string
}

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
	rootCmd.AddCommand(initCommand)
	initCommand.Flags().StringVarP(&initOptions.address, "address", "a", "", "address of the management-api")
	initCommand.Flags().StringVarP(&initOptions.cert, "cert", "c", "", "path to certificate")
	initCommand.Flags().StringVarP(&initOptions.key, "key", "k", "", "path to private key")
	initCommand.Flags().StringVarP(&initOptions.ca, "ca", "", "", "path to CA used to verify the connection to the management-api")

	err := initCommand.MarkFlagRequired("key")
	if err != nil {
		panic(err)
	}

	err = initCommand.MarkFlagRequired("cert")
	if err != nil {
		panic(err)
	}

	err = initCommand.MarkFlagRequired("address")
	if err != nil {
		panic(err)
	}
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize nlx-ctl",
	Long:  `use init to initialize the nlx-ctl with the address of the NLX management API and cert key pair`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			log.Panic(err)
		}

		viper.Set("api-address", initOptions.address)
		viper.Set("key-path", initOptions.key)
		viper.Set("cert-path", initOptions.cert)
		viper.Set("ca-path", initOptions.ca)
		err = viper.WriteConfigAs(fmt.Sprintf("%s/.nlxctl-config.yaml", home))
		if err != nil {
			log.Panic(err)
		}
	},
}
