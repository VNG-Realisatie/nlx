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
 * @interface ManagementTXLogOrder
 */
export interface ManagementTXLogOrder {
    /**
     * 
     * @type {string}
     * @memberof ManagementTXLogOrder
     */
    delegator?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementTXLogOrder
     */
    reference?: string;
}

export function ManagementTXLogOrderFromJSON(json: any): ManagementTXLogOrder {
    return ManagementTXLogOrderFromJSONTyped(json, false);
}

export function ManagementTXLogOrderFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementTXLogOrder {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'delegator': !exists(json, 'delegator') ? undefined : json['delegator'],
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
    };
}

export function ManagementTXLogOrderToJSON(value?: ManagementTXLogOrder | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'delegator': value.delegator,
        'reference': value.reference,
    };
}

