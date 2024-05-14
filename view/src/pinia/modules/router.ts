import { defineStore } from 'pinia'
import { routerTreeAPI } from '@/api/router'
import router from '@/router'
import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router';
import type {RouterRecord} from '@/api/router'


export const useRouterStore = defineStore('router', () => {

    const menData = ref()
    const defaultRouter = ref()

    const setRouterData = (val:RouterRecord[]) => {
        menData.value = val
        val.forEach((item:RouterRecord) => {
            if (defaultRouter.value === undefined && item.is_menu && item.children?.length == 0) {
                console.log('default', item)
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

    const formatRouter = (viewRouter:RouteRecordRaw, routerRecordList:RouterRecord[]) => {
        routerRecordList.forEach((routerRecord: RouterRecord) => {
            const subViewRouter:RouteRecordRaw  = {
                name: routerRecord.name,
                path: routerRecord.path,
                component: () => import('@/core/page/src/components/PageComponent.vue'),
                meta: {
                    title: routerRecord.title,
                    isMenu: routerRecord.is_menu,
                },
                children: (routerRecord.children.length > 0) ? []: undefined
            }
            formatRouter(subViewRouter, routerRecord.children)
            viewRouter.children?.push(subViewRouter)
        });
    }
    const updateRouter = async () => {
        try {
            const routerTreeResponse = await routerTreeAPI()
            if (routerTreeResponse.code != '0') {
                return Promise.reject(routerTreeResponse.message);
            }
            setRouterData(routerTreeResponse.data.routers)

            const baseRouter = {
                path: '/',
                name: 'layout',
                component: () => import('@/views/AdminLayout.vue'),
                children: []
            }


            formatRouter(baseRouter, routerTreeResponse.data.routers)
            router.addRoute(baseRouter)

            console.log('router', router.getRoutes())

        } catch (err) {
            return Promise.reject("网络错误")
        }
    }

    return {
        updateRouter,
        getDefaultRouter,
        hasUpdated,
    }
})