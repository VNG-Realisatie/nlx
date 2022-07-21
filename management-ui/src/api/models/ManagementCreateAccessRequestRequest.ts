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
 * @interface ManagementCreateAccessRequestRequest
 */
export interface ManagementCreateAccessRequestRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementCreateAccessRequestRequest
     */
    organizationSerialNumber?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementCreateAccessRequestRequest
     */
    serviceName?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementCreateAccessRequestRequest
     */
    publicKeyPEM?: string;
}

/**
 * Check if a given object implements the ManagementCreateAccessRequestRequest interface.
 */
export function instanceOfManagementCreateAccessRequestRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementCreateAccessRequestRequestFromJSON(json: any): ManagementCreateAccessRequestRequest {
    return ManagementCreateAccessRequestRequestFromJSONTyped(json, false);
}

export function ManagementCreateAccessRequestRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementCreateAccessRequestRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organizationSerialNumber': !exists(json, 'organizationSerialNumber') ? undefined : json['organizationSerialNumber'],
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'publicKeyPEM': !exists(json, 'publicKeyPEM') ? undefined : json['publicKeyPEM'],
    };
}

export function ManagementCreateAccessRequestRequestToJSON(value?: ManagementCreateAccessRequestRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'organizationSerialNumber': value.organizationSerialNumber,
        'serviceName': value.serviceName,
        'publicKeyPEM': value.publicKeyPEM,
    };
}

