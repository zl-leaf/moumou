/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_GetPermissionListRequest } from '../models/server_api_GetPermissionListRequest';
import type { server_api_GetPermissionListResponse } from '../models/server_api_GetPermissionListResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class PermissionHandlerService {
    /**
     * @param requestBody
     * @returns server_api_GetPermissionListResponse OK
     * @throws ApiError
     */
    public static permissionHandlerGetPermissionList(
        requestBody: server_api_GetPermissionListRequest,
    ): CancelablePromise<server_api_GetPermissionListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
