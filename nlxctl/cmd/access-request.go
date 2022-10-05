// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.nlx.io/nlx/management-api/api"
)

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
	rootCmd.AddCommand(accessRequestCommand)
	accessRequestCommand.AddCommand(listAccessRequestCommand)
	accessRequestCommand.AddCommand(createAccessRequestCommand)
	accessRequestCommand.AddCommand(approveAccessRequestCommand)

	createAccessRequestCommand.Flags().StringVarP(&accessRequestOptions.organizationSerialNumber, "organization", "", "", "Serial number of the organization")
	createAccessRequestCommand.Flags().StringVarP(&accessRequestOptions.serviceName, "service", "", "", "Name of the service")

	listAccessRequestCommand.Flags().StringVarP(&listAccessRequestsOptions.serviceName, "service", "s", "", "Name of the service")

	err := createAccessRequestCommand.MarkFlagRequired("organization")
	if err != nil {
		panic(err)
	}

	err = createAccessRequestCommand.MarkFlagRequired("service")
	if err != nil {
		panic(err)
	}

	approveAccessRequestCommand.Flags().Uint64VarP(&accessRequestOptions.id, "id", "", 0, "Access request ID")
	approveAccessRequestCommand.Flags().StringVarP(&accessRequestOptions.serviceName, "service", "", "", "Name of the service")
}

type accessRequestDetails struct {
	ID                       uint64
	State                    api.AccessRequestState
	OrganizationName         string
	OrganizationSerialNumber string
	ServiceName              string
	CreatedAt                *timestamppb.Timestamp
	UpdatedAt                *timestamppb.Timestamp
}

func printAccessRequest(details accessRequestDetails) {
	createdAt := timestamppb.New(details.CreatedAt.AsTime())
	updatedAt := timestamppb.New(details.UpdatedAt.AsTime())
	state := "UNKNOWN"

	if name, ok := api.AccessRequestState_name[int32(details.State)]; ok {
		state = name
	}

	fmt.Printf(
		"ID:%d\tState:%s\tOrganization:%s\tService:%s\tCreatedAt:%s\tUpdatedAt:%s\n",
		details.ID,
		state,
		details.OrganizationName,
		details.ServiceName,
		createdAt.AsTime().Format(time.RFC3339),
		updatedAt.AsTime().Format(time.RFC3339),
	)
}

var accessRequestCommand = &cobra.Command{
	Use:   "access-request",
	Short: "Manage access requests",
}

var accessRequestOptions struct {
	id                       uint64
	organizationSerialNumber string
	serviceName              string
}

var approveAccessRequestCommand = &cobra.Command{
	Use:   "approve",
	Short: "Approve an existing incoming access request",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getManagementClient().ApproveIncomingAccessRequest(newRequestContext(), &api.ApproveIncomingAccessRequestRequest{
			ServiceName:     accessRequestOptions.serviceName,
			AccessRequestId: accessRequestOptions.id,
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Access request ID=%d is approved\n", accessRequestOptions.id)
	},
}

var createAccessRequestCommand = &cobra.Command{
	Use:   "create",
	Short: "Request access to another service",
	Run: func(cmd *cobra.Command, args []string) {
		accessRequest, err := getManagementClient().SendAccessRequest(newRequestContext(), &api.SendAccessRequestRequest{
			OrganizationSerialNumber: accessRequestOptions.organizationSerialNumber,
			ServiceName:              accessRequestOptions.serviceName,
		})
		if err != nil {
			log.Fatal(err)
		}

		printAccessRequest(accessRequestDetails{
			ID:                       accessRequest.OutgoingAccessRequest.Id,
			State:                    accessRequest.OutgoingAccessRequest.State,
			OrganizationSerialNumber: accessRequest.OutgoingAccessRequest.Organization.SerialNumber,
			ServiceName:              accessRequest.OutgoingAccessRequest.ServiceName,
			CreatedAt:                accessRequest.OutgoingAccessRequest.CreatedAt,
			UpdatedAt:                accessRequest.OutgoingAccessRequest.UpdatedAt,
		})
	},
}

var listAccessRequestsOptions struct {
	serviceName string
}

//nolint:dupl // incoming/outgoing requests aren't duplicates
var listAccessRequestCommand = &cobra.Command{
	Use:   "list",
	Short: "List incoming access requests",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getManagementClient().ListIncomingAccessRequests(newRequestContext(), &api.ListIncomingAccessRequestsRequest{
			ServiceName: listAccessRequestsOptions.serviceName,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, accessRequest := range response.AccessRequests {
			printAccessRequest(accessRequestDetails{
				ID:                       accessRequest.Id,
				State:                    accessRequest.State,
				OrganizationName:         accessRequest.Organization.Name,
				OrganizationSerialNumber: accessRequest.Organization.SerialNumber,
				ServiceName:              accessRequest.ServiceName,
				CreatedAt:                accessRequest.CreatedAt,
				UpdatedAt:                accessRequest.UpdatedAt,
			})
		}
	},
}
