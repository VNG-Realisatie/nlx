package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configName = ".nlxctl-config"
	configType = "yaml"
)

var configLocation string

var rootCmd = &cobra.Command{
	Use:   "nlxctl",
	Short: "nlxctl is a command line tool to communicate with the config API",
	Long:  `nlxctl is a command line tool to communicate with the config API`,
}

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)
	}

	home := usr.HomeDir

	viper.SetEnvPrefix("nlx")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.AddConfigPath(home)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	configLocation = fmt.Sprintf("%s/%s.%s", home, configName, configType)

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
