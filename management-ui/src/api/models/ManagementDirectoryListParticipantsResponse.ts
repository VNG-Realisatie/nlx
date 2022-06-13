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
    DirectoryListParticipantsResponseDirectoryParticipant,
    DirectoryListParticipantsResponseDirectoryParticipantFromJSON,
    DirectoryListParticipantsResponseDirectoryParticipantFromJSONTyped,
    DirectoryListParticipantsResponseDirectoryParticipantToJSON,
} from './DirectoryListParticipantsResponseDirectoryParticipant';

/**
 * 
 * @export
 * @interface ManagementDirectoryListParticipantsResponse
 */
export interface ManagementDirectoryListParticipantsResponse {
    /**
     * 
     * @type {Array<DirectoryListParticipantsResponseDirectoryParticipant>}
     * @memberof ManagementDirectoryListParticipantsResponse
     */
    participants?: Array<DirectoryListParticipantsResponseDirectoryParticipant>;
}

export function ManagementDirectoryListParticipantsResponseFromJSON(json: any): ManagementDirectoryListParticipantsResponse {
    return ManagementDirectoryListParticipantsResponseFromJSONTyped(json, false);
}

export function ManagementDirectoryListParticipantsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementDirectoryListParticipantsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'participants': !exists(json, 'participants') ? undefined : ((json['participants'] as Array<any>).map(DirectoryListParticipantsResponseDirectoryParticipantFromJSON)),
    };
}

export function ManagementDirectoryListParticipantsResponseToJSON(value?: ManagementDirectoryListParticipantsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'participants': value.participants === undefined ? undefined : ((value.participants as Array<any>).map(DirectoryListParticipantsResponseDirectoryParticipantToJSON)),
    };
}

