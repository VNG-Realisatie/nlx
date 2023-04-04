// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "nlx-directory-ui",
	Short: "NLX Directory UI",
}

func Execute() error {
	return RootCmd.Execute()
}

//nolint:gochecknoinits // this is the recommended way to use cobra
func init() {
	RootCmd.AddCommand(serveCommand)
}
