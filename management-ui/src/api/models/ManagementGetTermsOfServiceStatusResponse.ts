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
 * @interface ManagementGetTermsOfServiceStatusResponse
 */
export interface ManagementGetTermsOfServiceStatusResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ManagementGetTermsOfServiceStatusResponse
     */
    accepted?: boolean;
}

export function ManagementGetTermsOfServiceStatusResponseFromJSON(json: any): ManagementGetTermsOfServiceStatusResponse {
    return ManagementGetTermsOfServiceStatusResponseFromJSONTyped(json, false);
}

export function ManagementGetTermsOfServiceStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetTermsOfServiceStatusResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accepted': !exists(json, 'accepted') ? undefined : json['accepted'],
    };
}

export function ManagementGetTermsOfServiceStatusResponseToJSON(value?: ManagementGetTermsOfServiceStatusResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accepted': value.accepted,
    };
}
