import axios from "axios";
import crypto from 'crypto-js';
import {http} from './request'

type UserData = {
    user_id: string
}

type EncryptData = {
    key: string;
    iv: string
}

type EncryptResponse = {
    code: string
    message: string
    data: EncryptData
}


type LoginData = {
    user: UserData
    token: string
}

type LoginResonse = {
    code: string
    message: string
    data: LoginData
}

type LogoutResponse = {
    code: string
    message: string
}

type SelfResponse = {
    code: string
    message: string
    data: UserData 
}



export const helloAPI = async ():Promise<EncryptResponse>  => {
    let url = '/hello';
    const response = await http.post(url)
    return response.data
}

export const loginAPI = async (data: any):Promise<LoginResonse> => {
    let url = '/login';
    const response = await http.post(url, {
        username: data.username,
        password: data.password,
    })
    return response.data
}

export const selfAPI = async (token:string):Promise<SelfResponse> => {
    let url = '/self';
    const response = await http.post(url, {}, {
        headers: {
            "x-token": token
        }
    })
    return response.data
}

export const logoutAPI = async (token:string):Promise<LogoutResponse> => {
    let url = '/logout';
    const response = await http.post(url, {}, {
        headers: {
            "x-token": token
        }
    })
    return response.data
}