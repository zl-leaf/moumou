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
import router from '@/router/index'

const userStore = useUserStore()

const redirectToDefaultPage = () => {
    let redirect = router.currentRoute.value.query.redirect?.toString() ?? ''
    if (redirect === '') {
        redirect = '/'
    }
    router.replace({ path: redirect })
}

const onLoginSuccess = function (ret: any) {
    console.log(ret)
    userStore.SetToken(ret.token)
    redirectToDefaultPage()
}

</script>