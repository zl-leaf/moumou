<template>
    <a-form :model="formState" name="basic" :label-col="{ span: 8 }" :wrapper-col="{ span: 4 }" autocomplete="off"
        @finish="onFinish" @finishFailed="onFinishFailed">
        <a-form-item label="Username" name="username"
            :rules="[{ required: true, message: 'Please input your username!' }]">
            <a-input v-model:value="formState.username" />
        </a-form-item>

        <a-form-item label="Password" name="password"
            :rules="[{ required: true, message: 'Please input your password!' }]">
            <a-input-password v-model:value="formState.password" />
        </a-form-item>

        <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
            <a-button type="primary" html-type="submit" :loading="loading">Login</a-button>
        </a-form-item>
    </a-form>
</template>
<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import * as api from '@/api'

interface FormState {
    username: string;
    password: string;
}

export default defineComponent({
    data() {
        return {
            formState: reactive<FormState>({
                username: '',
                password: '',
            }),
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
                const loginResponse = await api.SecurityHandlerService.securityHandlerLogin({
                    username: this.formState.username,
                    password: this.formState.password,
                })
                if (loginResponse.code != 0) {
                    return Promise.reject(loginResponse.message);
                }
                return loginResponse.data
            } catch (err) {
                return Promise.reject("网络错误")
            }
        }
    }
})
</script>