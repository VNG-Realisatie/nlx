// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "nlx-management-api",
	Short: "NLX Management API",
}

func Execute() error {
	return RootCmd.Execute()
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	RootCmd.AddCommand(serveCommand)
	RootCmd.AddCommand(createUserCommand)
	RootCmd.AddCommand(migrateCommand)
}
