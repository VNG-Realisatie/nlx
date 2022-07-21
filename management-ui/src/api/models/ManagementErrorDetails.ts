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
import type { ManagementErrorCode } from './ManagementErrorCode';
import {
    ManagementErrorCodeFromJSON,
    ManagementErrorCodeFromJSONTyped,
    ManagementErrorCodeToJSON,
} from './ManagementErrorCode';

/**
 * 
 * @export
 * @interface ManagementErrorDetails
 */
export interface ManagementErrorDetails {
    /**
     * 
     * @type {ManagementErrorCode}
     * @memberof ManagementErrorDetails
     */
    code?: ManagementErrorCode;
    /**
     * 
     * @type {string}
     * @memberof ManagementErrorDetails
     */
    cause?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ManagementErrorDetails
     */
    stackTrace?: Array<string>;
}

/**
 * Check if a given object implements the ManagementErrorDetails interface.
 */
export function instanceOfManagementErrorDetails(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementErrorDetailsFromJSON(json: any): ManagementErrorDetails {
    return ManagementErrorDetailsFromJSONTyped(json, false);
}

export function ManagementErrorDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementErrorDetails {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : ManagementErrorCodeFromJSON(json['code']),
        'cause': !exists(json, 'cause') ? undefined : json['cause'],
        'stackTrace': !exists(json, 'stackTrace') ? undefined : json['stackTrace'],
    };
}

export function ManagementErrorDetailsToJSON(value?: ManagementErrorDetails | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': ManagementErrorCodeToJSON(value.code),
        'cause': value.cause,
        'stackTrace': value.stackTrace,
    };
}

