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
    ManagementServiceStatistics,
    ManagementServiceStatisticsFromJSON,
    ManagementServiceStatisticsFromJSONTyped,
    ManagementServiceStatisticsToJSON,
} from './ManagementServiceStatistics';

/**
 * 
 * @export
 * @interface ManagementGetStatisticsOfServicesResponse
 */
export interface ManagementGetStatisticsOfServicesResponse {
    /**
     * 
     * @type {Array<ManagementServiceStatistics>}
     * @memberof ManagementGetStatisticsOfServicesResponse
     */
    services?: Array<ManagementServiceStatistics>;
}

export function ManagementGetStatisticsOfServicesResponseFromJSON(json: any): ManagementGetStatisticsOfServicesResponse {
    return ManagementGetStatisticsOfServicesResponseFromJSONTyped(json, false);
}

export function ManagementGetStatisticsOfServicesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetStatisticsOfServicesResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementServiceStatisticsFromJSON)),
    };
}

export function ManagementGetStatisticsOfServicesResponseToJSON(value?: ManagementGetStatisticsOfServicesResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementServiceStatisticsToJSON)),
    };
}

