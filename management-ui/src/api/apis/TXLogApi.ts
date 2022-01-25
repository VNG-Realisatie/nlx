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


import * as runtime from '../runtime';
import {
    ManagementTXLogListRecordsResponse,
    ManagementTXLogListRecordsResponseFromJSON,
    ManagementTXLogListRecordsResponseToJSON,
    RpcStatus,
    RpcStatusFromJSON,
    RpcStatusToJSON,
} from '../models';

/**
 * 
 */
export class TXLogApi extends runtime.BaseAPI {

    /**
     */
    async tXLogListRecordsRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<ManagementTXLogListRecordsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v1/txlog/records`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ManagementTXLogListRecordsResponseFromJSON(jsonValue));
    }

    /**
     */
    async tXLogListRecords(initOverrides?: RequestInit): Promise<ManagementTXLogListRecordsResponse> {
        const response = await this.tXLogListRecordsRaw(initOverrides);
        return await response.value();
    }

}
