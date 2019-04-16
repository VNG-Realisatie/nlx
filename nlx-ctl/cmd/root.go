package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "nlx-ctl",
	Short: "nlx-ctl is a command line tool to communicate with the config API",
	Long:  `nlx-ctl is a command line tool to communicate with the config API`,
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName(".nlx-ctl-config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		err = viper.WriteConfigAs(fmt.Sprintf("%s/.nlx-ctl-config.yaml", home))
		if err != nil {
			log.Panic(err)
		}
	}

}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
