// nolint:dupl
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"go.nlx.io/nlx/management-api/configapi"
)

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(inwayCommand)
	inwayCommand.AddCommand(listInwaysCommand)
	inwayCommand.AddCommand(createInwayCommand)
	inwayCommand.AddCommand(updateInwayCommand)
	inwayCommand.AddCommand(deleteInwayCommand)

	createInwayCommand.Flags().StringVarP(&inwayOptions.configPath, "config", "c", "", "config of inway")
	err := createInwayCommand.MarkFlagRequired("config")
	if err != nil {
		panic(err)
	}

	updateInwayCommand.Flags().StringVarP(&inwayOptions.name, "name", "n", "", "name of inway")
	err = updateInwayCommand.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	updateInwayCommand.Flags().StringVarP(&inwayOptions.configPath, "config", "c", "", "config of inway")
	err = updateInwayCommand.MarkFlagRequired("config")
	if err != nil {
		panic(err)
	}

	deleteInwayCommand.Flags().StringVarP(&inwayOptions.name, "name", "n", "", "name of inway")
	err = deleteInwayCommand.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}

var inwayOptions struct {
	name       string
	configPath string
}

var inwayCommand = &cobra.Command{
	Use:   "inway",
	Short: "Manage inways",
}

var listInwaysCommand = &cobra.Command{
	Use:   "list",
	Short: "List inways",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().ListInways(context.Background(), &configapi.ListInwaysRequest{})
		if err != nil {
			log.Fatal(err)
		}

		for _, inway := range response.Inways {
			fmt.Printf("%+v\n\n", inway)
		}
	},
}

var createInwayCommand = &cobra.Command{
	Use:   "create",
	Short: "Create an inway",
	Run: func(cmd *cobra.Command, arg []string) {
		configBytes, err := ioutil.ReadFile(inwayOptions.configPath)
		if err != nil {
			panic(err)
		}

		inwayConfig := splitConfigString(string(configBytes))
		for _, configString := range inwayConfig {
			inway := &configapi.Inway{}
			err = json.Unmarshal([]byte(configString), inway)
			if err != nil {
				panic(err)
			}

			ctx := context.Background()
			_, err = getConfigClient().CreateInway(ctx, inway)
			if err != nil {
				panic(err)
			}

			println(fmt.Sprintf("created inway with name: %+v", inway))
		}

	},
}

var updateInwayCommand = &cobra.Command{
	Use:   "update",
	Short: "Update an inway",
	Run: func(cmd *cobra.Command, arg []string) {
		configBytes, err := ioutil.ReadFile(inwayOptions.configPath)
		if err != nil {
			panic(err)
		}

		inway := &configapi.Inway{}
		err = json.Unmarshal(configBytes, inway)
		if err != nil {
			panic(err)
		}

		updateInwayRequest := &configapi.UpdateInwayRequest{
			Name:  inwayOptions.name,
			Inway: inway,
		}

		ctx := context.Background()
		_, err = getConfigClient().UpdateInway(ctx, updateInwayRequest)
		if err != nil {
			panic(err)
		}

		println(fmt.Sprintf("updated inway with name: %s", inwayOptions.name))
	},
}

var deleteInwayCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete an inway",
	Run: func(cmd *cobra.Command, arg []string) {
		ctx := context.Background()

		deleteInwayRequest := &configapi.DeleteInwayRequest{
			Name: inwayOptions.name,
		}

		_, err := getConfigClient().DeleteInway(ctx, deleteInwayRequest)
		if err != nil {
			panic(err)
		}

		println(fmt.Sprintf("deleted inway with name: %s", inwayOptions.name))
	},
}
