/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_LoginRequest } from '../models/server_api_LoginRequest';
import type { server_api_LoginResponse } from '../models/server_api_LoginResponse';
import type { server_api_LogoutRequest } from '../models/server_api_LogoutRequest';
import type { server_api_LogoutResponse } from '../models/server_api_LogoutResponse';
import type { server_api_SelfRequest } from '../models/server_api_SelfRequest';
import type { server_api_SelfResponse } from '../models/server_api_SelfResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class SecurityHandlerService {
    /**
     * @param requestBody
     * @returns server_api_LoginResponse OK
     * @throws ApiError
     */
    public static securityHandlerLogin(
        requestBody: server_api_LoginRequest,
    ): CancelablePromise<server_api_LoginResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_LogoutResponse OK
     * @throws ApiError
     */
    public static securityHandlerLogout(
        requestBody: server_api_LogoutRequest,
    ): CancelablePromise<server_api_LogoutResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/logout',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_SelfResponse OK
     * @throws ApiError
     */
    public static securityHandlerSelf(
        requestBody: server_api_SelfRequest,
    ): CancelablePromise<server_api_SelfResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/security/self',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
