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
    LOGINSUCCESS = 'loginSuccess',
    LOGINFAIL = 'loginFail',
    LOGOUT = 'logout',
    INCOMINGACCESSREQUESTACCEPT = 'incomingAccessRequestAccept',
    INCOMINGACCESSREQUESTREJECT = 'incomingAccessRequestReject',
    ACCESSGRANTREVOKE = 'accessGrantRevoke',
    OUTGOINGACCESSREQUESTCREATE = 'outgoingAccessRequestCreate',
    OUTGOINGACCESSREQUESTFAIL = 'outgoingAccessRequestFail',
    SERVICECREATE = 'serviceCreate',
    SERVICEUPDATE = 'serviceUpdate',
    SERVICEDELETE = 'serviceDelete',
    ORGANIZATIONSETTINGSUPDATE = 'organizationSettingsUpdate',
    ORDERCREATE = 'orderCreate',
    ORDEROUTGOINGREVOKE = 'orderOutgoingRevoke',
    ORDERINCOMINGREVOKE = 'orderIncomingRevoke'
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

