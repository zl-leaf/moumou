import axios from "axios";
import {http} from '../request'

export type UserModel = {
    username: string
}

type GetSysUserListResponseData = {
    list: UserModel[]
    total: number
}

type GetSysUserListResponse = {
    code: string
    message: string
    data: GetSysUserListResponseData
}

type GetSysUserListFilter = {

}

type GetSysUserInfoResponse = {
    code: string
    message: string
    data: UserModel
}

export const getSysUserListAPI = async (filter:GetSysUserListFilter, currentPage: number, pageSize: number):Promise<GetSysUserListResponse> => {
    let url = '/sys_user/list'
    const response = await http.post(url, {
        filter: filter,
        current_page: currentPage,
        page_size: pageSize
    })
    return response.data
}


export const getSysUserInfoAPI = async (id: number):Promise<GetSysUserInfoResponse> => {
    let url = '/sys_user/info'
    const response = await http.post(url, {
        id: id
    })
    return response.data
}