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
 * @interface ManagementServiceStatistics
 */
export interface ManagementServiceStatistics {
    /**
     * 
     * @type {string}
     * @memberof ManagementServiceStatistics
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof ManagementServiceStatistics
     */
    incomingAccessRequestCount?: number;
}

/**
 * Check if a given object implements the ManagementServiceStatistics interface.
 */
export function instanceOfManagementServiceStatistics(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementServiceStatisticsFromJSON(json: any): ManagementServiceStatistics {
    return ManagementServiceStatisticsFromJSONTyped(json, false);
}

export function ManagementServiceStatisticsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementServiceStatistics {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'incomingAccessRequestCount': !exists(json, 'incomingAccessRequestCount') ? undefined : json['incomingAccessRequestCount'],
    };
}

export function ManagementServiceStatisticsToJSON(value?: ManagementServiceStatistics | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'incomingAccessRequestCount': value.incomingAccessRequestCount,
    };
}

