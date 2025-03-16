/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_CreatePermissionRequest } from '../models/server_api_CreatePermissionRequest';
import type { server_api_CreatePermissionResponse } from '../models/server_api_CreatePermissionResponse';
import type { server_api_DeletePermissionRequest } from '../models/server_api_DeletePermissionRequest';
import type { server_api_DeletePermissionResponse } from '../models/server_api_DeletePermissionResponse';
import type { server_api_GetPermissionInfoRequest } from '../models/server_api_GetPermissionInfoRequest';
import type { server_api_GetPermissionInfoResponse } from '../models/server_api_GetPermissionInfoResponse';
import type { server_api_GetPermissionListRequest } from '../models/server_api_GetPermissionListRequest';
import type { server_api_GetPermissionListResponse } from '../models/server_api_GetPermissionListResponse';
import type { server_api_GetPermissionTreeRequest } from '../models/server_api_GetPermissionTreeRequest';
import type { server_api_GetPermissionTreeResponse } from '../models/server_api_GetPermissionTreeResponse';
import type { server_api_GetUserPermissionPathRequest } from '../models/server_api_GetUserPermissionPathRequest';
import type { server_api_GetUserPermissionPathResponse } from '../models/server_api_GetUserPermissionPathResponse';
import type { server_api_UpdatePermissionRequest } from '../models/server_api_UpdatePermissionRequest';
import type { server_api_UpdatePermissionResponse } from '../models/server_api_UpdatePermissionResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class PermissionHandlerService {
    /**
     * @param requestBody
     * @returns server_api_CreatePermissionResponse OK
     * @throws ApiError
     */
    public static permissionHandlerCreatePermission(
        requestBody: server_api_CreatePermissionRequest,
    ): CancelablePromise<server_api_CreatePermissionResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/create',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_DeletePermissionResponse OK
     * @throws ApiError
     */
    public static permissionHandlerDeletePermission(
        requestBody: server_api_DeletePermissionRequest,
    ): CancelablePromise<server_api_DeletePermissionResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/delete',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetPermissionInfoResponse OK
     * @throws ApiError
     */
    public static permissionHandlerGetPermissionInfo(
        requestBody: server_api_GetPermissionInfoRequest,
    ): CancelablePromise<server_api_GetPermissionInfoResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
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
    /**
     * @param requestBody
     * @returns server_api_GetPermissionTreeResponse OK
     * @throws ApiError
     */
    public static permissionHandlerGetPermissionTree(
        requestBody: server_api_GetPermissionTreeRequest,
    ): CancelablePromise<server_api_GetPermissionTreeResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/tree',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_UpdatePermissionResponse OK
     * @throws ApiError
     */
    public static permissionHandlerUpdatePermission(
        requestBody: server_api_UpdatePermissionRequest,
    ): CancelablePromise<server_api_UpdatePermissionResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/update',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetUserPermissionPathResponse OK
     * @throws ApiError
     */
    public static permissionHandlerGetUserPermissionPath(
        requestBody: server_api_GetUserPermissionPathRequest,
    ): CancelablePromise<server_api_GetUserPermissionPathResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/permission/user_permission',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
