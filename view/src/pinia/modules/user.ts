import {defineStore} from 'pinia'
import {ref, watch} from 'vue'
import {helloAPI, loginAPI, logoutAPI, selfAPI} from '@/core/auth/src/api/user'
import crypto from 'crypto-js';
import cookie from 'js-cookie'
import { useRouterStore } from './router';
import router from '@/router';

export const useUserStore = defineStore('user', () => {
    const userInfo = ref({
        userID: '',
    })
    const token = ref(window.localStorage.getItem('token') || cookie.get('x-token') || '')

    const SetUserInfo = (val:any) => {
        userInfo.value = val
    }

    const GetUserInfo = () =>{
        return userInfo
    }

    const SetToken = (val:any) => {
        token.value = val
    }

    const GetToken = () => {
        return token
    }

    const Logout = async() => {
        try {
            const logoutResponse = await logoutAPI(GetToken().value)
            if (logoutResponse.code != '0') {
                return Promise.reject(logoutResponse.message);
            }

            // 清理数据
            SetToken('')
            SetUserInfo({})
            sessionStorage.clear()
            window.localStorage.removeItem('token')
            cookie.remove('x-token')

            // 跳转
            router.push({ name: 'login', replace: true })
            window.location.reload()

        } catch (err) {
            return Promise.reject("网络错误")
        }
    }

    const Self = async() => {
        try {
            const selfResponse = await selfAPI(GetToken().value)
            if (selfResponse.code != '0') {
                return Promise.reject(selfResponse.message)
            }
            SetUserInfo({
                userID: selfResponse.data.user_id
            })
        } catch (err) {
            return Promise.reject('网络错误')
        }
    }

    watch(() => token.value, () => {
        window.localStorage.setItem('token', token.value)
    })
    
    return {
        Logout,
        Self,
        SetUserInfo,
        GetUserInfo,
        SetToken,
        GetToken,
    }
})