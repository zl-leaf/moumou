/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_CreateRoleRequest } from '../models/server_api_CreateRoleRequest';
import type { server_api_CreateRoleResponse } from '../models/server_api_CreateRoleResponse';
import type { server_api_DeleteRoleRequest } from '../models/server_api_DeleteRoleRequest';
import type { server_api_DeleteRoleResponse } from '../models/server_api_DeleteRoleResponse';
import type { server_api_GetRoleListRequest } from '../models/server_api_GetRoleListRequest';
import type { server_api_GetRoleListResponse } from '../models/server_api_GetRoleListResponse';
import type { server_api_UpdateRoleRequest } from '../models/server_api_UpdateRoleRequest';
import type { server_api_UpdateRoleResponse } from '../models/server_api_UpdateRoleResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class RoleHandlerService {
    /**
     * @param requestBody
     * @returns server_api_CreateRoleResponse OK
     * @throws ApiError
     */
    public static roleHandlerCreateRole(
        requestBody: server_api_CreateRoleRequest,
    ): CancelablePromise<server_api_CreateRoleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/role/create',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_DeleteRoleResponse OK
     * @throws ApiError
     */
    public static roleHandlerDeleteRole(
        requestBody: server_api_DeleteRoleRequest,
    ): CancelablePromise<server_api_DeleteRoleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/role/delete',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetRoleListResponse OK
     * @throws ApiError
     */
    public static roleHandlerGetRoleList(
        requestBody: server_api_GetRoleListRequest,
    ): CancelablePromise<server_api_GetRoleListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/role/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_UpdateRoleResponse OK
     * @throws ApiError
     */
    public static roleHandlerUpdateRole(
        requestBody: server_api_UpdateRoleRequest,
    ): CancelablePromise<server_api_UpdateRoleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/role/update',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
