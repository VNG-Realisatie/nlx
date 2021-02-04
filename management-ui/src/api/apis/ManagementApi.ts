/* tslint:disable */
/* eslint-disable */
/**
 * management.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    ManagementAccessGrant,
    ManagementAccessGrantFromJSON,
    ManagementAccessGrantToJSON,
    ManagementCreateAccessRequestRequest,
    ManagementCreateAccessRequestRequestFromJSON,
    ManagementCreateAccessRequestRequestToJSON,
    ManagementCreateServiceRequest,
    ManagementCreateServiceRequestFromJSON,
    ManagementCreateServiceRequestToJSON,
    ManagementCreateServiceResponse,
    ManagementCreateServiceResponseFromJSON,
    ManagementCreateServiceResponseToJSON,
    ManagementGetServiceResponse,
    ManagementGetServiceResponseFromJSON,
    ManagementGetServiceResponseToJSON,
    ManagementGetStatisticsOfServicesResponse,
    ManagementGetStatisticsOfServicesResponseFromJSON,
    ManagementGetStatisticsOfServicesResponseToJSON,
    ManagementInsightConfiguration,
    ManagementInsightConfigurationFromJSON,
    ManagementInsightConfigurationToJSON,
    ManagementInway,
    ManagementInwayFromJSON,
    ManagementInwayToJSON,
    ManagementListAccessGrantsForServiceResponse,
    ManagementListAccessGrantsForServiceResponseFromJSON,
    ManagementListAccessGrantsForServiceResponseToJSON,
    ManagementListAuditLogsResponse,
    ManagementListAuditLogsResponseFromJSON,
    ManagementListAuditLogsResponseToJSON,
    ManagementListIncomingAccessRequestsResponse,
    ManagementListIncomingAccessRequestsResponseFromJSON,
    ManagementListIncomingAccessRequestsResponseToJSON,
    ManagementListInwaysResponse,
    ManagementListInwaysResponseFromJSON,
    ManagementListInwaysResponseToJSON,
    ManagementListOutgoingAccessRequestsResponse,
    ManagementListOutgoingAccessRequestsResponseFromJSON,
    ManagementListOutgoingAccessRequestsResponseToJSON,
    ManagementListServicesResponse,
    ManagementListServicesResponseFromJSON,
    ManagementListServicesResponseToJSON,
    ManagementOutgoingAccessRequest,
    ManagementOutgoingAccessRequestFromJSON,
    ManagementOutgoingAccessRequestToJSON,
    ManagementSettings,
    ManagementSettingsFromJSON,
    ManagementSettingsToJSON,
    ManagementUpdateServiceRequest,
    ManagementUpdateServiceRequestFromJSON,
    ManagementUpdateServiceRequestToJSON,
    ManagementUpdateServiceResponse,
    ManagementUpdateServiceResponseFromJSON,
    ManagementUpdateServiceResponseToJSON,
    ManagementUpdateSettingsRequest,
    ManagementUpdateSettingsRequestFromJSON,
    ManagementUpdateSettingsRequestToJSON,
    RuntimeError,
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
} from '../models';

export interface ManagementApproveIncomingAccessRequestRequest {
    serviceName: string;
    accessRequestID: string;
}

export interface ManagementCreateAccessRequestOperationRequest {
    body: ManagementCreateAccessRequestRequest;
}

export interface ManagementCreateInwayRequest {
    body: ManagementInway;
}

export interface ManagementCreateServiceOperationRequest {
    body: ManagementCreateServiceRequest;
}

export interface ManagementDeleteInwayRequest {
    name: string;
}

export interface ManagementDeleteServiceRequest {
    name: string;
}

export interface ManagementGetInwayRequest {
    name: string;
}

export interface ManagementGetServiceRequest {
    name: string;
}

export interface ManagementListAccessGrantsForServiceRequest {
    serviceName: string;
}

export interface ManagementListIncomingAccessRequestRequest {
    serviceName: string;
}

export interface ManagementListOutgoingAccessRequestsRequest {
    organizationName: string;
    serviceName: string;
}

export interface ManagementListServicesRequest {
    inwayName?: string;
}

export interface ManagementPutInsightConfigurationRequest {
    body: ManagementInsightConfiguration;
}

export interface ManagementRejectIncomingAccessRequestRequest {
    serviceName: string;
    accessRequestID: string;
}

export interface ManagementRevokeAccessGrantRequest {
    serviceName: string;
    organizationName: string;
    accessGrantID: string;
}

export interface ManagementSendAccessRequestRequest {
    organizationName: string;
    serviceName: string;
    accessRequestID: string;
}

export interface ManagementUpdateInwayRequest {
    name: string;
    body: ManagementInway;
}

export interface ManagementUpdateServiceOperationRequest {
    name: string;
    body: ManagementUpdateServiceRequest;
}

export interface ManagementUpdateSettingsOperationRequest {
    body: ManagementUpdateSettingsRequest;
}

/**
 * 
 */
