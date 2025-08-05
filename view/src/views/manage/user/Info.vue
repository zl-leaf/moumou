<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
            <a-form-item label="账号">
                <a-input v-model:value="formState.username" />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                <a-button type="primary" @click="onSubmit" :loading="loading" v-permission="'ManageUserWrite'">提交</a-button>
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
import { defineComponent,ref } from 'vue';
import { message } from 'ant-design-vue';
import * as api from '@/api'

export default defineComponent({
    data() {
        return {
            dataId: "",
            formState: ref<api.server_api_User>({}),
            loading: ref(false),
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function(userId:string) {
            this.loading = true
            try {
                let infoResponse = await api.UserHandlerService.userHandlerGetUserInfo({id: userId})
                if (infoResponse.code != 0) {
                    throw new Error(infoResponse.message)
                }
                this.formState = infoResponse.data ?? {}
                this.loading = false
            } catch(err) {
                message.error('网络错误')
            }
        },
        onSubmit: async function() {
            this.loading = true
            let reqData: api.server_api_UpdateUserRequestData = {
                id: this.dataId,
                username: this.formState.username,
            }

            try {
                let response = await api.UserHandlerService.userHandlerUpdateUser({user: reqData})
                if(response.code != 0) {
                    throw new Error(response.message)
                }
                message.success('操作完成')
            } catch(err) {
                message.error('网络错误')
            } finally {
                this.loading = false
            }
            
        }
    }
})
</script>