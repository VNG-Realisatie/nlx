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

/**
 * 
 * @export
 * @enum {string}
 */
export enum ManagementErrorCode {
    INTERNAL = 'INTERNAL',
    NO_INWAY_SELECTED = 'NO_INWAY_SELECTED'
}

export function ManagementErrorCodeFromJSON(json: any): ManagementErrorCode {
    return ManagementErrorCodeFromJSONTyped(json, false);
}

export function ManagementErrorCodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementErrorCode {
    return json as ManagementErrorCode;
}

export function ManagementErrorCodeToJSON(value?: ManagementErrorCode | null): any {
    return value as any;
}

