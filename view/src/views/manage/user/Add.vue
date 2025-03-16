<template>
    <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
        <a-form-item label="账号" :rules="[{ required: true, message: '请输入账号' }]" >
            <a-input v-model:value="formState.username" />
        </a-form-item>
        <a-form-item label="密码" :rules="[{ required: true, message: '请输入密码' }]" >
            <a-input-password v-model:value="formState.password" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
            <a-button type="primary" @click="onSubmit" :loading="loading">提交</a-button>
            <a-button style="margin-left: 10px" @click="$router.back()">返回</a-button>
        </a-form-item>
    </a-form>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { message } from 'ant-design-vue';

export default defineComponent({
    data() {
        return {
            formState: ref<api.server_api_CreateUserRequestData>({}),
            loading: ref<boolean>(false),
        }
    },
    methods: {
        onSubmit: async function() {
            if (!this.formState.username || !this.formState.password) {
                message.error('请填写必填项');
                return;
            }

            this.loading = true;
            let reqData: api.server_api_CreateUserRequestData = {
                username: this.formState.username,
                password: this.formState.password,
            }
            try {
                let response = await api.UserHandlerService.userHandlerCreateUser({user: reqData})
                if(response.code != 0) {
                    throw new Error(response.message)
                }
                if (!response.data) {
                    throw new Error('详情缺少参数')
                }

                router.back()
            } catch (err) {
                message.error('网络错误')
            }
            this.loading = false;
        }
    }
})
</script> 