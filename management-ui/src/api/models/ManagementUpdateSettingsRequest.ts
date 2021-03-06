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
 * @interface ManagementUpdateSettingsRequest
 */
export interface ManagementUpdateSettingsRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateSettingsRequest
     */
    organizationInway?: string;
}

export function ManagementUpdateSettingsRequestFromJSON(json: any): ManagementUpdateSettingsRequest {
    return ManagementUpdateSettingsRequestFromJSONTyped(json, false);
}

export function ManagementUpdateSettingsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementUpdateSettingsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organizationInway': !exists(json, 'organizationInway') ? undefined : json['organizationInway'],
    };
}

export function ManagementUpdateSettingsRequestToJSON(value?: ManagementUpdateSettingsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'organizationInway': value.organizationInway,
    };
}


