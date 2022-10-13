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
 * @interface ManagementUpdateServiceResponse
 */
export interface ManagementUpdateServiceResponse {
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    endpointUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    documentationUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    apiSpecificationUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementUpdateServiceResponse
     */
    internal?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    techSupportContact?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ManagementUpdateServiceResponse
     */
    inways?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceResponse
     */
    oneTimeCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceResponse
     */
    monthlyCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceResponse
     */
    requestCosts?: number;
}

/**
 * Check if a given object implements the ManagementUpdateServiceResponse interface.
 */
export function instanceOfManagementUpdateServiceResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementUpdateServiceResponseFromJSON(json: any): ManagementUpdateServiceResponse {
    return ManagementUpdateServiceResponseFromJSONTyped(json, false);
}

export function ManagementUpdateServiceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementUpdateServiceResponse {
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
        'inways': !exists(json, 'inways') ? undefined : json['inways'],
        'oneTimeCosts': !exists(json, 'one_time_costs') ? undefined : json['one_time_costs'],
        'monthlyCosts': !exists(json, 'monthly_costs') ? undefined : json['monthly_costs'],
        'requestCosts': !exists(json, 'request_costs') ? undefined : json['request_costs'],
    };
}

export function ManagementUpdateServiceResponseToJSON(value?: ManagementUpdateServiceResponse | null): any {
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
        'inways': value.inways,
        'one_time_costs': value.oneTimeCosts,
        'monthly_costs': value.monthlyCosts,
        'request_costs': value.requestCosts,
    };
}

