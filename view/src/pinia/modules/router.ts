import { defineStore } from 'pinia'
import * as api from '@/api'
import router from '@/router'
import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router';
const viewModules = import.meta.glob('@/views/**/*.vue')

export const useRouterStore = defineStore('router', () => {

    const menData = ref()
    const defaultRouter = ref()

    const setRouterData = (val:api.server_api_Router[]) => {
        menData.value = val
        val.forEach((item:api.server_api_Router) => {
            if (defaultRouter.value === undefined && item.is_menu && item.children?.length == 0) {
                defaultRouter.value = item
            }
        })
    }

    const getDefaultRouter = () => {
        return defaultRouter
    }

    const hasUpdated = ():boolean => {
        return menData.value?.length > 0
    }
    // 数据转换为vue路由
    const formatRouter = (viewRouter:RouteRecordRaw, routerRecordList:api.server_api_Router[]) => {
        routerRecordList.forEach((routerRecord: api.server_api_Router) => {
            let subViewRouter:RouteRecordRaw = {
                name: routerRecord.name,
                path: routerRecord.path ?? '',
                meta: {
                    title: routerRecord.title,
                    isMenu: routerRecord.is_menu,
                },
                children: []
            }
            

            if (routerRecord.component) {
                subViewRouter.component = dynamicImport(viewModules, routerRecord.component);
            } else if (routerRecord.children?.length == 0) {
                subViewRouter.component = () => import('@/components/PageComponent.vue')
                subViewRouter.path += '?page_id=1' // TODO 看下怎么改这个
            }
            formatRouter(subViewRouter, routerRecord.children ?? [])
            viewRouter.children?.push(subViewRouter)
        });
    }
    // 更新路由
    const updateRouter = async () => {
        try {
            const routerTreeResponse = await api.SecurityHandlerService.securityHandlerGetSecurityRouterTree({})
            if (routerTreeResponse.code != 0) {
                return Promise.reject(routerTreeResponse.message);
            }

            const baseRouter = {
                path: '/',
                name: 'layout',
                component: () => import('@/views/AdminLayout.vue'),
                children: [],
            }

            if (routerTreeResponse.data?.routers) {
                setRouterData(routerTreeResponse.data.routers)
                formatRouter(baseRouter, routerTreeResponse.data.routers)
            }

            router.addRoute(baseRouter)
        } catch (err) {
            return Promise.reject("网络错误")
        }
    }

    // 动态加载
    const dynamicImport = (dynamicViewsModules:any, componentName:string) => {
        const matchKeys = Object.keys(dynamicViewsModules).filter((key:string) => {
            const viewPaths = key.split('views/')
            if (!viewPaths.length) {
                return false
            }
            const viewPathLastName = viewPaths[viewPaths.length - 1]
            if (viewPathLastName === componentName + '.vue') {
                return true
            }
            return false
        })
        return dynamicViewsModules[matchKeys[0]]
    }

    return {
        updateRouter,
        getDefaultRouter,
        hasUpdated,
    }
})