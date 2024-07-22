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
            routerId: "",
            formState: ref<api.server_api_Router>({}),
            routerTreeData: ref<api.server_api_Router[]>([
                {
                    id: "0",
                    title: "根目录"
                }
            ]),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.routerId = String(this.$route.query.id)
        this.initData(this.routerId)
    },
    methods: {
        initData: async function(routerId: string) {
            this.loading = true
            try {
                // 加载数节点数据
                let routerTreeResponse = await api.RouterHandlerService.routerHandlerGetRouterList({})
                if (routerTreeResponse.code != 0) {
                    throw new Error(routerTreeResponse.message)
                }
                this.routerTreeData[0].children = routerTreeResponse.data?.list ?? []

                // 加载详情数据
                let infoResponse = await api.RouterHandlerService.routerHandlerGetRouterInfo({id: this.routerId})
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

            try {
                let response = await api.RouterHandlerService.routerHandlerUpdateRouter({router: reqData})
                if(response.code != 0) {
                    throw new Error(response.message)
                }

                await routerStore.updateRouter()
                message.success('操作完成')
            } catch (err) {
                message.error('网络错误')
            }
            this.loading = false;
        }
    }
})
</script>