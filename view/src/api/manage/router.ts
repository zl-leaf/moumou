import axios from "axios";
import {http} from '../request'

export type RouterModel = {
    id: number
    name: string
    path: string
    title: string
    is_menu: boolean
    pid: number
    sort: number
    component: string
    children?: RouterModel[]
}

type GetSysRouterListResponseData = {
    list: RouterModel[]
}

type GetSysRouterListResponse = {
    code: string
    message: string
    data: GetSysRouterListResponseData
}


type GetSysRouterInfoResponse = {
    code: string
    message: string
    data: RouterModel
}

export const getSysRouterListAPI = async ():Promise<GetSysRouterListResponse> => {
    let url = '/sys_router/list'
    const response = await http.post(url, {})
    return response.data
}


export const getSysRouterInfoAPI = async (id: number):Promise<GetSysRouterInfoResponse> => {
    let url = '/sys_router/info'
    const response = await http.post(url, {
        id: id
    })
    return response.data
}