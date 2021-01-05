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
    ManagementListServicesResponseServiceAuthorizationSettingsAuthorization,
    ManagementListServicesResponseServiceAuthorizationSettingsAuthorizationFromJSON,
    ManagementListServicesResponseServiceAuthorizationSettingsAuthorizationFromJSONTyped,
    ManagementListServicesResponseServiceAuthorizationSettingsAuthorizationToJSON,
} from './';

/**
 * 
 * @export
 * @interface ManagementListServicesResponseServiceAuthorizationSettings
 */
export interface ManagementListServicesResponseServiceAuthorizationSettings {
    /**
     * 
     * @type {string}
     * @memberof ManagementListServicesResponseServiceAuthorizationSettings
     */
    mode?: string;
    /**
     * 
     * @type {Array<ManagementListServicesResponseServiceAuthorizationSettingsAuthorization>}
     * @memberof ManagementListServicesResponseServiceAuthorizationSettings
     */
    authorizations?: Array<ManagementListServicesResponseServiceAuthorizationSettingsAuthorization>;
}

export function ManagementListServicesResponseServiceAuthorizationSettingsFromJSON(json: any): ManagementListServicesResponseServiceAuthorizationSettings {
    return ManagementListServicesResponseServiceAuthorizationSettingsFromJSONTyped(json, false);
}

export function ManagementListServicesResponseServiceAuthorizationSettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementListServicesResponseServiceAuthorizationSettings {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mode': !exists(json, 'mode') ? undefined : json['mode'],
        'authorizations': !exists(json, 'authorizations') ? undefined : ((json['authorizations'] as Array<any>).map(ManagementListServicesResponseServiceAuthorizationSettingsAuthorizationFromJSON)),
    };
}

export function ManagementListServicesResponseServiceAuthorizationSettingsToJSON(value?: ManagementListServicesResponseServiceAuthorizationSettings | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'mode': value.mode,
        'authorizations': value.authorizations === undefined ? undefined : ((value.authorizations as Array<any>).map(ManagementListServicesResponseServiceAuthorizationSettingsAuthorizationToJSON)),
    };
}


