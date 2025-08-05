<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-tree
                    checkable
                    showLine
                    :tree-data="permissionTree"
                    v-model:checkedKeys="formState.values"
                ></a-tree>

                <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                    <a-button type="primary" @click="onSubmit" :loading="loading">提交</a-button>
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
import { message } from 'ant-design-vue';
import { useUserStore } from '@/pinia/modules/user';
import type { TreeDataItem } from 'ant-design-vue/es/tree/Tree';


type FormState = {
    values: string[]
}

export default defineComponent({
    data() {
        return {
            dataId: "",
            permissionTree: ref<TreeDataItem[]>([]),
            formState: ref<FormState>({
                values: []
            }),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function(dataId: string) {
            let that = this
            this.loading = true
            try {
                // 加载详情数据
                let permissionTreeResponse = await api.PermissionHandlerService.permissionHandlerGetPermissionTree({})
                if (permissionTreeResponse.code != 0) {
                    throw new Error(permissionTreeResponse.message)
                }
                this.permissionTree = this.formatPermissionToTreeDataItem(permissionTreeResponse.data?.list ?? [])

                let getRolePermissionResponse = await api.RoleHandlerService.roleHandlerGetRolePermission({
                    roleId: this.dataId,
                })
                if (getRolePermissionResponse.code !== 0) {
                    throw new Error(getRolePermissionResponse.message)
                }
                getRolePermissionResponse.data?.list?.forEach(permission => {
                    that.formState.values.push(String(permission.id))
                })

                this.loading = false
            } catch (err) {
                message.error("网络错误")
            }

        },
        onSubmit: async function() {
            this.loading = true
            try {
                let response = await api.RoleHandlerService.roleHandlerUpdateRolePermission({
                    id: this.dataId,
                    permissionIds: this.formState.values
                })
                if(response.code != 0) {
                    throw new Error(response.message)
                }

                await useUserStore().UpdatePermissions()

                message.success('操作完成')
                this.loading = false
            } catch(err) {
                message.error("网络异常")
            }
        },
        formatPermissionToTreeDataItem(permissionTreeNodes: api.server_api_PermissionTreeNode[]): Array<TreeDataItem> {
            const that = this
            return permissionTreeNodes.map(permissionTreeNode => {
                let treeDataItem: TreeDataItem = {
                    title: permissionTreeNode.name,
                    key: permissionTreeNode.id ?? '',
                }
                if (permissionTreeNode.children?.length) {
                    treeDataItem.children = that.formatPermissionToTreeDataItem(permissionTreeNode.children)
                }
                return treeDataItem
            })
        }
    }
})
</script>