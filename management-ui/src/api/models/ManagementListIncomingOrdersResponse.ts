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
import type { ManagementIncomingOrder } from './ManagementIncomingOrder';
import {
    ManagementIncomingOrderFromJSON,
    ManagementIncomingOrderFromJSONTyped,
    ManagementIncomingOrderToJSON,
} from './ManagementIncomingOrder';

/**
 * 
 * @export
 * @interface ManagementListIncomingOrdersResponse
 */
export interface ManagementListIncomingOrdersResponse {
    /**
     * 
     * @type {Array<ManagementIncomingOrder>}
     * @memberof ManagementListIncomingOrdersResponse
     */
    orders?: Array<ManagementIncomingOrder>;
}

/**
 * Check if a given object implements the ManagementListIncomingOrdersResponse interface.
 */
export function instanceOfManagementListIncomingOrdersResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementListIncomingOrdersResponseFromJSON(json: any): ManagementListIncomingOrdersResponse {
    return ManagementListIncomingOrdersResponseFromJSONTyped(json, false);
}

export function ManagementListIncomingOrdersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListIncomingOrdersResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'orders': !exists(json, 'orders') ? undefined : ((json['orders'] as Array<any>).map(ManagementIncomingOrderFromJSON)),
    };
}

export function ManagementListIncomingOrdersResponseToJSON(value?: ManagementListIncomingOrdersResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'orders': value.orders === undefined ? undefined : ((value.orders as Array<any>).map(ManagementIncomingOrderToJSON)),
    };
}

