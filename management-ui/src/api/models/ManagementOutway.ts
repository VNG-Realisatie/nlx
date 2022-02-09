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
 * @interface ManagementOutway
 */
export interface ManagementOutway {
    /**
     * 
     * @type {string}
     * @memberof ManagementOutway
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutway
     */
    ipAddress?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutway
     */
    publicKeyPEM?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutway
     */
    version?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutway
     */
    publicKeyFingerprint?: string;
}

export function ManagementOutwayFromJSON(json: any): ManagementOutway {
    return ManagementOutwayFromJSONTyped(json, false);
}

export function ManagementOutwayFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementOutway {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'ipAddress': !exists(json, 'ipAddress') ? undefined : json['ipAddress'],
        'publicKeyPEM': !exists(json, 'publicKeyPEM') ? undefined : json['publicKeyPEM'],
        'version': !exists(json, 'version') ? undefined : json['version'],
        'publicKeyFingerprint': !exists(json, 'publicKeyFingerprint') ? undefined : json['publicKeyFingerprint'],
    };
}

export function ManagementOutwayToJSON(value?: ManagementOutway | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'ipAddress': value.ipAddress,
        'publicKeyPEM': value.publicKeyPEM,
        'version': value.version,
        'publicKeyFingerprint': value.publicKeyFingerprint,
    };
}

