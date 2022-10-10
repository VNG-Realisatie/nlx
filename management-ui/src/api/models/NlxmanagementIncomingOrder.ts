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
    NlxmanagementOrderService,
    NlxmanagementOrderServiceFromJSON,
    NlxmanagementOrderServiceFromJSONTyped,
    NlxmanagementOrderServiceToJSON,
} from './NlxmanagementOrderService';
import {
    NlxmanagementOrganization,
    NlxmanagementOrganizationFromJSON,
    NlxmanagementOrganizationFromJSONTyped,
    NlxmanagementOrganizationToJSON,
} from './NlxmanagementOrganization';

/**
 * 
 * @export
 * @interface NlxmanagementIncomingOrder
 */
export interface NlxmanagementIncomingOrder {
    /**
     * 
     * @type {string}
     * @memberof NlxmanagementIncomingOrder
     */
    reference?: string;
    /**
     * 
     * @type {string}
     * @memberof NlxmanagementIncomingOrder
     */
    description?: string;
    /**
     * 
     * @type {NlxmanagementOrganization}
     * @memberof NlxmanagementIncomingOrder
     */
    delegator?: NlxmanagementOrganization;
    /**
     * 
     * @type {Date}
     * @memberof NlxmanagementIncomingOrder
     */
    validFrom?: Date;
    /**
     * 
     * @type {Date}
     * @memberof NlxmanagementIncomingOrder
     */
    validUntil?: Date;
    /**
     * 
     * @type {Array<NlxmanagementOrderService>}
     * @memberof NlxmanagementIncomingOrder
     */
    services?: Array<NlxmanagementOrderService>;
    /**
     * 
     * @type {Date}
     * @memberof NlxmanagementIncomingOrder
     */
    revokedAt?: Date;
}

export function NlxmanagementIncomingOrderFromJSON(json: any): NlxmanagementIncomingOrder {
    return NlxmanagementIncomingOrderFromJSONTyped(json, false);
}

export function NlxmanagementIncomingOrderFromJSONTyped(json: any, ignoreDiscriminator: boolean): NlxmanagementIncomingOrder {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reference': !exists(json, 'reference') ? undefined : json['reference'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'delegator': !exists(json, 'delegator') ? undefined : NlxmanagementOrganizationFromJSON(json['delegator']),
        'validFrom': !exists(json, 'valid_from') ? undefined : (new Date(json['valid_from'])),
        'validUntil': !exists(json, 'valid_until') ? undefined : (new Date(json['valid_until'])),
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(NlxmanagementOrderServiceFromJSON)),
        'revokedAt': !exists(json, 'revoked_at') ? undefined : (new Date(json['revoked_at'])),
    };
}

export function NlxmanagementIncomingOrderToJSON(value?: NlxmanagementIncomingOrder | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reference': value.reference,
        'description': value.description,
        'delegator': NlxmanagementOrganizationToJSON(value.delegator),
        'valid_from': value.validFrom === undefined ? undefined : (value.validFrom.toISOString()),
        'valid_until': value.validUntil === undefined ? undefined : (value.validUntil.toISOString()),
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(NlxmanagementOrderServiceToJSON)),
        'revoked_at': value.revokedAt === undefined ? undefined : (value.revokedAt.toISOString()),
    };
}

