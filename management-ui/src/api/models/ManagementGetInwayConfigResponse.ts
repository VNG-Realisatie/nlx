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
import type { ManagementGetInwayConfigResponseService } from './ManagementGetInwayConfigResponseService';
import {
    ManagementGetInwayConfigResponseServiceFromJSON,
    ManagementGetInwayConfigResponseServiceFromJSONTyped,
    ManagementGetInwayConfigResponseServiceToJSON,
} from './ManagementGetInwayConfigResponseService';

/**
 * 
 * @export
 * @interface ManagementGetInwayConfigResponse
 */
export interface ManagementGetInwayConfigResponse {
    /**
     * 
     * @type {Array<ManagementGetInwayConfigResponseService>}
     * @memberof ManagementGetInwayConfigResponse
     */
    services?: Array<ManagementGetInwayConfigResponseService>;
    /**
     * 
     * @type {boolean}
     * @memberof ManagementGetInwayConfigResponse
     */
    isOrganizationInway?: boolean;
}

/**
 * Check if a given object implements the ManagementGetInwayConfigResponse interface.
 */
export function instanceOfManagementGetInwayConfigResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementGetInwayConfigResponseFromJSON(json: any): ManagementGetInwayConfigResponse {
    return ManagementGetInwayConfigResponseFromJSONTyped(json, false);
}

export function ManagementGetInwayConfigResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetInwayConfigResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'services': !exists(json, 'services') ? undefined : ((json['services'] as Array<any>).map(ManagementGetInwayConfigResponseServiceFromJSON)),
        'isOrganizationInway': !exists(json, 'is_organization_inway') ? undefined : json['is_organization_inway'],
    };
}

export function ManagementGetInwayConfigResponseToJSON(value?: ManagementGetInwayConfigResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'services': value.services === undefined ? undefined : ((value.services as Array<any>).map(ManagementGetInwayConfigResponseServiceToJSON)),
        'is_organization_inway': value.isOrganizationInway,
    };
}

