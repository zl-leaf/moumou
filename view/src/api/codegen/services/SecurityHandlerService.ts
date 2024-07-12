/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { moumou_server_api_GetPublicKeyRequest } from '../models/moumou_server_api_GetPublicKeyRequest';
import type { moumou_server_api_GetPublicKeyResponse } from '../models/moumou_server_api_GetPublicKeyResponse';
import type { moumou_server_api_GetSecurityRouterTreeRequest } from '../models/moumou_server_api_GetSecurityRouterTreeRequest';
import type { moumou_server_api_GetSecurityRouterTreeResponse } from '../models/moumou_server_api_GetSecurityRouterTreeResponse';
import type { moumou_server_api_LoginRequest } from '../models/moumou_server_api_LoginRequest';
import type { moumou_server_api_LoginResponse } from '../models/moumou_server_api_LoginResponse';
import type { moumou_server_api_LogoutRequest } from '../models/moumou_server_api_LogoutRequest';
import type { moumou_server_api_LogoutResponse } from '../models/moumou_server_api_LogoutResponse';
import type { moumou_server_api_SelfRequest } from '../models/moumou_server_api_SelfRequest';
import type { moumou_server_api_SelfResponse } from '../models/moumou_server_api_SelfResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class SecurityHandlerService {
    /**
     * @param requestBody
     * @returns moumou_server_api_GetPublicKeyResponse OK
     * @throws ApiError
     */
    public static securityHandlerGetPublicKey(
        requestBody: moumou_server_api_GetPublicKeyRequest,
    ): CancelablePromise<moumou_server_api_GetPublicKeyResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/hello',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_LoginResponse OK
     * @throws ApiError
     */
    public static securityHandlerLogin(
        requestBody: moumou_server_api_LoginRequest,
    ): CancelablePromise<moumou_server_api_LoginResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_LogoutResponse OK
     * @throws ApiError
     */
    public static securityHandlerLogout(
        requestBody: moumou_server_api_LogoutRequest,
    ): CancelablePromise<moumou_server_api_LogoutResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/logout',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_GetSecurityRouterTreeResponse OK
     * @throws ApiError
     */
    public static securityHandlerGetSecurityRouterTree(
        requestBody: moumou_server_api_GetSecurityRouterTreeRequest,
    ): CancelablePromise<moumou_server_api_GetSecurityRouterTreeResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/router_tree',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns moumou_server_api_SelfResponse OK
     * @throws ApiError
     */
    public static securityHandlerSelf(
        requestBody: moumou_server_api_SelfRequest,
    ): CancelablePromise<moumou_server_api_SelfResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/self',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
