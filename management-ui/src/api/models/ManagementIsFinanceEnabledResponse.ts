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
 * @interface ManagementIsFinanceEnabledResponse
 */
export interface ManagementIsFinanceEnabledResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ManagementIsFinanceEnabledResponse
     */
    enabled?: boolean;
}

/**
 * Check if a given object implements the ManagementIsFinanceEnabledResponse interface.
 */
export function instanceOfManagementIsFinanceEnabledResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementIsFinanceEnabledResponseFromJSON(json: any): ManagementIsFinanceEnabledResponse {
    return ManagementIsFinanceEnabledResponseFromJSONTyped(json, false);
}

export function ManagementIsFinanceEnabledResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementIsFinanceEnabledResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'enabled': !exists(json, 'enabled') ? undefined : json['enabled'],
    };
}

export function ManagementIsFinanceEnabledResponseToJSON(value?: ManagementIsFinanceEnabledResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'enabled': value.enabled,
    };
}

