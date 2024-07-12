/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { moumou_server_api_GetUserInfoRequest } from '../models/moumou_server_api_GetUserInfoRequest';
import type { moumou_server_api_GetUserInfoResponse } from '../models/moumou_server_api_GetUserInfoResponse';
import type { moumou_server_api_GetUserListRequest } from '../models/moumou_server_api_GetUserListRequest';
import type { moumou_server_api_GetUserListResponse } from '../models/moumou_server_api_GetUserListResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class UserHandlerService {
    /**
     * @param requestBody
     * @returns moumou_server_api_GetUserInfoResponse OK
     * @throws ApiError
     */
    public static userHandlerGetUserInfo(
        requestBody: moumou_server_api_GetUserInfoRequest,
    ): CancelablePromise<moumou_server_api_GetUserInfoResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/user/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_GetUserListResponse OK
     * @throws ApiError
     */
    public static userHandlerGetUserList(
        requestBody: moumou_server_api_GetUserListRequest,
    ): CancelablePromise<moumou_server_api_GetUserListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/user/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
