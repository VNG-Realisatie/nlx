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
    DirectoryServiceState,
    DirectoryServiceStateFromJSON,
    DirectoryServiceStateFromJSONTyped,
    DirectoryServiceStateToJSON,
    ManagementAccessProof,
    ManagementAccessProofFromJSON,
    ManagementAccessProofFromJSONTyped,
    ManagementAccessProofToJSON,
    ManagementOrganization,
    ManagementOrganizationFromJSON,
    ManagementOrganizationFromJSONTyped,
    ManagementOrganizationToJSON,
    ManagementOutgoingAccessRequest,
    ManagementOutgoingAccessRequestFromJSON,
    ManagementOutgoingAccessRequestFromJSONTyped,
    ManagementOutgoingAccessRequestToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementDirectoryService
 */
export interface ManagementDirectoryService {
    /**
     * 
     * @type {string}
     * @memberof ManagementDirectoryService
     */
    serviceName?: string;
    /**
     * 
     * @type {ManagementOrganization}
     * @memberof ManagementDirectoryService
     */
    organization?: ManagementOrganization;
    /**
     * 
     * @type {string}
     * @memberof ManagementDirectoryService
     */
    apiSpecificationType?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementDirectoryService
     */
    documentationURL?: string;
    /**
     * 
     * @type {string}
     * @memberof ManagementDirectoryService
     */
    publicSupportContact?: string;
    /**
     * 
     * @type {DirectoryServiceState}
     * @memberof ManagementDirectoryService
     */
    state?: DirectoryServiceState;
    /**
     * 
     * @type {ManagementOutgoingAccessRequest}
     * @memberof ManagementDirectoryService
     */
    latestAccessRequest?: ManagementOutgoingAccessRequest;
    /**
     * 
     * @type {ManagementAccessProof}
     * @memberof ManagementDirectoryService
     */
    latestAccessProof?: ManagementAccessProof;
    /**
     * 
     * @type {number}
     * @memberof ManagementDirectoryService
     */
    oneTimeCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementDirectoryService
     */
    monthlyCosts?: number;
    /**
     * 
     * @type {number}
     * @memberof ManagementDirectoryService
     */
    requestCosts?: number;
}

export function ManagementDirectoryServiceFromJSON(json: any): ManagementDirectoryService {
    return ManagementDirectoryServiceFromJSONTyped(json, false);
}

export function ManagementDirectoryServiceFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementDirectoryService {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'serviceName': !exists(json, 'serviceName') ? undefined : json['serviceName'],
        'organization': !exists(json, 'organization') ? undefined : ManagementOrganizationFromJSON(json['organization']),
        'apiSpecificationType': !exists(json, 'apiSpecificationType') ? undefined : json['apiSpecificationType'],
        'documentationURL': !exists(json, 'documentationURL') ? undefined : json['documentationURL'],
        'publicSupportContact': !exists(json, 'publicSupportContact') ? undefined : json['publicSupportContact'],
        'state': !exists(json, 'state') ? undefined : DirectoryServiceStateFromJSON(json['state']),
        'latestAccessRequest': !exists(json, 'latestAccessRequest') ? undefined : ManagementOutgoingAccessRequestFromJSON(json['latestAccessRequest']),
        'latestAccessProof': !exists(json, 'latestAccessProof') ? undefined : ManagementAccessProofFromJSON(json['latestAccessProof']),
        'oneTimeCosts': !exists(json, 'oneTimeCosts') ? undefined : json['oneTimeCosts'],
        'monthlyCosts': !exists(json, 'monthlyCosts') ? undefined : json['monthlyCosts'],
        'requestCosts': !exists(json, 'requestCosts') ? undefined : json['requestCosts'],
    };
}

export function ManagementDirectoryServiceToJSON(value?: ManagementDirectoryService | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'serviceName': value.serviceName,
        'organization': ManagementOrganizationToJSON(value.organization),
        'apiSpecificationType': value.apiSpecificationType,
        'documentationURL': value.documentationURL,
        'publicSupportContact': value.publicSupportContact,
        'state': DirectoryServiceStateToJSON(value.state),
        'latestAccessRequest': ManagementOutgoingAccessRequestToJSON(value.latestAccessRequest),
        'latestAccessProof': ManagementAccessProofToJSON(value.latestAccessProof),
        'oneTimeCosts': value.oneTimeCosts,
        'monthlyCosts': value.monthlyCosts,
        'requestCosts': value.requestCosts,
    };
}


