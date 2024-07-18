<template>
    <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
        <a-form-item label="账号">
            <a-input v-model:value="formState.username" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
            <a-button type="primary" @click="onSubmit">提交</a-button>
            <a-button style="margin-left: 10px">返回</a-button>
        </a-form-item>
    </a-form>
</template>

<script lang="ts">
import { defineComponent,ref } from 'vue';
import * as api from '@/api'

export default defineComponent({
    data() {
        return {
            formState: ref<api.server_api_User>({}),
        }
    },
    created() {
        api.UserHandlerService.userHandlerGetUserInfo({
            id: String(this.$route.query.id)
        }).then((response) => {
            if (response.code != 0) {
                return Promise.reject(response.message)
            }
            return response.data
        }).then((data) => {
             this.formState = data ?? {}
        }).catch(err => {
            console.log('err:', err)
        })
    },
    methods: {
        onSubmit: function() {
        }
    }
})
</script>