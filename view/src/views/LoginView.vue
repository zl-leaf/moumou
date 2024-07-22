<template>
    <a-layout style="min-height: 100vh">
        <a-layout-header style="background: #fff; padding: 0" />
        <a-layout-content>
            <LoginComponent @onLoginSuccess="onLoginSuccess" />
        </a-layout-content>
        <a-layout-footer style="text-align: center">
            Ant Design ©2018 Created by Ant UED
        </a-layout-footer>
    </a-layout>
</template>

<script setup lang="ts">
import LoginComponent from '@/components/LoginComponent.vue'
import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router';
import router from '@/router/index'

const userStore = useUserStore()
const routerStore = useRouterStore()

const redirectToDefaultPage = async () => {
    try {
        await routerStore.updateRouter()
        router.replace({ name: '/' })
    } catch (err) {
        console.log('error', err)
    }
}

const onLoginSuccess = function (ret: any) {
    console.log(ret)

    userStore.SetUserInfo({
        userID: ret.user.user_id
    })
    userStore.SetToken(ret.token)
    redirectToDefaultPage()
}

</script>