<template>
    <img
        :src="captchaData"
        alt="验证码"
        width="80"
        height="32"
        class="captcha-img"
        @click="reloadCaptcha"
    />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import * as api from '@/api'

const captchaId = defineModel('captchaId')
const captchaData = ref('')

async function reloadCaptcha() {
    try {
        const captchaResponse = await api.SecurityHandlerService.securityHandlerCaptcha({})
        if (captchaResponse.code !== 0) {
            throw new Error('获取验证码失败')
        }

        captchaId.value = captchaResponse.data?.randomId
        const imageData = captchaResponse.data?.image ?? ''
        captchaData.value = 'data:image/png;base64,' + imageData
    } catch (err) {
        console.log(err)
    }
}

reloadCaptcha()
</script>

<style scoped>
.captcha-img {
    cursor: pointer;
    border-radius: 4px;
}
</style>