export class ManagementApi extends runtime.BaseAPI {

    /**
     */
    async managementApproveIncomingAccessRequestRaw(requestParameters: ManagementApproveIncomingAccessRequestRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementApproveIncomingAccessRequest.');
        }

        if (requestParameters.accessRequestID === null || requestParameters.accessRequestID === undefined) {
            throw new runtime.RequiredError('accessRequestID','Required parameter requestParameters.accessRequestID was null or undefined when calling managementApproveIncomingAccessRequest.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-requests/incoming/services/{serviceName}/{accessRequestID}/approve`.replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))).replace(`{${"accessRequestID"}}`, encodeURIComponent(String(requestParameters.accessRequestID))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async managementApproveIncomingAccessRequest(requestParameters: ManagementApproveIncomingAccessRequestRequest): Promise<object> {
        const response = await this.managementApproveIncomingAccessRequestRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementCreateAccessRequestRaw(requestParameters: ManagementCreateAccessRequestOperationRequest): Promise<runtime.ApiResponse<ManagementOutgoingAccessRequest>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementCreateAccessRequest.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/access-requests`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementCreateAccessRequestRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementOutgoingAccessRequestFromJSON(jsonValue));
    }

    /**
     */
    async managementCreateAccessRequest(requestParameters: ManagementCreateAccessRequestOperationRequest): Promise<ManagementOutgoingAccessRequest> {
        const response = await this.managementCreateAccessRequestRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementCreateInwayRaw(requestParameters: ManagementCreateInwayRequest): Promise<runtime.ApiResponse<ManagementInway>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementCreateInway.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/inways`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementInwayToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementInwayFromJSON(jsonValue));
    }

