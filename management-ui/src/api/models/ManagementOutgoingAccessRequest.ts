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
    ManagementErrorDetails,
    ManagementErrorDetailsFromJSON,
    ManagementErrorDetailsFromJSONTyped,
    ManagementErrorDetailsToJSON,
} from './';

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
     * @type {string}
     * @memberof ManagementOutgoingAccessRequest
     */
    organizationName?: string;
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
        'organizationName': !exists(json, 'organizationName') ? undefined : json['organizationName'],
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'state': !exists(json, 'state') ? undefined : ManagementAccessRequestStateFromJSON(json['state']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
        'errorDetails': !exists(json, 'errorDetails') ? undefined : ManagementErrorDetailsFromJSON(json['errorDetails']),
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
        'organizationName': value.organizationName,
        'serviceName': value.serviceName,
        'state': ManagementAccessRequestStateToJSON(value.state),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
        'errorDetails': ManagementErrorDetailsToJSON(value.errorDetails),
    };
}


