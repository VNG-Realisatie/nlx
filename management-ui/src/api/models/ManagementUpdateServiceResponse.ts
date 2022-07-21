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
    endpointURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    documentationURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceResponse
     */
    apiSpecificationURL?: string;
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

export function ManagementUpdateServiceResponseToJSON(value?: ManagementUpdateServiceResponse | null): any {
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

