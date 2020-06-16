// nolint:dupl
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/pkg/configapi"
)

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(serviceCommand)
	serviceCommand.AddCommand(listServicesCommand)
	serviceCommand.AddCommand(createServiceCommand)
	serviceCommand.AddCommand(updateServiceCommand)
	serviceCommand.AddCommand(deleteServiceCommand)

	listServicesCommand.Flags().StringVarP(&serviceListOptions.inwayName, "inway", "i", "", "name of the inway of which you want to list the services")

	createServiceCommand.Flags().StringVarP(&serviceOptions.configPath, "config", "c", "", "config of service")
	err := createServiceCommand.MarkFlagRequired("config")
	if err != nil {
		panic(err)
	}

	updateServiceCommand.Flags().StringVarP(&serviceOptions.name, "name", "n", "", "name of service")
	err = updateServiceCommand.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	updateServiceCommand.Flags().StringVarP(&serviceOptions.configPath, "config", "c", "", "config of service")
	err = createServiceCommand.MarkFlagRequired("config")
	if err != nil {
		panic(err)
	}

	deleteServiceCommand.Flags().StringVarP(&serviceOptions.name, "name", "n", "", "name of service")
	err = deleteServiceCommand.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}

var serviceListOptions struct {
	inwayName string
}

var serviceOptions struct {
	name       string
	configPath string
}

var serviceCommand = &cobra.Command{
	Use:   "service",
	Short: "Manage services",
}

var listServicesCommand = &cobra.Command{
	Use:   "list",
	Short: "List services",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getConfigClient().ListServices(context.Background(), &configapi.ListServicesRequest{})
		if err != nil {
			log.Fatal(err)
		}

		for _, service := range response.Services {
			fmt.Printf("%+v\n\n", service)
		}
	},
}

var createServiceCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a service",
	Run: func(cmd *cobra.Command, arg []string) {
		configBytes, err := ioutil.ReadFile(serviceOptions.configPath)
		if err != nil {
			panic(err)
		}

		serviceConfigs := splitConfigString(string(configBytes))
		for _, configString := range serviceConfigs {
			service := &configapi.Service{}
			err = json.Unmarshal([]byte(configString), service)
			if err != nil {
				panic(err)
			}

			ctx := context.Background()
			_, err = getConfigClient().CreateService(ctx, service)
			if err != nil {
				panic(err)
			}

			println(fmt.Sprintf("created service: %+v", service))
		}

	},
}

var updateServiceCommand = &cobra.Command{
	Use:   "update",
	Short: "Update a service",
	Run: func(cmd *cobra.Command, arg []string) {
		configBytes, err := ioutil.ReadFile(serviceOptions.configPath)
		if err != nil {
			panic(err)
		}

		service := &configapi.Service{}
		err = json.Unmarshal(configBytes, service)
		if err != nil {
			panic(err)
		}

		updateServiceRequest := &configapi.UpdateServiceRequest{
			Name:    serviceOptions.name,
			Service: service,
		}

		ctx := context.Background()
		_, err = getConfigClient().UpdateService(ctx, updateServiceRequest)
		if err != nil {
			panic(err)
		}

		println(fmt.Sprintf("updated service with name: %s", serviceOptions.name))
	},
}

var deleteServiceCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a service",
	Run: func(cmd *cobra.Command, arg []string) {
		ctx := context.Background()

		deleteServiceRequest := &configapi.DeleteServiceRequest{
			Name: serviceOptions.name,
		}

		_, err := getConfigClient().DeleteService(ctx, deleteServiceRequest)
		if err != nil {
			panic(err)
		}

		println(fmt.Sprintf("deleted service with name: %s", serviceOptions.name))
	},
}
