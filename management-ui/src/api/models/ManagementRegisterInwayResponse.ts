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
import type { ManagementInway } from './ManagementInway';
import {
    ManagementInwayFromJSON,
    ManagementInwayFromJSONTyped,
    ManagementInwayToJSON,
} from './ManagementInway';

/**
 * 
 * @export
 * @interface ManagementRegisterInwayResponse
 */
export interface ManagementRegisterInwayResponse {
    /**
     * 
     * @type {ManagementInway}
     * @memberof ManagementRegisterInwayResponse
     */
    inway?: ManagementInway;
}

/**
 * Check if a given object implements the ManagementRegisterInwayResponse interface.
 */
export function instanceOfManagementRegisterInwayResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementRegisterInwayResponseFromJSON(json: any): ManagementRegisterInwayResponse {
    return ManagementRegisterInwayResponseFromJSONTyped(json, false);
}

export function ManagementRegisterInwayResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementRegisterInwayResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'inway': !exists(json, 'inway') ? undefined : ManagementInwayFromJSON(json['inway']),
    };
}

export function ManagementRegisterInwayResponseToJSON(value?: ManagementRegisterInwayResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'inway': ManagementInwayToJSON(value.inway),
    };
}

