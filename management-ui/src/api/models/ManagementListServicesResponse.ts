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
import type { ManagementListServicesResponseService } from './ManagementListServicesResponseService';
import {
    ManagementListServicesResponseServiceFromJSON,
    ManagementListServicesResponseServiceFromJSONTyped,
    ManagementListServicesResponseServiceToJSON,
} from './ManagementListServicesResponseService';

/**
 * 
 * @export
 * @interface ManagementListServicesResponse
 */
export interface ManagementListServicesResponse {
    /**
     * 
     * @type {Array<ManagementListServicesResponseService>}
     * @memberof ManagementListServicesResponse
     */
    services?: Array<ManagementListServicesResponseService>;
}

/**
 * Check if a given object implements the ManagementListServicesResponse interface.
 */
export function instanceOfManagementListServicesResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementListServicesResponseFromJSON(json: any): ManagementListServicesResponse {
    return ManagementListServicesResponseFromJSONTyped(json, false);
}

export function ManagementListServicesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListServicesResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementListServicesResponseServiceFromJSON)),
    };
}

export function ManagementListServicesResponseToJSON(value?: ManagementListServicesResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementListServicesResponseServiceToJSON)),
    };
}

