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
 * @interface ManagementIsInwayOrganizationInwayResponse
 */
export interface ManagementIsInwayOrganizationInwayResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ManagementIsInwayOrganizationInwayResponse
     */
    isOrganizationInway?: boolean;
}

export function ManagementIsInwayOrganizationInwayResponseFromJSON(json: any): ManagementIsInwayOrganizationInwayResponse {
    return ManagementIsInwayOrganizationInwayResponseFromJSONTyped(json, false);
}

export function ManagementIsInwayOrganizationInwayResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementIsInwayOrganizationInwayResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'isOrganizationInway': !exists(json, 'isOrganizationInway') ? undefined : json['isOrganizationInway'],
    };
}

export function ManagementIsInwayOrganizationInwayResponseToJSON(value?: ManagementIsInwayOrganizationInwayResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'isOrganizationInway': value.isOrganizationInway,
    };
}

