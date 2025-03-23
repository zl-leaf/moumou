import { defineStore } from 'pinia'
import * as api from '@/api'
import router from '@/router'
import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router';
import { Components } from 'ant-design-vue/es/date-picker/generatePicker';
const viewModules = import.meta.glob('@/views/**/*.vue')

// TODO 这里需要增加一个权限的判断，用权限过滤菜单
export const useRouterStore = defineStore('router', () => {
    const updateRouterFlag = ref(0) // 记录更新菜单次数，出发监听

    const hasUpdated = ():boolean => {
        return true
    }

    // 更新路由
    const updateRouter = async () => {
        updateRouterFlag.value = (updateRouterFlag.value + 1) % 10
        return
    }

    return {
        updateRouter,
        hasUpdated,
        updateRouterFlag,
    }
})