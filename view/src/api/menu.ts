import axios from "axios";
import {http} from './request'

export type MenuItem = {
    name: string
    path: string
    title: string
    is_menu: boolean
    children: MenuItem[]
}

type MenuResponseData = {
    menu_items: MenuItem[]
}

type MenuResponse = {
    code: string
    message: string
    data: MenuResponseData
}

export const menuAPI = async ():Promise<MenuResponse> => {
    let url = '/menu'
    const response = await http.post(url)
    return response.data
}