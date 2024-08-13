import { defineStore } from 'pinia'
import * as api from '@/api'
import router from '@/router'
import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router';
import { Components } from 'ant-design-vue/es/date-picker/generatePicker';
const viewModules = import.meta.glob('@/views/**/*.vue')

export const useRouterStore = defineStore('router', () => {
    const permissions = ref<string[]>([]) // 记录用户的权限
    const updateRouterFlag = ref(0) // 记录更新菜单次数，出发监听

    const hasUpdated = ():boolean => {
        return true
    }

    // 更新路由
    const updateRouter = async () => {
        updateRouterFlag.value = (updateRouterFlag.value + 1) % 10
        return
    }

    const updatePermissions = async () => {
        permissions.value = []
        try {
            let response = await api.PermissionHandlerService.permissionHandlerGetUserPermissionPath({})
            if (response.code != 0) {
                throw new Error(response.message)
            }
            permissions.value = response.data?.permissions ?? []
        } catch (err) {
            return false
        }
        return true
    }

    return {
        updatePermissions,
        updateRouter,
        hasUpdated,
        updateRouterFlag,
    }
})