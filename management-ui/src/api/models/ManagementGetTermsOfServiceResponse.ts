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
 * @interface ManagementGetTermsOfServiceResponse
 */
export interface ManagementGetTermsOfServiceResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ManagementGetTermsOfServiceResponse
     */
    enabled?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetTermsOfServiceResponse
     */
    url?: string;
}

/**
 * Check if a given object implements the ManagementGetTermsOfServiceResponse interface.
 */
export function instanceOfManagementGetTermsOfServiceResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementGetTermsOfServiceResponseFromJSON(json: any): ManagementGetTermsOfServiceResponse {
    return ManagementGetTermsOfServiceResponseFromJSONTyped(json, false);
}

export function ManagementGetTermsOfServiceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetTermsOfServiceResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'enabled': !exists(json, 'enabled') ? undefined : json['enabled'],
        'url': !exists(json, 'url') ? undefined : json['url'],
    };
}

export function ManagementGetTermsOfServiceResponseToJSON(value?: ManagementGetTermsOfServiceResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'enabled': value.enabled,
        'url': value.url,
    };
}

