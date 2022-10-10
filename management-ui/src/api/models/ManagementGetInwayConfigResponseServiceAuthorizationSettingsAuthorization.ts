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
    NlxmanagementOrganization,
    NlxmanagementOrganizationFromJSON,
    NlxmanagementOrganizationFromJSONTyped,
    NlxmanagementOrganizationToJSON,
} from './NlxmanagementOrganization';

/**
 * 
 * @export
 * @interface ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization
 */
export interface ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization {
    /**
     * 
     * @type {NlxmanagementOrganization}
     * @memberof ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization
     */
    organization?: NlxmanagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization
     */
    publicKeyHash?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization
     */
    publicKeyPem?: string;
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSON(json: any): ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization {
    return ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSONTyped(json, false);
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'organization': !exists(json, 'organization') ? undefined : NlxmanagementOrganizationFromJSON(json['organization']),
        'publicKeyHash': !exists(json, 'public_key_hash') ? undefined : json['public_key_hash'],
        'publicKeyPem': !exists(json, 'public_key_pem') ? undefined : json['public_key_pem'],
    };
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationToJSON(value?: ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'organization': NlxmanagementOrganizationToJSON(value.organization),
        'public_key_hash': value.publicKeyHash,
        'public_key_pem': value.publicKeyPem,
    };
}

