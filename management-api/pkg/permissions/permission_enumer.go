// Code generated by "enumer -type=Permission -linecomment"; DO NOT EDIT.

package permissions

import (
	"fmt"
	"strings"
)

const _PermissionName = "permissions.incoming_access_request.approvepermissions.incoming_access_request.rejectpermissions.incoming_access_requests.readpermissions.outgoing_access_request.createpermissions.outgoing_access_request.updatepermissions.access_grants.readpermissions.access_grant.revokepermissions.audit_logs.readpermissions.finance_report.readpermissions.inway.readpermissions.inway.updatepermissions.inway.deletepermissions.inways.readpermissions.outgoing_order.createpermissions.outgoing_order.updatepermissions.outgoing_order.revokepermissions.outgoing_orders.readpermissions.incoming_orders.readpermissions.incoming_orders.synchronizepermissions.outways.readpermissions.service.createpermissions.service.readpermissions.service.updatepermissions.service.deletepermissions.services.readpermissions.services_statistics.readpermissions.organization_settings.readpermissions.organization_settings.updatepermissions.terms_of_service.acceptpermissions.terms_of_service_status.readpermissions.transaction_logs.read"

var _PermissionIndex = [...]uint16{0, 43, 85, 126, 168, 210, 240, 271, 298, 329, 351, 375, 399, 422, 455, 488, 521, 553, 585, 624, 648, 674, 698, 724, 750, 775, 811, 849, 889, 924, 964, 997}

const _PermissionLowerName = "permissions.incoming_access_request.approvepermissions.incoming_access_request.rejectpermissions.incoming_access_requests.readpermissions.outgoing_access_request.createpermissions.outgoing_access_request.updatepermissions.access_grants.readpermissions.access_grant.revokepermissions.audit_logs.readpermissions.finance_report.readpermissions.inway.readpermissions.inway.updatepermissions.inway.deletepermissions.inways.readpermissions.outgoing_order.createpermissions.outgoing_order.updatepermissions.outgoing_order.revokepermissions.outgoing_orders.readpermissions.incoming_orders.readpermissions.incoming_orders.synchronizepermissions.outways.readpermissions.service.createpermissions.service.readpermissions.service.updatepermissions.service.deletepermissions.services.readpermissions.services_statistics.readpermissions.organization_settings.readpermissions.organization_settings.updatepermissions.terms_of_service.acceptpermissions.terms_of_service_status.readpermissions.transaction_logs.read"

