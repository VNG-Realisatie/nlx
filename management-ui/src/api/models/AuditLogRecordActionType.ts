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

/**
 * 
 * @export
 * @enum {string}
 */
export enum AuditLogRecordActionType {
    LOGIN_SUCCESS = 'loginSuccess',
    LOGIN_FAIL = 'loginFail',
    LOGOUT = 'logout',
    INCOMING_ACCESS_REQUEST_ACCEPT = 'incomingAccessRequestAccept',
    INCOMING_ACCESS_REQUEST_REJECT = 'incomingAccessRequestReject',
    ACCESS_GRANT_REVOKE = 'accessGrantRevoke',
    OUTGOING_ACCESS_REQUEST_CREATE = 'outgoingAccessRequestCreate',
    OUTGOING_ACCESS_REQUEST_FAIL = 'outgoingAccessRequestFail',
    SERVICE_CREATE = 'serviceCreate',
    SERVICE_UPDATE = 'serviceUpdate',
    SERVICE_DELETE = 'serviceDelete',
    ORGANIZATION_SETTINGS_UPDATE = 'organizationSettingsUpdate',
    ORDER_CREATE = 'orderCreate',
    ORDER_OUTGOING_REVOKE = 'orderOutgoingRevoke',
    ORDER_INCOMING_REVOKE = 'orderIncomingRevoke',
    INWAY_DELETE = 'inwayDelete',
    ORDER_OUTGOING_UPDATE = 'orderOutgoingUpdate'
}

export function AuditLogRecordActionTypeFromJSON(json: any): AuditLogRecordActionType {
    return AuditLogRecordActionTypeFromJSONTyped(json, false);
}

export function AuditLogRecordActionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuditLogRecordActionType {
    return json as AuditLogRecordActionType;
}

export function AuditLogRecordActionTypeToJSON(value?: AuditLogRecordActionType | null): any {
    return value as any;
}

