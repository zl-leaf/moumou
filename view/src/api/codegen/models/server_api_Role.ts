/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { server_api_Permission } from './server_api_Permission';
import type { server_api_User } from './server_api_User';
export type server_api_Role = {
    id?: string;
    name?: string;
    permissions?: Array<server_api_Permission>;
    users?: Array<server_api_User>;
};

