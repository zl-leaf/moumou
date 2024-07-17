/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { moumou_server_api_CreateRouterRequest } from '../models/moumou_server_api_CreateRouterRequest';
import type { moumou_server_api_CreateRouterResponse } from '../models/moumou_server_api_CreateRouterResponse';
import type { moumou_server_api_DeleteRouterRequest } from '../models/moumou_server_api_DeleteRouterRequest';
import type { moumou_server_api_DeleteRouterResponse } from '../models/moumou_server_api_DeleteRouterResponse';
import type { moumou_server_api_GetRouterInfoRequest } from '../models/moumou_server_api_GetRouterInfoRequest';
import type { moumou_server_api_GetRouterInfoResponse } from '../models/moumou_server_api_GetRouterInfoResponse';
import type { moumou_server_api_GetRouterListRequest } from '../models/moumou_server_api_GetRouterListRequest';
import type { moumou_server_api_GetRouterListResponse } from '../models/moumou_server_api_GetRouterListResponse';
import type { moumou_server_api_UpdateRouterRequest } from '../models/moumou_server_api_UpdateRouterRequest';
import type { moumou_server_api_UpdateRouterResponse } from '../models/moumou_server_api_UpdateRouterResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class RouterHandlerService {
    /**
     * @param requestBody
     * @returns moumou_server_api_CreateRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerCreateRouter(
        requestBody: moumou_server_api_CreateRouterRequest,
    ): CancelablePromise<moumou_server_api_CreateRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/create',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_DeleteRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerDeleteRouter(
        requestBody: moumou_server_api_DeleteRouterRequest,
    ): CancelablePromise<moumou_server_api_DeleteRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/delete',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_GetRouterInfoResponse OK
     * @throws ApiError
     */
    public static routerHandlerGetRouterInfo(
        requestBody: moumou_server_api_GetRouterInfoRequest,
    ): CancelablePromise<moumou_server_api_GetRouterInfoResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_GetRouterListResponse OK
     * @throws ApiError
     */
    public static routerHandlerGetRouterList(
        requestBody: moumou_server_api_GetRouterListRequest,
    ): CancelablePromise<moumou_server_api_GetRouterListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_UpdateRouterResponse OK
     * @throws ApiError
     */
    public static routerHandlerUpdateRouter(
        requestBody: moumou_server_api_UpdateRouterRequest,
    ): CancelablePromise<moumou_server_api_UpdateRouterResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/router/update',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
