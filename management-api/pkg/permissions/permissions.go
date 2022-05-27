// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package permissions

type Permission int

//go:generate go run github.com/dmarkham/enumer -type=Permission -linecomment

// The comments after the enum determines the string representation of the enum value
const (
	ApproveIncomingAccessRequest Permission = iota + 1 // permissions.incoming_access_request.approve
	RejectIncomingAccessRequest                        // permissions.incoming_access_request.reject
	ReadIncomingAccessRequests                         // permissions.incoming_access_requests.read
	CreateOutgoingAccessRequest                        // permissions.outgoing_access_request.create
	UpdateOutgoingAccessRequest                        // permissions.outgoing_access_request.update
	ReadAccessGrants                                   // permissions.access_grants.read
	RevokeAccessGrant                                  // permissions.access_grant.revoke
	ReadAuditLogs                                      // permissions.audit_logs.read
	ReadFinanceReport                                  // permissions.finance_report.read
	ReadInway                                          // permissions.inway.read
	UpdateInway                                        // permissions.inway.update
	DeleteInway                                        // permissions.inway.delete
	ReadInways                                         // permissions.inways.read
	CreateOutgoingOrder                                // permissions.outgoing_order.create
	UpdateOutgoingOrder                                // permissions.outgoing_order.update
	RevokeOutgoingOrder                                // permissions.outgoing_order.revoke
	ReadOutgoingOrders                                 // permissions.outgoing_orders.read
	ReadIncomingOrders                                 // permissions.incoming_orders.read
	SynchronizeIncomingOrders                          // permissions.incoming_orders.synchronize
	ReadOutways                                        // permissions.outways.read
	CreateService                                      // permissions.service.create
	ReadService                                        // permissions.service.read
	UpdateService                                      // permissions.service.update
	DeleteService                                      // permissions.service.delete
	ReadServices                                       // permissions.services.read
	ReadServicesStatistics                             // permissions.services_statistics.read
	ReadOrganizationSettings                           // permissions.organization_settings.read
	UpdateOrganizationSettings                         // permissions.organization_settings.update
	AcceptTermsOfService                               // permissions.terms_of_service.accept
	ReadTermsOfServiceStatus                           // permissions.terms_of_service_status.read
	ReadTransactionLogs                                // permissions.transaction_logs.read
)
