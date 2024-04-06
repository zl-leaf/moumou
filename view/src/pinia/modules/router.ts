import { defineStore } from 'pinia'
import { menuAPI } from '@/api/menu'
import router from '@/router'
import { ref } from 'vue'
import type { RouteRecordRaw } from 'vue-router';
import type {MenuItem} from '@/api/menu'


export const useRouterStore = defineStore('router', () => {

    const menData = ref()
    const defaultRouter = ref()

    const setMenuData = (val:MenuItem[]) => {
        menData.value = val
        val.forEach((item:MenuItem) => {
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

    const formatRouter = (viewRouter:RouteRecordRaw, menuItems:MenuItem[]) => {
        menuItems.forEach((menuItem: MenuItem) => {
            const subViewRouter:RouteRecordRaw  = {
                name: menuItem.name,
                path: menuItem.path,
                component: import('@/components/TableComponent.vue'),
                meta: {
                    title: menuItem.title,
                    isMenu: menuItem.is_menu,
                },
                children: (menuItem.children.length > 0) ? []: undefined
            }
            formatRouter(subViewRouter, menuItem.children)
            viewRouter.children?.push(subViewRouter)
        });
    }
    const updateRouter = async () => {
        try {
            const menuResponse = await menuAPI()
            if (menuResponse.code != '0') {
                return Promise.reject(menuResponse.message);
            }
            setMenuData(menuResponse.data.menu_items)

            const baseRouter = {
                path: '/',
                name: 'layout',
                component: () => import('@/views/AdminLayout.vue'),
                children: []
            }


            formatRouter(baseRouter, menuResponse.data.menu_items)
            router.addRoute(baseRouter)

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