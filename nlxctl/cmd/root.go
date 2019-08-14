package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "nlxctl",
	Short: "nlxctl is a command line tool to communicate with the config API",
	Long:  `nlxctl is a command line tool to communicate with the config API`,
}

func init() { //nolint:gochecknoinits
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}

	viper.SetEnvPrefix("nlx")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.AddConfigPath(home)
	viper.SetConfigName(".nlxctl-config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err == nil {
		return
	}

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		return
	}

	log.Panic(err)
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
