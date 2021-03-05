package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/management-api/api"
)

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
	rootCmd.AddCommand(accessRequestCommand)
	accessRequestCommand.AddCommand(listAccessRequestCommand)
	accessRequestCommand.AddCommand(listOutgoingAccessRequestCommand)
	accessRequestCommand.AddCommand(createAccessRequestCommand)
	accessRequestCommand.AddCommand(approveAccessRequestCommand)

	createAccessRequestCommand.Flags().StringVarP(&accessRequestOptions.organizationName, "organization", "", "", "Name of the organization")
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
	ID               uint64
	State            api.AccessRequestState
	OrganizationName string
	ServiceName      string
	CreatedAt        *types.Timestamp
	UpdatedAt        *types.Timestamp
}

func printAccessRequest(details accessRequestDetails) {
	createdAt, err := types.TimestampFromProto(details.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	updatedAt, err := types.TimestampFromProto(details.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}

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
		createdAt.Format(time.RFC3339),
		updatedAt.Format(time.RFC3339),
	)
}

var accessRequestCommand = &cobra.Command{
	Use:   "access-request",
	Short: "Manage access requests",
}

var accessRequestOptions struct {
	id               uint64
	organizationName string
	serviceName      string
}

var approveAccessRequestCommand = &cobra.Command{
	Use:   "approve",
	Short: "Approve an existing incoming access request",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getManagementClient().ApproveIncomingAccessRequest(context.Background(), &api.ApproveIncomingAccessRequestRequest{
			ServiceName:     accessRequestOptions.serviceName,
			AccessRequestID: accessRequestOptions.id,
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
		accessRequest, err := getManagementClient().CreateAccessRequest(context.Background(), &api.CreateAccessRequestRequest{
			OrganizationName: accessRequestOptions.organizationName,
			ServiceName:      accessRequestOptions.serviceName,
		})
		if err != nil {
			log.Fatal(err)
		}

		printAccessRequest(accessRequestDetails{
			ID:               accessRequest.Id,
			State:            accessRequest.State,
			OrganizationName: accessRequest.OrganizationName,
			ServiceName:      accessRequest.ServiceName,
			CreatedAt:        accessRequest.CreatedAt,
			UpdatedAt:        accessRequest.UpdatedAt,
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
		response, err := getManagementClient().ListIncomingAccessRequest(context.Background(), &api.ListIncomingAccessRequestsRequests{
			ServiceName: listAccessRequestsOptions.serviceName,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, accessRequest := range response.AccessRequests {
			printAccessRequest(accessRequestDetails{
				ID:               accessRequest.Id,
				State:            accessRequest.State,
				OrganizationName: accessRequest.OrganizationName,
				ServiceName:      accessRequest.ServiceName,
				CreatedAt:        accessRequest.CreatedAt,
				UpdatedAt:        accessRequest.UpdatedAt,
			})
		}
	},
}

//nolint:dupl // incoming/outgoing requests aren't duplicates
var listOutgoingAccessRequestCommand = &cobra.Command{
	Use:   "list-outgoing",
	Short: "List outgoing access requests",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := getManagementClient().ListOutgoingAccessRequests(context.Background(), &api.ListOutgoingAccessRequestsRequest{})
		if err != nil {
			log.Fatal(err)
		}

		for _, accessRequest := range response.AccessRequests {
			printAccessRequest(accessRequestDetails{
				ID:               accessRequest.Id,
				State:            accessRequest.State,
				OrganizationName: accessRequest.OrganizationName,
				ServiceName:      accessRequest.ServiceName,
				CreatedAt:        accessRequest.CreatedAt,
				UpdatedAt:        accessRequest.UpdatedAt,
			})
		}
	},
}
