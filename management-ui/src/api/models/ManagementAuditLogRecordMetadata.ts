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
    NlxmanagementOrganization,
    NlxmanagementOrganizationFromJSON,
    NlxmanagementOrganizationFromJSONTyped,
    NlxmanagementOrganizationToJSON,
} from './NlxmanagementOrganization';

/**
 * 
 * @export
 * @interface ManagementAuditLogRecordMetadata
 */
export interface ManagementAuditLogRecordMetadata {
    /**
     * 
     * @type {NlxmanagementOrganization}
     * @memberof ManagementAuditLogRecordMetadata
     */
    delegatee?: NlxmanagementOrganization;
    /**
     * 
     * @type {NlxmanagementOrganization}
     * @memberof ManagementAuditLogRecordMetadata
     */
    delegator?: NlxmanagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecordMetadata
     */
    reference?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecordMetadata
     */
    inwayName?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAuditLogRecordMetadata
     */
    outwayName?: string;
}

export function ManagementAuditLogRecordMetadataFromJSON(json: any): ManagementAuditLogRecordMetadata {
    return ManagementAuditLogRecordMetadataFromJSONTyped(json, false);
}

export function ManagementAuditLogRecordMetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementAuditLogRecordMetadata {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'delegatee': !exists(json, 'delegatee') ? undefined : NlxmanagementOrganizationFromJSON(json['delegatee']),
        'delegator': !exists(json, 'delegator') ? undefined : NlxmanagementOrganizationFromJSON(json['delegator']),
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'inwayName': !exists(json, 'inway_name') ? undefined : json['inway_name'],
        'outwayName': !exists(json, 'outway_name') ? undefined : json['outway_name'],
    };
}

export function ManagementAuditLogRecordMetadataToJSON(value?: ManagementAuditLogRecordMetadata | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'delegatee': NlxmanagementOrganizationToJSON(value.delegatee),
        'delegator': NlxmanagementOrganizationToJSON(value.delegator),
        'reference': value.reference,
        'inway_name': value.inwayName,
        'outway_name': value.outwayName,
    };
}

