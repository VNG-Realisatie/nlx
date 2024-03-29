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
import type { ManagementTXLogRecord } from './ManagementTXLogRecord';
import {
    ManagementTXLogRecordFromJSON,
    ManagementTXLogRecordFromJSONTyped,
    ManagementTXLogRecordToJSON,
} from './ManagementTXLogRecord';

/**
 * 
 * @export
 * @interface ManagementTXLogServiceListRecordsResponse
 */
export interface ManagementTXLogServiceListRecordsResponse {
    /**
     * 
     * @type {Array<ManagementTXLogRecord>}
     * @memberof ManagementTXLogServiceListRecordsResponse
     */
    records?: Array<ManagementTXLogRecord>;
}

/**
 * Check if a given object implements the ManagementTXLogServiceListRecordsResponse interface.
 */
export function instanceOfManagementTXLogServiceListRecordsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementTXLogServiceListRecordsResponseFromJSON(json: any): ManagementTXLogServiceListRecordsResponse {
    return ManagementTXLogServiceListRecordsResponseFromJSONTyped(json, false);
}

export function ManagementTXLogServiceListRecordsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementTXLogServiceListRecordsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'records': !exists(json, 'records') ? undefined : ((json['records'] as Array<any>).map(ManagementTXLogRecordFromJSON)),
    };
}

export function ManagementTXLogServiceListRecordsResponseToJSON(value?: ManagementTXLogServiceListRecordsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'records': value.records === undefined ? undefined : ((value.records as Array<any>).map(ManagementTXLogRecordToJSON)),
    };
}

