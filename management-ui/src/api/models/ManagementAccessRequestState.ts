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
 */
export const ManagementAccessRequestState = {
    UNSPECIFIED: 'UNSPECIFIED',
    FAILED: 'FAILED',
    CREATED: 'CREATED',
    RECEIVED: 'RECEIVED',
    APPROVED: 'APPROVED',
    REJECTED: 'REJECTED',
    REVOKED: 'REVOKED'
} as const;
export type ManagementAccessRequestState = typeof ManagementAccessRequestState[keyof typeof ManagementAccessRequestState];


export function ManagementAccessRequestStateFromJSON(json: any): ManagementAccessRequestState {
    return ManagementAccessRequestStateFromJSONTyped(json, false);
}

export function ManagementAccessRequestStateFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementAccessRequestState {
    return json as ManagementAccessRequestState;
}

export function ManagementAccessRequestStateToJSON(value?: ManagementAccessRequestState | null): any {
    return value as any;
}

