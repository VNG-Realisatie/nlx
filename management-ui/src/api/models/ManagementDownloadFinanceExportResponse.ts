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
 * @interface ManagementDownloadFinanceExportResponse
 */
export interface ManagementDownloadFinanceExportResponse {
    /**
     * 
     * @type {string}
     * @memberof ManagementDownloadFinanceExportResponse
     */
    data?: string;
}

/**
 * Check if a given object implements the ManagementDownloadFinanceExportResponse interface.
 */
export function instanceOfManagementDownloadFinanceExportResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementDownloadFinanceExportResponseFromJSON(json: any): ManagementDownloadFinanceExportResponse {
    return ManagementDownloadFinanceExportResponseFromJSONTyped(json, false);
}

export function ManagementDownloadFinanceExportResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementDownloadFinanceExportResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'data': !exists(json, 'data') ? undefined : json['data'],
    };
}

export function ManagementDownloadFinanceExportResponseToJSON(value?: ManagementDownloadFinanceExportResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'data': value.data,
    };
}

