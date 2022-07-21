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
import type { ManagementAccessRequestState } from './ManagementAccessRequestState';
import {
    ManagementAccessRequestStateFromJSON,
    ManagementAccessRequestStateFromJSONTyped,
    ManagementAccessRequestStateToJSON,
} from './ManagementAccessRequestState';
import type { ManagementErrorDetails } from './ManagementErrorDetails';
import {
    ManagementErrorDetailsFromJSON,
    ManagementErrorDetailsFromJSONTyped,
    ManagementErrorDetailsToJSON,
} from './ManagementErrorDetails';
import type { ManagementOrganization } from './ManagementOrganization';
import {
    ManagementOrganizationFromJSON,
    ManagementOrganizationFromJSONTyped,
    ManagementOrganizationToJSON,
} from './ManagementOrganization';

/**
 * 
 * @export
 * @interface ManagementOutgoingAccessRequest
 */
export interface ManagementOutgoingAccessRequest {
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingAccessRequest
     */
    id?: string;
    /**
     * 
     * @type {ManagementOrganization}
     * @memberof ManagementOutgoingAccessRequest
     */
    organization?: ManagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingAccessRequest
     */
    serviceName?: string;
    /**
     * 
     * @type {ManagementAccessRequestState}
     * @memberof ManagementOutgoingAccessRequest
     */
    state?: ManagementAccessRequestState;
    /**
     * 
     * @type {Date}
     * @memberof ManagementOutgoingAccessRequest
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof ManagementOutgoingAccessRequest
     */
    updatedAt?: Date;
    /**
     * 
     * @type {ManagementErrorDetails}
     * @memberof ManagementOutgoingAccessRequest
     */
    errorDetails?: ManagementErrorDetails;
    /**
     * 
     * @type {string}
     * @memberof ManagementOutgoingAccessRequest
     */
    publicKeyFingerprint?: string;
}

/**
 * Check if a given object implements the ManagementOutgoingAccessRequest interface.
 */
export function instanceOfManagementOutgoingAccessRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementOutgoingAccessRequestFromJSON(json: any): ManagementOutgoingAccessRequest {
    return ManagementOutgoingAccessRequestFromJSONTyped(json, false);
}

export function ManagementOutgoingAccessRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementOutgoingAccessRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'organization': !exists(json, 'organization') ? undefined : ManagementOrganizationFromJSON(json['organization']),
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'state': !exists(json, 'state') ? undefined : ManagementAccessRequestStateFromJSON(json['state']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
        'errorDetails': !exists(json, 'errorDetails') ? undefined : ManagementErrorDetailsFromJSON(json['errorDetails']),
        'publicKeyFingerprint': !exists(json, 'publicKeyFingerprint') ? undefined : json['publicKeyFingerprint'],
    };
}

export function ManagementOutgoingAccessRequestToJSON(value?: ManagementOutgoingAccessRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'organization': ManagementOrganizationToJSON(value.organization),
        'serviceName': value.serviceName,
        'state': ManagementAccessRequestStateToJSON(value.state),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
        'errorDetails': ManagementErrorDetailsToJSON(value.errorDetails),
        'publicKeyFingerprint': value.publicKeyFingerprint,
    };
}

