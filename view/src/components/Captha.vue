<template>
    <img :src="captchaData" alt="" width="60px" height="30px;">
</template>

<script lang="ts" setup>
import { defineModel, ref } from 'vue';
import * as api from '@/api'
const captchaId = defineModel('captchaId')
const captchaData = ref('')
reloadCaptcha()
async function reloadCaptcha() {
    try {
        const captchaResponse = await api.SecurityHandlerService.securityHandlerCaptcha({})
        if (captchaResponse.code !== 0) {
            throw new Error('获取验证码失败')
        }

        captchaId.value = captchaResponse.data?.randomId
        const imageData = captchaResponse.data?.image ?? ''
        
        captchaData.value = 'data:image/png;base64,' + imageData
    } catch(err) {
        console.log(err)
    }
}
</script>
