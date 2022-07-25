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
 * @interface ManagementGetServiceResponse
 */
export interface ManagementGetServiceResponse {
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    endpointURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    documentationURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    apiSpecificationURL?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementGetServiceResponse
     */
    internal?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    techSupportContact?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetServiceResponse
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ManagementGetServiceResponse
     */
    inways?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetServiceResponse
     */
    oneTimeCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetServiceResponse
     */
    monthlyCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementGetServiceResponse
     */
    requestCosts?: number;
}

export function ManagementGetServiceResponseFromJSON(json: any): ManagementGetServiceResponse {
    return ManagementGetServiceResponseFromJSONTyped(json, false);
}

export function ManagementGetServiceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetServiceResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'endpointURL': !exists(json, 'endpointURL') ? undefined : json['endpointURL'],
        'documentationURL': !exists(json, 'documentationURL') ? undefined : json['documentationURL'],
        'apiSpecificationURL': !exists(json, 'apiSpecificationURL') ? undefined : json['apiSpecificationURL'],
        'internal': !exists(json, 'internal') ? undefined : json['internal'],
        'techSupportContact': !exists(json, 'techSupportContact') ? undefined : json['techSupportContact'],
        'publicSupportContact': !exists(json, 'publicSupportContact') ? undefined : json['publicSupportContact'],
        'inways': !exists(json, 'inways') ? undefined : json['inways'],
        'oneTimeCosts': !exists(json, 'oneTimeCosts') ? undefined : json['oneTimeCosts'],
        'monthlyCosts': !exists(json, 'monthlyCosts') ? undefined : json['monthlyCosts'],
        'requestCosts': !exists(json, 'requestCosts') ? undefined : json['requestCosts'],
    };
}

export function ManagementGetServiceResponseToJSON(value?: ManagementGetServiceResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'endpointURL': value.endpointURL,
        'documentationURL': value.documentationURL,
        'apiSpecificationURL': value.apiSpecificationURL,
        'internal': value.internal,
        'techSupportContact': value.techSupportContact,
        'publicSupportContact': value.publicSupportContact,
        'inways': value.inways,
        'oneTimeCosts': value.oneTimeCosts,
        'monthlyCosts': value.monthlyCosts,
        'requestCosts': value.requestCosts,
    };
}

