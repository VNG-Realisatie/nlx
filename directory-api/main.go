// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package main

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/directory-api/cmd"
)

func setupCmdFlagForEnvironment(key, value string, command *cobra.Command) {
	flag := command.Flags().Lookup(key)

	if flag != nil {
		if err := flag.Value.Set(value); err != nil {
			log.Fatal(err)
		}

		// remove required flag if the value was set by an environment variable
		if value != "" {
			delete(flag.Annotations, cobra.BashCompOneRequiredFlag)
		}
	}

	for _, subCmd := range command.Commands() {
		setupCmdFlagForEnvironment(key, value, subCmd)
	}
}

// All flags can also be set using environment variables.
// Environment variable names are all caps and '-' is replaced by '_'.
// Example: 'listen-address' becomes 'LISTEN_ADDRESS'
func setupFlagsForEnvironment(rootCmd *cobra.Command) {
	// pass environment variables to the arguments
	for _, keyval := range os.Environ() {
		const amountOfParts = 2
		components := strings.SplitN(keyval, "=", amountOfParts)

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
	setupFlagsForEnvironment(cmd.RootCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
