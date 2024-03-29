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
import type { ManagementGetInwayConfigResponseServiceAuthorizationSettings } from './ManagementGetInwayConfigResponseServiceAuthorizationSettings';
import {
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSON,
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSONTyped,
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsToJSON,
} from './ManagementGetInwayConfigResponseServiceAuthorizationSettings';

/**
 * 
 * @export
 * @interface ManagementGetInwayConfigResponseService
 */
export interface ManagementGetInwayConfigResponseService {
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    endpointUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    documentationUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    apiSpecificationUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementGetInwayConfigResponseService
     */
    internal?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    techSupportContact?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseService
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {ManagementGetInwayConfigResponseServiceAuthorizationSettings}
     * @memberof ManagementGetInwayConfigResponseService
     */
    authorizationSettings?: ManagementGetInwayConfigResponseServiceAuthorizationSettings;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetInwayConfigResponseService
     */
    oneTimeCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetInwayConfigResponseService
     */
    monthlyCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetInwayConfigResponseService
     */
    requestCosts?: number;
}

/**
 * Check if a given object implements the ManagementGetInwayConfigResponseService interface.
 */
export function instanceOfManagementGetInwayConfigResponseService(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementGetInwayConfigResponseServiceFromJSON(json: any): ManagementGetInwayConfigResponseService {
    return ManagementGetInwayConfigResponseServiceFromJSONTyped(json, false);
}

export function ManagementGetInwayConfigResponseServiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetInwayConfigResponseService {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'endpointUrl': !exists(json, 'endpoint_url') ? undefined : json['endpoint_url'],
        'documentationUrl': !exists(json, 'documentation_url') ? undefined : json['documentation_url'],
        'apiSpecificationUrl': !exists(json, 'api_specification_url') ? undefined : json['api_specification_url'],
        'internal': !exists(json, 'internal') ? undefined : json['internal'],
        'techSupportContact': !exists(json, 'tech_support_contact') ? undefined : json['tech_support_contact'],
        'publicSupportContact': !exists(json, 'public_support_contact') ? undefined : json['public_support_contact'],
        'authorizationSettings': !exists(json, 'authorization_settings') ? undefined : ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSON(json['authorization_settings']),
        'oneTimeCosts': !exists(json, 'one_time_costs') ? undefined : json['one_time_costs'],
        'monthlyCosts': !exists(json, 'monthly_costs') ? undefined : json['monthly_costs'],
        'requestCosts': !exists(json, 'request_costs') ? undefined : json['request_costs'],
    };
}

export function ManagementGetInwayConfigResponseServiceToJSON(value?: ManagementGetInwayConfigResponseService | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'endpoint_url': value.endpointUrl,
        'documentation_url': value.documentationUrl,
        'api_specification_url': value.apiSpecificationUrl,
        'internal': value.internal,
        'tech_support_contact': value.techSupportContact,
        'public_support_contact': value.publicSupportContact,
        'authorization_settings': ManagementGetInwayConfigResponseServiceAuthorizationSettingsToJSON(value.authorizationSettings),
        'one_time_costs': value.oneTimeCosts,
        'monthly_costs': value.monthlyCosts,
        'request_costs': value.requestCosts,
    };
}

