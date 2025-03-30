<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-form-item label="名称">
                    <a-input v-model:value="formState.name" />
                </a-form-item>
                <a-form-item label="编码">
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
import { message } from 'ant-design-vue';
import * as api from '@/api'

export default defineComponent({
    data() {
        return {
            dataId: "",
            formState: ref<api.server_api_UpdatePermissionRequestData>({}),
            loading: ref(false),
            parentOptions: ref<SelectProps['options']>([]),
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function (permissionId: string) {
            this.loading = true
            try {
                let response = await api.PermissionHandlerService.permissionHandlerGetPermissionList({
                    filter: {},
                })
                let list = response.data?.list ?? []
                this.parentOptions = list.map((item) => ({
                    label: item.name,
                    value: item.id,
                }))
                let infoResponse = await api.PermissionHandlerService.permissionHandlerGetPermissionInfo({
                    id: permissionId
                })
                if (infoResponse.code != 0) {
                    throw new Error(infoResponse.message)
                }
                let info = infoResponse.data ?? {}
                this.formState = {
                    id: info.id,
                    name: info.name,
                    code: info.code,
                    pid: info.pid,
                    sort: info.sort,
                }
                this.loading = false
            } catch (err) {
                message.error('网络错误')
            }
        },
        onSubmit: async function () {
            this.loading = true
            try {
                let response = await api.PermissionHandlerService.permissionHandlerUpdatePermission({ permission: this.formState })
                if (response.code != 0) {
                    throw new Error(response.message)
                }
                message.success('操作完成')
            } catch (err) {
                message.error('网络错误')
            } finally {
                this.loading = false
            }

        }
    }
})
</script>