func (i Permission) String() string {
	i -= 1
	if i < 0 || i >= Permission(len(_PermissionIndex)-1) {
		return fmt.Sprintf("Permission(%d)", i+1)
	}
	return _PermissionName[_PermissionIndex[i]:_PermissionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PermissionNoOp() {
	var x [1]struct{}
	_ = x[ApproveIncomingAccessRequest-(1)]
	_ = x[RejectIncomingAccessRequest-(2)]
	_ = x[ReadIncomingAccessRequests-(3)]
	_ = x[CreateOutgoingAccessRequest-(4)]
	_ = x[UpdateOutgoingAccessRequest-(5)]
	_ = x[ReadAccessGrants-(6)]
	_ = x[RevokeAccessGrant-(7)]
	_ = x[ReadAuditLogs-(8)]
	_ = x[ReadFinanceReport-(9)]
	_ = x[ReadInway-(10)]
	_ = x[UpdateInway-(11)]
	_ = x[DeleteInway-(12)]
	_ = x[ReadInways-(13)]
	_ = x[CreateOutgoingOrder-(14)]
	_ = x[UpdateOutgoingOrder-(15)]
	_ = x[RevokeOutgoingOrder-(16)]
	_ = x[ReadOutgoingOrders-(17)]
	_ = x[ReadIncomingOrders-(18)]
	_ = x[SynchronizeIncomingOrders-(19)]
	_ = x[ReadOutways-(20)]
	_ = x[CreateService-(21)]
	_ = x[ReadService-(22)]
	_ = x[UpdateService-(23)]
	_ = x[DeleteService-(24)]
	_ = x[ReadServices-(25)]
	_ = x[ReadServicesStatistics-(26)]
	_ = x[ReadOrganizationSettings-(27)]
	_ = x[UpdateOrganizationSettings-(28)]
	_ = x[AcceptTermsOfService-(29)]
	_ = x[ReadTermsOfServiceStatus-(30)]
	_ = x[ReadTransactionLogs-(31)]
}

var _PermissionValues = []Permission{ApproveIncomingAccessRequest, RejectIncomingAccessRequest, ReadIncomingAccessRequests, CreateOutgoingAccessRequest, UpdateOutgoingAccessRequest, ReadAccessGrants, RevokeAccessGrant, ReadAuditLogs, ReadFinanceReport, ReadInway, UpdateInway, DeleteInway, ReadInways, CreateOutgoingOrder, UpdateOutgoingOrder, RevokeOutgoingOrder, ReadOutgoingOrders, ReadIncomingOrders, SynchronizeIncomingOrders, ReadOutways, CreateService, ReadService, UpdateService, DeleteService, ReadServices, ReadServicesStatistics, ReadOrganizationSettings, UpdateOrganizationSettings, AcceptTermsOfService, ReadTermsOfServiceStatus, ReadTransactionLogs}

var _PermissionNameToValueMap = map[string]Permission{
	_PermissionName[0:43]:         ApproveIncomingAccessRequest,
	_PermissionLowerName[0:43]:    ApproveIncomingAccessRequest,
	_PermissionName[43:85]:        RejectIncomingAccessRequest,
	_PermissionLowerName[43:85]:   RejectIncomingAccessRequest,
	_PermissionName[85:126]:       ReadIncomingAccessRequests,
	_PermissionLowerName[85:126]:  ReadIncomingAccessRequests,
	_PermissionName[126:168]:      CreateOutgoingAccessRequest,
	_PermissionLowerName[126:168]: CreateOutgoingAccessRequest,
	_PermissionName[168:210]:      UpdateOutgoingAccessRequest,
	_PermissionLowerName[168:210]: UpdateOutgoingAccessRequest,
	_PermissionName[210:240]:      ReadAccessGrants,
	_PermissionLowerName[210:240]: ReadAccessGrants,
	_PermissionName[240:271]:      RevokeAccessGrant,
	_PermissionLowerName[240:271]: RevokeAccessGrant,
	_PermissionName[271:298]:      ReadAuditLogs,
	_PermissionLowerName[271:298]: ReadAuditLogs,
	_PermissionName[298:329]:      ReadFinanceReport,
	_PermissionLowerName[298:329]: ReadFinanceReport,
	_PermissionName[329:351]:      ReadInway,
	_PermissionLowerName[329:351]: ReadInway,
	_PermissionName[351:375]:      UpdateInway,
	_PermissionLowerName[351:375]: UpdateInway,
	_PermissionName[375:399]:      DeleteInway,
	_PermissionLowerName[375:399]: DeleteInway,
	_PermissionName[399:422]:      ReadInways,
	_PermissionLowerName[399:422]: ReadInways,
	_PermissionName[422:455]:      CreateOutgoingOrder,
	_PermissionLowerName[422:455]: CreateOutgoingOrder,
	_PermissionName[455:488]:      UpdateOutgoingOrder,
	_PermissionLowerName[455:488]: UpdateOutgoingOrder,
	_PermissionName[488:521]:      RevokeOutgoingOrder,
	_PermissionLowerName[488:521]: RevokeOutgoingOrder,
	_PermissionName[521:553]:      ReadOutgoingOrders,
	_PermissionLowerName[521:553]: ReadOutgoingOrders,
	_PermissionName[553:585]:      ReadIncomingOrders,
	_PermissionLowerName[553:585]: ReadIncomingOrders,
	_PermissionName[585:624]:      SynchronizeIncomingOrders,
	_PermissionLowerName[585:624]: SynchronizeIncomingOrders,
	_PermissionName[624:648]:      ReadOutways,
	_PermissionLowerName[624:648]: ReadOutways,
	_PermissionName[648:674]:      CreateService,
	_PermissionLowerName[648:674]: CreateService,
	_PermissionName[674:698]:      ReadService,
	_PermissionLowerName[674:698]: ReadService,
	_PermissionName[698:724]:      UpdateService,
	_PermissionLowerName[698:724]: UpdateService,
	_PermissionName[724:750]:      DeleteService,
	_PermissionLowerName[724:750]: DeleteService,
	_PermissionName[750:775]:      ReadServices,
	_PermissionLowerName[750:775]: ReadServices,
	_PermissionName[775:811]:      ReadServicesStatistics,
	_PermissionLowerName[775:811]: ReadServicesStatistics,
	_PermissionName[811:849]:      ReadOrganizationSettings,
	_PermissionLowerName[811:849]: ReadOrganizationSettings,
	_PermissionName[849:889]:      UpdateOrganizationSettings,
	_PermissionLowerName[849:889]: UpdateOrganizationSettings,
	_PermissionName[889:924]:      AcceptTermsOfService,
	_PermissionLowerName[889:924]: AcceptTermsOfService,
	_PermissionName[924:964]:      ReadTermsOfServiceStatus,
	_PermissionLowerName[924:964]: ReadTermsOfServiceStatus,
	_PermissionName[964:997]:      ReadTransactionLogs,
	_PermissionLowerName[964:997]: ReadTransactionLogs,
}

var _PermissionNames = []string{
	_PermissionName[0:43],
	_PermissionName[43:85],
	_PermissionName[85:126],
	_PermissionName[126:168],
	_PermissionName[168:210],
	_PermissionName[210:240],
	_PermissionName[240:271],
	_PermissionName[271:298],
	_PermissionName[298:329],
	_PermissionName[329:351],
	_PermissionName[351:375],
	_PermissionName[375:399],
	_PermissionName[399:422],
	_PermissionName[422:455],
	_PermissionName[455:488],
	_PermissionName[488:521],
	_PermissionName[521:553],
	_PermissionName[553:585],
	_PermissionName[585:624],
	_PermissionName[624:648],
	_PermissionName[648:674],
	_PermissionName[674:698],
	_PermissionName[698:724],
	_PermissionName[724:750],
	_PermissionName[750:775],
	_PermissionName[775:811],
	_PermissionName[811:849],
	_PermissionName[849:889],
	_PermissionName[889:924],
	_PermissionName[924:964],
	_PermissionName[964:997],
}

// PermissionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PermissionString(s string) (Permission, error) {
	if val, ok := _PermissionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PermissionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Permission values", s)
}

// PermissionValues returns all values of the enum
func PermissionValues() []Permission {
	return _PermissionValues
}

// PermissionStrings returns a slice of all String values of the enum
func PermissionStrings() []string {
	strs := make([]string, len(_PermissionNames))
	copy(strs, _PermissionNames)
	return strs
}

// IsAPermission returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Permission) IsAPermission() bool {
	for _, v := range _PermissionValues {
		if i == v {
			return true
		}
	}
	return false
}
