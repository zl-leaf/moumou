<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
            <a-form-item label="名称">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                <a-button type="primary" @click="onSubmit" :loading="loading" v-permission="'ManageRoleWrite'">提交</a-button>
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
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { message } from 'ant-design-vue';

export default defineComponent({
    data() {
        return {
            formState: ref<api.server_api_Role>({}),
            loading: ref<boolean>(false),
        }
    },
    methods: {
        onSubmit: async function() {
            this.loading = true;
            let reqData: api.server_api_CreateRoleRequestData = {
                name: this.formState.name,
            }
            try {
                let response = await api.RoleHandlerService.roleHandlerCreateRole({role: reqData})
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