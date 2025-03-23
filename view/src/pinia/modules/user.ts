import {defineStore} from 'pinia'
import {ref, watch} from 'vue'
import * as api from '@/api'
import crypto from 'crypto-js';
import cookie from 'js-cookie'
import { useRouterStore } from './router';
import router from '@/router';

export const useUserStore = defineStore('user', () => {
    const token = ref(window.localStorage.getItem('token') || cookie.get('x-token') || '')
    const permissions = ref<string[]>([]) // 记录用户的权限


    const SetToken = (val:any) => {
        token.value = val
    }

    const GetToken = () => {
        return token
    }

    const Logout = async() => {
        try {
            const logoutResponse = await api.SecurityHandlerService.securityHandlerLogout({})
            if (logoutResponse.code != 0) {
                return Promise.reject(logoutResponse.message);
            }

            // 清理数据
            SetToken('')
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

    const UpdatePermissions = async () => {
        permissions.value = []
        try {
            let response = await api.PermissionHandlerService.permissionHandlerGetUserPermission({})
            if (response.code != 0) {
                throw new Error(response.message)
            }
            permissions.value = response.data?.permissions ?? []
        } catch (err) {
            return false
        }
        return true
    }

    const HasPermission = (permission: string) => {
        return permissions.value.includes(permission)
    }

    watch(() => token.value, () => {
        window.localStorage.setItem('token', token.value)
    })
    
    return {
        Logout,
        SetToken,
        GetToken,
        UpdatePermissions,
        HasPermission,
    }
})