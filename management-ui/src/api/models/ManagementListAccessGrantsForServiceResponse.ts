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
import type { NlxmanagementAccessGrant } from './NlxmanagementAccessGrant';
import {
    NlxmanagementAccessGrantFromJSON,
    NlxmanagementAccessGrantFromJSONTyped,
    NlxmanagementAccessGrantToJSON,
} from './NlxmanagementAccessGrant';

/**
 * 
 * @export
 * @interface ManagementListAccessGrantsForServiceResponse
 */
export interface ManagementListAccessGrantsForServiceResponse {
    /**
     * 
     * @type {Array<NlxmanagementAccessGrant>}
     * @memberof ManagementListAccessGrantsForServiceResponse
     */
    accessGrants?: Array<NlxmanagementAccessGrant>;
}

/**
 * Check if a given object implements the ManagementListAccessGrantsForServiceResponse interface.
 */
export function instanceOfManagementListAccessGrantsForServiceResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementListAccessGrantsForServiceResponseFromJSON(json: any): ManagementListAccessGrantsForServiceResponse {
    return ManagementListAccessGrantsForServiceResponseFromJSONTyped(json, false);
}

export function ManagementListAccessGrantsForServiceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListAccessGrantsForServiceResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accessGrants': !exists(json, 'access_grants') ? undefined : ((json['access_grants'] as Array<any>).map(NlxmanagementAccessGrantFromJSON)),
    };
}

export function ManagementListAccessGrantsForServiceResponseToJSON(value?: ManagementListAccessGrantsForServiceResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'access_grants': value.accessGrants === undefined ? undefined : ((value.accessGrants as Array<any>).map(NlxmanagementAccessGrantToJSON)),
    };
}

