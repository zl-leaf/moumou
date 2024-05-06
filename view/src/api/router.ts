import axios from "axios";
import {http} from './request'

export type RouterRecord = {
    name: string
    path: string
    title: string
    is_menu: boolean
    children: RouterRecord[]
}

type RouterTreeResponseData = {
    routers: RouterRecord[]
}

type RouterTreeResponse = {
    code: string
    message: string
    data: RouterTreeResponseData
}

export const routerTreeAPI = async ():Promise<RouterTreeResponse> => {
    let url = '/router_tree'
    const response = await http.post(url)
    return response.data
}