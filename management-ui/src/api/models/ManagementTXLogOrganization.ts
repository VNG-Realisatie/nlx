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
/**
 * 
 * @export
 * @interface ManagementTXLogOrganization
 */
export interface ManagementTXLogOrganization {
    /**
     * 
     * @type {string}
     * @memberof ManagementTXLogOrganization
     */
    serialNumber?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementTXLogOrganization
     */
    name?: string;
}

export function ManagementTXLogOrganizationFromJSON(json: any): ManagementTXLogOrganization {
    return ManagementTXLogOrganizationFromJSONTyped(json, false);
}

export function ManagementTXLogOrganizationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementTXLogOrganization {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'serialNumber': !exists(json, 'serialNumber') ? undefined : json['serialNumber'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function ManagementTXLogOrganizationToJSON(value?: ManagementTXLogOrganization | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'serialNumber': value.serialNumber,
        'name': value.name,
    };
}

