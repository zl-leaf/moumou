<template>
    <a-form :model="formState" name="basic" :label-col="{ span: 10 }" :wrapper-col="{ span: 4 }" autocomplete="off"
        @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="账号" name="username"
            :rules="[{ required: true, message: '请输入账号' }]">
            <a-input v-model:value="formState.username" />
        </a-form-item>

        <a-form-item label="密码" name="password"
            :rules="[{ required: true, message: '请输入密码' }]">
            <a-input-password v-model:value="formState.password" />
        </a-form-item>

        <a-form-item label="验证码" name="captcha"
            :rules="[{ required: true, message: '请输入验证码' }]">
            <a-input v-model:value="formState.captcha" >
                <template #addonAfter>
                    <Captha v-model:captcha-id="formState.captchaId" />
                </template>
            </a-input>
        </a-form-item>

        <a-form-item :wrapper-col="{ offset: 10, span: 16 }">
            <a-button type="primary" html-type="submit" :loading="loading">Login</a-button>
        </a-form-item>
    </a-form>
</template>
<script setup lang=ts>
import Captha from '@/components/Captha.vue'
</script>
<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import * as api from '@/api'

export default defineComponent({
    data() {
        return {
            formState: reactive<api.server_api_LoginRequest>({}),
            loading: ref<boolean>(false),
        }
    },
    emits: {
        // 回调成功
        onLoginSuccess: function(loginResponseData:any) {}
    },
    methods: {
        onFinish: function (values: any) {
            this.loading = true;
            this.login().then((loginResponseData: any) => {
                setTimeout(() => {
                    this.$emit('onLoginSuccess', loginResponseData)
                }, 1000);
            }).catch(err => {
                message.error(err)
                this.loading = false
            })
        },
        // 异常
        onFinishFailed: function (errorInfo: any) {
            console.log('Failed:', errorInfo);
        },
        // 登录
        login: async function () {
            try {
                const loginResponse = await api.SecurityHandlerService.securityHandlerLogin(this.formState)
                if (loginResponse.code != 0) {
                    return Promise.reject(loginResponse.message);
                }
                return loginResponse.data
            } catch (err) {
                return Promise.reject("网络错误")
            }
        },
    }
})
</script>