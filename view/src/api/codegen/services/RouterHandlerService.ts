/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_CreateRouterRequest } from '../models/server_api_CreateRouterRequest';
import type { server_api_CreateRouterResponse } from '../models/server_api_CreateRouterResponse';
import type { server_api_DeleteRouterRequest } from '../models/server_api_DeleteRouterRequest';
import type { server_api_DeleteRouterResponse } from '../models/server_api_DeleteRouterResponse';
import type { server_api_GetRouterInfoRequest } from '../models/server_api_GetRouterInfoRequest';
import type { server_api_GetRouterInfoResponse } from '../models/server_api_GetRouterInfoResponse';
import type { server_api_GetRouterListRequest } from '../models/server_api_GetRouterListRequest';
import type { server_api_GetRouterListResponse } from '../models/server_api_GetRouterListResponse';
import type { server_api_UpdateRouterRequest } from '../models/server_api_UpdateRouterRequest';
import type { server_api_UpdateRouterResponse } from '../models/server_api_UpdateRouterResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class RouterHandlerService {
    /**
     * @param requestBody
     * @returns server_api_CreateRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerCreateRouter(
        requestBody: server_api_CreateRouterRequest,
    ): CancelablePromise<server_api_CreateRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/create',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_DeleteRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerDeleteRouter(
        requestBody: server_api_DeleteRouterRequest,
    ): CancelablePromise<server_api_DeleteRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/delete',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetRouterInfoResponse OK
     * @throws ApiError
     */
    public static routerHandlerGetRouterInfo(
        requestBody: server_api_GetRouterInfoRequest,
    ): CancelablePromise<server_api_GetRouterInfoResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetRouterListResponse OK
     * @throws ApiError
     */
    public static routerHandlerGetRouterList(
        requestBody: server_api_GetRouterListRequest,
    ): CancelablePromise<server_api_GetRouterListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_UpdateRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerUpdateRouter(
        requestBody: server_api_UpdateRouterRequest,
    ): CancelablePromise<server_api_UpdateRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/update',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
