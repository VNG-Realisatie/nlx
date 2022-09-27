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
    ManagementAccessRequestState,
    ManagementAccessRequestStateFromJSON,
    ManagementAccessRequestStateFromJSONTyped,
    ManagementAccessRequestStateToJSON,
} from './ManagementAccessRequestState';
import {
    ManagementOrganization,
    ManagementOrganizationFromJSON,
    ManagementOrganizationFromJSONTyped,
    ManagementOrganizationToJSON,
} from './ManagementOrganization';

/**
 * 
 * @export
 * @interface ManagementIncomingAccessRequest
 */
export interface ManagementIncomingAccessRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    id?: string;
    /**
     * 
     * @type {ManagementOrganization}
     * @memberof ManagementIncomingAccessRequest
     */
    organization?: ManagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    serviceName?: string;
    /**
     * 
     * @type {ManagementAccessRequestState}
     * @memberof ManagementIncomingAccessRequest
     */
    state?: ManagementAccessRequestState;
    /**
     * 
     * @type {Date}
     * @memberof ManagementIncomingAccessRequest
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ManagementIncomingAccessRequest
     */
    updatedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof ManagementIncomingAccessRequest
     */
    publicKeyFingerprint?: string;
}

export function ManagementIncomingAccessRequestFromJSON(json: any): ManagementIncomingAccessRequest {
    return ManagementIncomingAccessRequestFromJSONTyped(json, false);
}

export function ManagementIncomingAccessRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementIncomingAccessRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'organization': !exists(json, 'organization') ? undefined : ManagementOrganizationFromJSON(json['organization']),
        'serviceName': !exists(json, 'service_name') ? undefined : json['service_name'],
        'state': !exists(json, 'state') ? undefined : ManagementAccessRequestStateFromJSON(json['state']),
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updatedAt': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
        'publicKeyFingerprint': !exists(json, 'public_key_fingerprint') ? undefined : json['public_key_fingerprint'],
    };
}

export function ManagementIncomingAccessRequestToJSON(value?: ManagementIncomingAccessRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'organization': ManagementOrganizationToJSON(value.organization),
        'service_name': value.serviceName,
        'state': ManagementAccessRequestStateToJSON(value.state),
        'created_at': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updated_at': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
        'public_key_fingerprint': value.publicKeyFingerprint,
    };
}

