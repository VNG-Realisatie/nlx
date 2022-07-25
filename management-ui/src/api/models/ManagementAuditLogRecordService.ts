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
    ManagementOrganization,
    ManagementOrganizationFromJSON,
    ManagementOrganizationFromJSONTyped,
    ManagementOrganizationToJSON,
} from './ManagementOrganization';

/**
 * 
 * @export
 * @interface ManagementAuditLogRecordService
 */
export interface ManagementAuditLogRecordService {
    /**
     * 
     * @type {ManagementOrganization}
     * @memberof ManagementAuditLogRecordService
     */
    organization?: ManagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecordService
     */
    service?: string;
}

export function ManagementAuditLogRecordServiceFromJSON(json: any): ManagementAuditLogRecordService {
    return ManagementAuditLogRecordServiceFromJSONTyped(json, false);
}

export function ManagementAuditLogRecordServiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementAuditLogRecordService {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organization': !exists(json, 'organization') ? undefined : ManagementOrganizationFromJSON(json['organization']),
        'service': !exists(json, 'service') ? undefined : json['service'],
    };
}

export function ManagementAuditLogRecordServiceToJSON(value?: ManagementAuditLogRecordService | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'organization': ManagementOrganizationToJSON(value.organization),
        'service': value.service,
    };
}

