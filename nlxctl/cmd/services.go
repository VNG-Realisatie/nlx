// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/management-api/api"
)

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
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
		response, err := getManagementClient().ListServices(newRequestContext(), &api.ListServicesRequest{})
		if err != nil {
			log.Fatalf("nlxctl list services: failed to list services from api: %v", err)
		}

		for _, service := range response.Services {
			fmt.Printf("%+v\n\n", service)
		}
	},
}

//nolint:dupl // service command looks like inway command
var createServiceCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a service",
	Run: func(cmd *cobra.Command, arg []string) {
		configBytes, err := ioutil.ReadFile(serviceOptions.configPath)
		if err != nil {
			log.Fatalf("nlxctl create service: read config: %v", err)
		}

		serviceConfigs := splitConfigString(string(configBytes))
		for _, configString := range serviceConfigs {
			service := &api.CreateServiceRequest{}

			err = json.Unmarshal([]byte(configString), service)
			if err != nil {
				log.Fatalf("nlxctl create service: unmarshal config json: %v", err)
			}

			_, err = getManagementClient().CreateService(newRequestContext(), service)
			if err != nil {
				log.Fatalf("nlxctl create service: create service api: %v", err)
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

		service := &api.UpdateServiceRequest{}
		err = json.Unmarshal(configBytes, service)
		if err != nil {
			panic(err)
		}

		updateServiceRequest := &api.UpdateServiceRequest{
			Name:                 serviceOptions.name,
			EndpointURL:          service.EndpointURL,
			DocumentationURL:     service.DocumentationURL,
			ApiSpecificationURL:  service.ApiSpecificationURL,
			Internal:             service.Internal,
			TechSupportContact:   service.TechSupportContact,
			PublicSupportContact: service.PublicSupportContact,
			Inways:               service.Inways,
		}

		_, err = getManagementClient().UpdateService(newRequestContext(), updateServiceRequest)
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
		deleteServiceRequest := &api.DeleteServiceRequest{
			Name: serviceOptions.name,
		}

		_, err := getManagementClient().DeleteService(newRequestContext(), deleteServiceRequest)
		if err != nil {
			panic(err)
		}

		println(fmt.Sprintf("deleted service with name: %s", serviceOptions.name))
	},
}
