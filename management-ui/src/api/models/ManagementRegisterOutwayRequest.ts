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
 * @interface ManagementRegisterOutwayRequest
 */
export interface ManagementRegisterOutwayRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementRegisterOutwayRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementRegisterOutwayRequest
     */
    publicKeyPEM?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementRegisterOutwayRequest
     */
    version?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementRegisterOutwayRequest
     */
    selfAddressAPI?: string;
}

/**
 * Check if a given object implements the ManagementRegisterOutwayRequest interface.
 */
export function instanceOfManagementRegisterOutwayRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementRegisterOutwayRequestFromJSON(json: any): ManagementRegisterOutwayRequest {
    return ManagementRegisterOutwayRequestFromJSONTyped(json, false);
}

export function ManagementRegisterOutwayRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementRegisterOutwayRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'publicKeyPEM': !exists(json, 'publicKeyPEM') ? undefined : json['publicKeyPEM'],
        'version': !exists(json, 'version') ? undefined : json['version'],
        'selfAddressAPI': !exists(json, 'selfAddressAPI') ? undefined : json['selfAddressAPI'],
    };
}

export function ManagementRegisterOutwayRequestToJSON(value?: ManagementRegisterOutwayRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'publicKeyPEM': value.publicKeyPEM,
        'version': value.version,
        'selfAddressAPI': value.selfAddressAPI,
    };
}

