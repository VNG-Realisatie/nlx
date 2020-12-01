// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nlx-management-api",
	Short: "NLX Management API",
}

func setupCmdFlagForEnvironment(key, value string, cmd *cobra.Command) {
	flag := cmd.Flags().Lookup(key)

	if flag != nil {
		if err := flag.Value.Set(value); err != nil {
			log.Fatal(err)
		}

		// remove required flag if the value was set by an environment variable
		if value != "" {
			delete(flag.Annotations, cobra.BashCompOneRequiredFlag)
		}
	}

	for _, subCmd := range cmd.Commands() {
		setupCmdFlagForEnvironment(key, value, subCmd)
	}
}

// All flags can also be set using environment variables.
// Environment variable names are all caps and '-' is replaced by '_'.
// Example: 'listen-address' becomes 'LISTEN_ADDRESS'
func setupFlagsForEnvironment() {
	// pass environment variables to the arguments
	for _, keyval := range os.Environ() {
		components := strings.SplitN(keyval, "=", 2)

		//nolint:gomnd // key value pair has to have two components
		if len(components) != 2 {
			continue
		}

		key := strings.ReplaceAll(strings.ToLower(components[0]), "_", "-")
		value := components[1]

		setupCmdFlagForEnvironment(key, value, rootCmd)
	}
}

func main() {
	setupFlagsForEnvironment()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
