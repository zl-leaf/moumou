import router from '@/router'
import {useUserStore} from '@/pinia/modules/user'
import { useRouterStore } from './pinia/modules/router'

const whiteList = ['login']

router.beforeEach(async(to, from) => {
    console.log('permission from', from)
    console.log('permission to', to)

    const userStore = useUserStore()
    const token = userStore.GetToken().value
    if (!token) {
        // 没有登录状态
        if (whiteList.indexOf(String(to.name)) > -1) {
            // 打开登录页面
            return true
        } else {
            // 否则跳转到登录页面
            return {
                name: 'login',
                query: {
                    redirect: document.location.hash
                }
            }
        }
    } else {
        const routerStore = useRouterStore()
        if (whiteList.indexOf(String(to.name)) > -1) {
            // 登录过后不允许再打开登录页面
            if (!routerStore.updateRouterFlag) {
                // 没有更新路由
                try {
                    await routerStore.updateRouter()
                } catch (err) {
                    // 更新路由失败，重新登陆
                    userStore.SetToken("")
                    return {
                        name: 'login',
                        query: {
                            redirect: document.location.hash
                        }
                    }
                }
            }
            return { name: routerStore.getDefaultRouter().value.name, replace: true }
        } else {
            // 正常访问
            if (!routerStore.updateRouterFlag) {
                // 没有更新路由
                await routerStore.updateRouter()
                return { ...to, replace: true }
            }

            return true
        }
    }
})