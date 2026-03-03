/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_InitializeRequest } from '../models/server_api_InitializeRequest';
import type { server_api_InitializeResponse } from '../models/server_api_InitializeResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class SystemHandlerService {
    /**
     * @param requestBody
     * @returns server_api_InitializeResponse OK
     * @throws ApiError
     */
    public static systemHandlerInitialize(
        requestBody: server_api_InitializeRequest,
    ): CancelablePromise<server_api_InitializeResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/system/initialize',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
