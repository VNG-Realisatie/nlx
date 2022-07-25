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
 * @interface ManagementOrderService
 */
export interface ManagementOrderService {
    /**
     * 
     * @type {ManagementOrganization}
     * @memberof ManagementOrderService
     */
    organization?: ManagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementOrderService
     */
    service?: string;
}

export function ManagementOrderServiceFromJSON(json: any): ManagementOrderService {
    return ManagementOrderServiceFromJSONTyped(json, false);
}

export function ManagementOrderServiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementOrderService {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organization': !exists(json, 'organization') ? undefined : ManagementOrganizationFromJSON(json['organization']),
        'service': !exists(json, 'service') ? undefined : json['service'],
    };
}

export function ManagementOrderServiceToJSON(value?: ManagementOrderService | null): any {
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

