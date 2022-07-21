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
import type { ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization } from './ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization';
import {
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSON,
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSONTyped,
    ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationToJSON,
} from './ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization';

/**
 * 
 * @export
 * @interface ManagementGetInwayConfigResponseServiceAuthorizationSettings
 */
export interface ManagementGetInwayConfigResponseServiceAuthorizationSettings {
    /**
     * 
     * @type {Array<ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization>}
     * @memberof ManagementGetInwayConfigResponseServiceAuthorizationSettings
     */
    authorizations?: Array<ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorization>;
}

/**
 * Check if a given object implements the ManagementGetInwayConfigResponseServiceAuthorizationSettings interface.
 */
export function instanceOfManagementGetInwayConfigResponseServiceAuthorizationSettings(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSON(json: any): ManagementGetInwayConfigResponseServiceAuthorizationSettings {
    return ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSONTyped(json, false);
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ManagementGetInwayConfigResponseServiceAuthorizationSettings {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authorizations': !exists(json, 'authorizations') ? undefined : ((json['authorizations'] as Array<any>).map(ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationFromJSON)),
    };
}

export function ManagementGetInwayConfigResponseServiceAuthorizationSettingsToJSON(value?: ManagementGetInwayConfigResponseServiceAuthorizationSettings | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authorizations': value.authorizations === undefined ? undefined : ((value.authorizations as Array<any>).map(ManagementGetInwayConfigResponseServiceAuthorizationSettingsAuthorizationToJSON)),
    };
}