    /**
     */
    async managementCreateInway(requestParameters: ManagementCreateInwayRequest): Promise<ManagementInway> {
        const response = await this.managementCreateInwayRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementCreateServiceRaw(requestParameters: ManagementCreateServiceOperationRequest): Promise<runtime.ApiResponse<ManagementCreateServiceResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementCreateService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/services`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementCreateServiceRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementCreateServiceResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementCreateService(requestParameters: ManagementCreateServiceOperationRequest): Promise<ManagementCreateServiceResponse> {
        const response = await this.managementCreateServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementDeleteInwayRaw(requestParameters: ManagementDeleteInwayRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementDeleteInway.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/inways/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async managementDeleteInway(requestParameters: ManagementDeleteInwayRequest): Promise<object> {
        const response = await this.managementDeleteInwayRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementDeleteServiceRaw(requestParameters: ManagementDeleteServiceRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementDeleteService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/services/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async managementDeleteService(requestParameters: ManagementDeleteServiceRequest): Promise<object> {
        const response = await this.managementDeleteServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementGetInsightConfigurationRaw(): Promise<runtime.ApiResponse<ManagementInsightConfiguration>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/insight-configuration`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementInsightConfigurationFromJSON(jsonValue));
    }

    /**
     */
    async managementGetInsightConfiguration(): Promise<ManagementInsightConfiguration> {
        const response = await this.managementGetInsightConfigurationRaw();
        return await response.value();
    }

    /**
     */
    async managementGetInwayRaw(requestParameters: ManagementGetInwayRequest): Promise<runtime.ApiResponse<ManagementInway>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementGetInway.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/inways/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementInwayFromJSON(jsonValue));
    }

    /**
     */
    async managementGetInway(requestParameters: ManagementGetInwayRequest): Promise<ManagementInway> {
        const response = await this.managementGetInwayRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementGetServiceRaw(requestParameters: ManagementGetServiceRequest): Promise<runtime.ApiResponse<ManagementGetServiceResponse>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementGetService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/services/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementGetServiceResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementGetService(requestParameters: ManagementGetServiceRequest): Promise<ManagementGetServiceResponse> {
        const response = await this.managementGetServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementGetSettingsRaw(): Promise<runtime.ApiResponse<ManagementSettings>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/settings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementSettingsFromJSON(jsonValue));
    }

    /**
     */
    async managementGetSettings(): Promise<ManagementSettings> {
        const response = await this.managementGetSettingsRaw();
        return await response.value();
    }

    /**
     */
    async managementGetStatisticsOfServicesRaw(): Promise<runtime.ApiResponse<ManagementGetStatisticsOfServicesResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/statistics/services`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementGetStatisticsOfServicesResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementGetStatisticsOfServices(): Promise<ManagementGetStatisticsOfServicesResponse> {
        const response = await this.managementGetStatisticsOfServicesRaw();
        return await response.value();
    }

    /**
     */
    async managementListAccessGrantsForServiceRaw(requestParameters: ManagementListAccessGrantsForServiceRequest): Promise<runtime.ApiResponse<ManagementListAccessGrantsForServiceResponse>> {
        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementListAccessGrantsForService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-grants/services/{serviceName}`.replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListAccessGrantsForServiceResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListAccessGrantsForService(requestParameters: ManagementListAccessGrantsForServiceRequest): Promise<ManagementListAccessGrantsForServiceResponse> {
        const response = await this.managementListAccessGrantsForServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementListAuditLogsRaw(): Promise<runtime.ApiResponse<ManagementListAuditLogsResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/audit-logs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListAuditLogsResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListAuditLogs(): Promise<ManagementListAuditLogsResponse> {
        const response = await this.managementListAuditLogsRaw();
        return await response.value();
    }

    /**
     */
    async managementListIncomingAccessRequestRaw(requestParameters: ManagementListIncomingAccessRequestRequest): Promise<runtime.ApiResponse<ManagementListIncomingAccessRequestsResponse>> {
        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementListIncomingAccessRequest.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-requests/incoming/services/{serviceName}`.replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListIncomingAccessRequestsResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListIncomingAccessRequest(requestParameters: ManagementListIncomingAccessRequestRequest): Promise<ManagementListIncomingAccessRequestsResponse> {
        const response = await this.managementListIncomingAccessRequestRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementListInwaysRaw(): Promise<runtime.ApiResponse<ManagementListInwaysResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/inways`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListInwaysResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListInways(): Promise<ManagementListInwaysResponse> {
        const response = await this.managementListInwaysRaw();
        return await response.value();
    }

    /**
     */
    async managementListOutgoingAccessRequestsRaw(requestParameters: ManagementListOutgoingAccessRequestsRequest): Promise<runtime.ApiResponse<ManagementListOutgoingAccessRequestsResponse>> {
        if (requestParameters.organizationName === null || requestParameters.organizationName === undefined) {
            throw new runtime.RequiredError('organizationName','Required parameter requestParameters.organizationName was null or undefined when calling managementListOutgoingAccessRequests.');
        }

        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementListOutgoingAccessRequests.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-requests/outgoing/organizations/{organizationName}/services/{serviceName}`.replace(`{${"organizationName"}}`, encodeURIComponent(String(requestParameters.organizationName))).replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListOutgoingAccessRequestsResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListOutgoingAccessRequests(requestParameters: ManagementListOutgoingAccessRequestsRequest): Promise<ManagementListOutgoingAccessRequestsResponse> {
        const response = await this.managementListOutgoingAccessRequestsRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementListServicesRaw(requestParameters: ManagementListServicesRequest): Promise<runtime.ApiResponse<ManagementListServicesResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.inwayName !== undefined) {
            queryParameters['inwayName'] = requestParameters.inwayName;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/services`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementListServicesResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementListServices(requestParameters: ManagementListServicesRequest): Promise<ManagementListServicesResponse> {
        const response = await this.managementListServicesRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementPutInsightConfigurationRaw(requestParameters: ManagementPutInsightConfigurationRequest): Promise<runtime.ApiResponse<ManagementInsightConfiguration>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementPutInsightConfiguration.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/insight-configuration`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementInsightConfigurationToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementInsightConfigurationFromJSON(jsonValue));
    }

    /**
     */
    async managementPutInsightConfiguration(requestParameters: ManagementPutInsightConfigurationRequest): Promise<ManagementInsightConfiguration> {
        const response = await this.managementPutInsightConfigurationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementRejectIncomingAccessRequestRaw(requestParameters: ManagementRejectIncomingAccessRequestRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementRejectIncomingAccessRequest.');
        }

        if (requestParameters.accessRequestID === null || requestParameters.accessRequestID === undefined) {
            throw new runtime.RequiredError('accessRequestID','Required parameter requestParameters.accessRequestID was null or undefined when calling managementRejectIncomingAccessRequest.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-requests/incoming/services/{serviceName}/{accessRequestID}/reject`.replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))).replace(`{${"accessRequestID"}}`, encodeURIComponent(String(requestParameters.accessRequestID))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async managementRejectIncomingAccessRequest(requestParameters: ManagementRejectIncomingAccessRequestRequest): Promise<object> {
        const response = await this.managementRejectIncomingAccessRequestRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementRevokeAccessGrantRaw(requestParameters: ManagementRevokeAccessGrantRequest): Promise<runtime.ApiResponse<ManagementAccessGrant>> {
        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementRevokeAccessGrant.');
        }

        if (requestParameters.organizationName === null || requestParameters.organizationName === undefined) {
            throw new runtime.RequiredError('organizationName','Required parameter requestParameters.organizationName was null or undefined when calling managementRevokeAccessGrant.');
        }

        if (requestParameters.accessGrantID === null || requestParameters.accessGrantID === undefined) {
            throw new runtime.RequiredError('accessGrantID','Required parameter requestParameters.accessGrantID was null or undefined when calling managementRevokeAccessGrant.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-grants/service/{serviceName}/organizations/{organizationName}/{accessGrantID}/revoke`.replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))).replace(`{${"organizationName"}}`, encodeURIComponent(String(requestParameters.organizationName))).replace(`{${"accessGrantID"}}`, encodeURIComponent(String(requestParameters.accessGrantID))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementAccessGrantFromJSON(jsonValue));
    }

    /**
     */
    async managementRevokeAccessGrant(requestParameters: ManagementRevokeAccessGrantRequest): Promise<ManagementAccessGrant> {
        const response = await this.managementRevokeAccessGrantRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementSendAccessRequestRaw(requestParameters: ManagementSendAccessRequestRequest): Promise<runtime.ApiResponse<ManagementOutgoingAccessRequest>> {
        if (requestParameters.organizationName === null || requestParameters.organizationName === undefined) {
            throw new runtime.RequiredError('organizationName','Required parameter requestParameters.organizationName was null or undefined when calling managementSendAccessRequest.');
        }

        if (requestParameters.serviceName === null || requestParameters.serviceName === undefined) {
            throw new runtime.RequiredError('serviceName','Required parameter requestParameters.serviceName was null or undefined when calling managementSendAccessRequest.');
        }

        if (requestParameters.accessRequestID === null || requestParameters.accessRequestID === undefined) {
            throw new runtime.RequiredError('accessRequestID','Required parameter requestParameters.accessRequestID was null or undefined when calling managementSendAccessRequest.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/access-requests/outgoing/organizations/{organizationName}/services/{serviceName}/{accessRequestID}/send`.replace(`{${"organizationName"}}`, encodeURIComponent(String(requestParameters.organizationName))).replace(`{${"serviceName"}}`, encodeURIComponent(String(requestParameters.serviceName))).replace(`{${"accessRequestID"}}`, encodeURIComponent(String(requestParameters.accessRequestID))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementOutgoingAccessRequestFromJSON(jsonValue));
    }

    /**
     */
    async managementSendAccessRequest(requestParameters: ManagementSendAccessRequestRequest): Promise<ManagementOutgoingAccessRequest> {
        const response = await this.managementSendAccessRequestRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementUpdateInwayRaw(requestParameters: ManagementUpdateInwayRequest): Promise<runtime.ApiResponse<ManagementInway>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementUpdateInway.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementUpdateInway.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/inways/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementInwayToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementInwayFromJSON(jsonValue));
    }

    /**
     */
    async managementUpdateInway(requestParameters: ManagementUpdateInwayRequest): Promise<ManagementInway> {
        const response = await this.managementUpdateInwayRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementUpdateServiceRaw(requestParameters: ManagementUpdateServiceOperationRequest): Promise<runtime.ApiResponse<ManagementUpdateServiceResponse>> {
        if (requestParameters.name === null || requestParameters.name === undefined) {
            throw new runtime.RequiredError('name','Required parameter requestParameters.name was null or undefined when calling managementUpdateService.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementUpdateService.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/services/{name}`.replace(`{${"name"}}`, encodeURIComponent(String(requestParameters.name))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementUpdateServiceRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementUpdateServiceResponseFromJSON(jsonValue));
    }

    /**
     */
    async managementUpdateService(requestParameters: ManagementUpdateServiceOperationRequest): Promise<ManagementUpdateServiceResponse> {
        const response = await this.managementUpdateServiceRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async managementUpdateSettingsRaw(requestParameters: ManagementUpdateSettingsOperationRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling managementUpdateSettings.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v1/settings`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: ManagementUpdateSettingsRequestToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async managementUpdateSettings(requestParameters: ManagementUpdateSettingsOperationRequest): Promise<object> {
        const response = await this.managementUpdateSettingsRaw(requestParameters);
        return await response.value();
    }

}
