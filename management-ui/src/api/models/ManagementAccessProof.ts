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
import type { ExternalOrganization } from './ExternalOrganization';
import {
    ExternalOrganizationFromJSON,
    ExternalOrganizationFromJSONTyped,
    ExternalOrganizationToJSON,
} from './ExternalOrganization';

/**
 * 
 * @export
 * @interface ManagementAccessProof
 */
export interface ManagementAccessProof {
    /**
     * 
     * @type {string}
     * @memberof ManagementAccessProof
     */
    id?: string;
    /**
     * 
     * @type {ExternalOrganization}
     * @memberof ManagementAccessProof
     */
    organization?: ExternalOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementAccessProof
     */
    serviceName?: string;
    /**
     * 
     * @type {Date}
     * @memberof ManagementAccessProof
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ManagementAccessProof
     */
    revokedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof ManagementAccessProof
     */
    accessRequestId?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementAccessProof
     */
    publicKeyFingerprint?: string;
    /**
     * 
     * @type {Date}
     * @memberof ManagementAccessProof
     */
    terminatedAt?: Date;
}

/**
 * Check if a given object implements the ManagementAccessProof interface.
 */
export function instanceOfManagementAccessProof(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementAccessProofFromJSON(json: any): ManagementAccessProof {
    return ManagementAccessProofFromJSONTyped(json, false);
}

export function ManagementAccessProofFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementAccessProof {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'organization': !exists(json, 'organization') ? undefined : ExternalOrganizationFromJSON(json['organization']),
        'serviceName': !exists(json, 'service_name') ? undefined : json['service_name'],
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'revokedAt': !exists(json, 'revoked_at') ? undefined : (new Date(json['revoked_at'])),
        'accessRequestId': !exists(json, 'access_request_id') ? undefined : json['access_request_id'],
        'publicKeyFingerprint': !exists(json, 'public_key_fingerprint') ? undefined : json['public_key_fingerprint'],
        'terminatedAt': !exists(json, 'terminated_at') ? undefined : (new Date(json['terminated_at'])),
    };
}

export function ManagementAccessProofToJSON(value?: ManagementAccessProof | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'organization': ExternalOrganizationToJSON(value.organization),
        'service_name': value.serviceName,
        'created_at': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'revoked_at': value.revokedAt === undefined ? undefined : (value.revokedAt.toISOString()),
        'access_request_id': value.accessRequestId,
        'public_key_fingerprint': value.publicKeyFingerprint,
        'terminated_at': value.terminatedAt === undefined ? undefined : (value.terminatedAt.toISOString()),
    };
}

