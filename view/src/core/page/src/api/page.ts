import axios from "axios";
import {http} from './request'

export type PageSchemaFieldAttribute = {
    key: string
    label: string
}

export type PageSchemaAttribute = {
    field_attribute: PageSchemaFieldAttribute
}

export type PageSchema = {
    page_id: number
    attributes: Array<PageSchemaAttribute>
}

export type Page = {
    title: string
    schema: PageSchema
}

type GetPageResponseData = {
    page: Page
}

type GetPageResponse = {
    code: string
    message: string
    data: GetPageResponseData
}

type GetPageDataListResponseData = {
    total: number,
    list: Array<any>
}

type GetPageDataListResponse = {
    code: string
    message: string
    data: GetPageDataListResponseData
}

export const getPageAPI = async (pageId: number):Promise<GetPageResponse> => {
    let url = '/page'
    const response = await http.post(url, {page_id:pageId})
    return response.data
}

export const getPageDataListAPI = async (pageId: number, currentPage: number, pageSize: number):Promise<GetPageDataListResponse> => {
    let url = 'page_data_list'
    const response = await http.post(url, {page_id:pageId, current_page:currentPage, page_size:pageSize})
    return response.data
}