<template>
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

<script lang="ts">
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { message } from 'ant-design-vue';
import { useRouterStore } from '@/pinia/modules/router';

const routerStore = useRouterStore()
export default defineComponent({
    data() {
        return {
            dataId: "",
            formState: ref<api.server_api_Role>({}),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function(dataId: string) {
            this.loading = true
            try {

                // 加载详情数据
                let infoResponse = await api.RoleHandlerService.roleHandlerGetRoleInfo({
                    id: this.dataId,
                })
                if (infoResponse.code != 0) {
                    throw new Error(infoResponse.message)
                }
                this.formState = infoResponse.data ?? {}
                this.loading = false

            } catch (err) {
                message.error('网络错误')
            }
        },
        onSubmit: async function() {
            this.loading = true;
            let reqData: api.server_api_UpdateRoleRequestData = {
                id: this.dataId,
                name: this.formState.name,
            }

            try {
                let response = await api.RoleHandlerService.roleHandlerUpdateRole({role: reqData})
                if(response.code != 0) {
                    throw new Error(response.message)
                }

                message.success('操作完成')
            } catch (err) {
                message.error('网络错误')
            }
            this.loading = false;
        }
    }
})
</script>