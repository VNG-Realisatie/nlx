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
    ManagementOrderService,
    ManagementOrderServiceFromJSON,
    ManagementOrderServiceFromJSONTyped,
    ManagementOrderServiceToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementOutgoingOrder
 */
export interface ManagementOutgoingOrder {
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingOrder
     */
    reference?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingOrder
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingOrder
     */
    delegatee?: string;
    /**
     * 
     * @type {Date}
     * @memberof ManagementOutgoingOrder
     */
    validFrom?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ManagementOutgoingOrder
     */
    validUntil?: Date;
    /**
     * 
     * @type {Array<ManagementOrderService>}
     * @memberof ManagementOutgoingOrder
     */
    services?: Array<ManagementOrderService>;
}

export function ManagementOutgoingOrderFromJSON(json: any): ManagementOutgoingOrder {
    return ManagementOutgoingOrderFromJSONTyped(json, false);
}

export function ManagementOutgoingOrderFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementOutgoingOrder {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'delegatee': !exists(json, 'delegatee') ? undefined : json['delegatee'],
        'validFrom': !exists(json, 'validFrom') ? undefined : (new Date(json['validFrom'])),
        'validUntil': !exists(json, 'validUntil') ? undefined : (new Date(json['validUntil'])),
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementOrderServiceFromJSON)),
    };
}

export function ManagementOutgoingOrderToJSON(value?: ManagementOutgoingOrder | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reference': value.reference,
        'description': value.description,
        'delegatee': value.delegatee,
        'validFrom': value.validFrom === undefined ? undefined : (value.validFrom.toISOString()),
        'validUntil': value.validUntil === undefined ? undefined : (value.validUntil.toISOString()),
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementOrderServiceToJSON)),
    };
}


