/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_CreateArticleRequest } from '../models/server_api_CreateArticleRequest';
import type { server_api_CreateArticleResponse } from '../models/server_api_CreateArticleResponse';
import type { server_api_DeleteArticleRequest } from '../models/server_api_DeleteArticleRequest';
import type { server_api_DeleteArticleResponse } from '../models/server_api_DeleteArticleResponse';
import type { server_api_GetArticleInfoRequest } from '../models/server_api_GetArticleInfoRequest';
import type { server_api_GetArticleInfoResponse } from '../models/server_api_GetArticleInfoResponse';
import type { server_api_GetArticleListRequest } from '../models/server_api_GetArticleListRequest';
import type { server_api_GetArticleListResponse } from '../models/server_api_GetArticleListResponse';
import type { server_api_UpdateArticleRequest } from '../models/server_api_UpdateArticleRequest';
import type { server_api_UpdateArticleResponse } from '../models/server_api_UpdateArticleResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class ArticleHandlerService {
    /**
     * @param requestBody
     * @returns server_api_CreateArticleResponse OK
     * @throws ApiError
     */
    public static articleHandlerCreateArticle(
        requestBody: server_api_CreateArticleRequest,
    ): CancelablePromise<server_api_CreateArticleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/article/create',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_DeleteArticleResponse OK
     * @throws ApiError
     */
    public static articleHandlerDeleteArticle(
        requestBody: server_api_DeleteArticleRequest,
    ): CancelablePromise<server_api_DeleteArticleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/article/delete',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetArticleInfoResponse OK
     * @throws ApiError
     */
    public static articleHandlerGetArticleInfo(
        requestBody: server_api_GetArticleInfoRequest,
    ): CancelablePromise<server_api_GetArticleInfoResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/article/info',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_GetArticleListResponse OK
     * @throws ApiError
     */
    public static articleHandlerGetArticleList(
        requestBody: server_api_GetArticleListRequest,
    ): CancelablePromise<server_api_GetArticleListResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/article/list',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * @param requestBody
     * @returns server_api_UpdateArticleResponse OK
     * @throws ApiError
     */
    public static articleHandlerUpdateArticle(
        requestBody: server_api_UpdateArticleRequest,
    ): CancelablePromise<server_api_UpdateArticleResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/article/update',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
