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

import { exists, mapValues } from '../runtime';
import {
    AuditLogRecordActionType,
    AuditLogRecordActionTypeFromJSON,
    AuditLogRecordActionTypeFromJSONTyped,
    AuditLogRecordActionTypeToJSON,
    ManagementAuditLogRecordService,
    ManagementAuditLogRecordServiceFromJSON,
    ManagementAuditLogRecordServiceFromJSONTyped,
    ManagementAuditLogRecordServiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementAuditLogRecord
 */
export interface ManagementAuditLogRecord {
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    id?: string;
    /**
     * 
     * @type {AuditLogRecordActionType}
     * @memberof ManagementAuditLogRecord
     */
    action?: AuditLogRecordActionType;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    operatingSystem?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    browser?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    client?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    user?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    delegatee?: string;
    /**
     * 
     * @type {Array<ManagementAuditLogRecordService>}
     * @memberof ManagementAuditLogRecord
     */
    services?: Array<ManagementAuditLogRecordService>;
    /**
     * 
     * @type {Date}
     * @memberof ManagementAuditLogRecord
     */
    createdAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecord
     */
    data?: string;
}

export function ManagementAuditLogRecordFromJSON(json: any): ManagementAuditLogRecord {
    return ManagementAuditLogRecordFromJSONTyped(json, false);
}

export function ManagementAuditLogRecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementAuditLogRecord {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'action': !exists(json, 'action') ? undefined : AuditLogRecordActionTypeFromJSON(json['action']),
        'operatingSystem': !exists(json, 'operatingSystem') ? undefined : json['operatingSystem'],
        'browser': !exists(json, 'browser') ? undefined : json['browser'],
        'client': !exists(json, 'client') ? undefined : json['client'],
        'user': !exists(json, 'user') ? undefined : json['user'],
        'delegatee': !exists(json, 'delegatee') ? undefined : json['delegatee'],
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementAuditLogRecordServiceFromJSON)),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'data': !exists(json, 'data') ? undefined : json['data'],
    };
}

export function ManagementAuditLogRecordToJSON(value?: ManagementAuditLogRecord | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'action': AuditLogRecordActionTypeToJSON(value.action),
        'operatingSystem': value.operatingSystem,
        'browser': value.browser,
        'client': value.client,
        'user': value.user,
        'delegatee': value.delegatee,
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementAuditLogRecordServiceToJSON)),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'data': value.data,
    };
}


