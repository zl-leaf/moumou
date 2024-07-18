<template>
    <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
        <a-form-item label="名称">
            <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="标题">
            <a-input v-model:value="formState.title" />
        </a-form-item>
        <a-form-item label="菜单">
            <a-radio-group v-model:value="formState.is_menu">
                <a-radio :value=true>是</a-radio>
                <a-radio :value=false>否</a-radio>
            </a-radio-group>
        </a-form-item>
        <a-form-item label="路径">
            <a-input v-model:value="formState.path" />
        </a-form-item>
        <a-form-item label="父节点">
            <a-tree-select 
                v-model:value="formState.pid" 
                :tree-data="routerTreeData" 
                :replaceFields="{key:'id', value:'id'}"
                tree-default-expand-all>
            </a-tree-select>
        </a-form-item>
        <a-form-item label="组件">
            <a-input v-model:value="formState.component" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
            <a-button type="primary" @click="onSubmit" :loading="loading">提交</a-button>
            <a-button style="margin-left: 10px">返回</a-button>
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
            routerId: "",
            formState: ref<api.server_api_Router>({}),
            routerTreeData: ref<api.server_api_Router[]>(),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.routerId = String(this.$route.query.id)

        // 加载数节点数据
        api.RouterHandlerService.routerHandlerGetRouterList({}).then((response) => {
            if (response.code != 0) {
                return Promise.reject(response.message)
            }
            return response.data
        }).then((data) => {
            this.routerTreeData = data?.list
        }).catch(err => {
            console.log('err:', err)
        })
        

        // 加载详情数据
        api.RouterHandlerService.routerHandlerGetRouterInfo({
            id: this.routerId
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
            this.loading = true;
            let reqData: api.server_api_UpdateRouterRequestData = {
                id: this.routerId,
                name: this.formState.name,
                path: this.formState.path,
                title: this.formState.title,
                is_menu: this.formState.is_menu,
                pid: this.formState.pid,
                sort: this.formState.sort,
                component: this.formState.component
            }

            api.RouterHandlerService.routerHandlerUpdateRouter({
                router: reqData,
            }).then((response) => {
                this.loading = false;
                if(response.code != 0) {
                    return Promise.reject(response.message)
                }
                message.success('操作完成')
            }).catch((err) => {
                this.loading = false;
                message.error('网络错误')
            })
        }
    }
})
</script>