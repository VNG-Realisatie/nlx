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
    user?: string;
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
    organization?: string;
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
        'user': !exists(json, 'user') ? undefined : json['user'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'organization': !exists(json, 'organization') ? undefined : json['organization'],
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
        'user': value.user,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'organization': value.organization,
    };
}


