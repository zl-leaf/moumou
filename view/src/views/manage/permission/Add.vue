<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-form-item label="名称" :rules="[{ required: true, message: '请输入名称' }]">
                    <a-input v-model:value="formState.name" />
                </a-form-item>
                <a-form-item label="编码" :rules="[{ required: true, message: '请输入编码' }]">
                    <a-input v-model:value="formState.code" />
                </a-form-item>
                <a-form-item label="父级">
                    <a-select v-model:value="formState.pid" :options="parentOptions" allow-clear />
                </a-form-item>
                <a-form-item label="排序">
                    <a-input v-model:value="formState.sort" />
                </a-form-item>
                <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                    <a-button type="primary" @click="onSubmit" :loading="loading" v-permission="'ManagePermissionWrite'">提交</a-button>
                    <a-button style="margin-left: 10px" @click="$router.back()">返回</a-button>
                </a-form-item>
            </a-form>
        </template>
    </ContentPage>
</template>

<script setup lang="ts">
import ContentPage from '@/components/ContentPage.vue';
</script>
<script lang="ts">
import type { SelectProps } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { message } from 'ant-design-vue';

export default defineComponent({
    data() {
        return {
            formState: ref<api.server_api_CreatePermissionRequestData>({}),
            parentOptions: ref<SelectProps['options']>([]),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.initData()
    },
    methods: {
        initData: async function () {
            let response = await api.PermissionHandlerService.permissionHandlerGetPermissionList({
                filter: {},
            })
            let list = response.data?.list ?? []
            this.parentOptions = list.map((item) => ({
                label: item.name,
                value: item.id,
            }))
        },
        onSubmit: async function () {
            if (!this.formState.name || !this.formState.code) {
                message.error('请填写必填项');
                return;
            }

            this.loading = true;
            try {
                let response = await api.PermissionHandlerService.permissionHandlerCreatePermission({ permission: this.formState })
                if (response.code != 0) {
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