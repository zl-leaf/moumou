<template>
    <a-form
        :model="formState"
        name="login-form"
        layout="vertical"
        autocomplete="off"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
    >
        <a-form-item
            label="账号"
            name="username"
            :rules="[{ required: true, message: '请输入账号' }]"
        >
            <a-input
                v-model:value="formState.username"
                placeholder="请输入账号"
                allow-clear
            />
        </a-form-item>

        <a-form-item
            label="密码"
            name="password"
            :rules="[{ required: true, message: '请输入密码' }]"
        >
            <a-input-password
                v-model:value="formState.password"
                placeholder="请输入密码"
            />
        </a-form-item>

        <a-form-item
            label="验证码"
            name="captcha"
            :rules="[{ required: true, message: '请输入验证码' }]"
        >
            <a-input
                v-model:value="formState.captcha"
                placeholder="请输入右侧验证码"
            >
                <template #addonAfter>
                    <Captha v-model:captcha-id="formState.captchaId" />
                </template>
            </a-input>
        </a-form-item>

        <a-form-item class="login-extra">
            <a-space style="width: 100%; justify-content: space-between">
                <a-checkbox v-model:checked="remember">记住登录状态</a-checkbox>
                <span class="login-link">忘记密码？</span>
            </a-space>
        </a-form-item>

        <a-form-item>
            <a-button
                type="primary"
                html-type="submit"
                :loading="loading"
                block
                size="large"
            >
                登录
            </a-button>
        </a-form-item>
    </a-form>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import * as api from '@/api'
import Captha from '@/components/Captha.vue'

const formState = reactive<api.server_api_LoginRequest>({} as api.server_api_LoginRequest)
const loading = ref(false)
const remember = ref(true)

const emit = defineEmits<{
    (e: 'onLoginSuccess', data: any): void
}>()

const onFinish = async () => {
    loading.value = true
    try {
        const loginResponse = await api.SecurityHandlerService.securityHandlerLogin(formState)
        if (loginResponse.code !== 0) {
            message.error(loginResponse.message || '登录失败')
            loading.value = false
            return
        }
        emit('onLoginSuccess', loginResponse.data)
    } catch (err) {
        message.error('网络错误')
        loading.value = false
    }
}

const onFinishFailed = (errorInfo: any) => {
    console.log('Login form validate failed:', errorInfo)
}
</script>

<style scoped>
.login-extra {
    margin-bottom: 8px;
}

.login-link {
    font-size: 13px;
    color: #8c8c8c;
    cursor: default;
}
</style>