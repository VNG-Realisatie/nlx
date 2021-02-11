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
 * @interface ManagementUpdateServiceRequest
 */
export interface ManagementUpdateServiceRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    endpointURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    documentationURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    apiSpecificationURL?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementUpdateServiceRequest
     */
    internal?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    techSupportContact?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementUpdateServiceRequest
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ManagementUpdateServiceRequest
     */
    inways?: Array<string>;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceRequest
     */
    oneTimeCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceRequest
     */
    monthlyCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementUpdateServiceRequest
     */
    requestCosts?: number;
}

export function ManagementUpdateServiceRequestFromJSON(json: any): ManagementUpdateServiceRequest {
    return ManagementUpdateServiceRequestFromJSONTyped(json, false);
}

export function ManagementUpdateServiceRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementUpdateServiceRequest {
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

export function ManagementUpdateServiceRequestToJSON(value?: ManagementUpdateServiceRequest | null): any {
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


