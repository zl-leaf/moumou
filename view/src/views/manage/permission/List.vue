<template>
    <ContentPage>
        <template #extra>
            <a-button type="primary" href="permission/add" v-permission="'ManagePermissionWrite'">添加</a-button>
        </template>

        <template #content>
            <a-table :columns="columns" :data-source="data" :pagination="false" @change="handleTableChange">
                <template #headerCell="{ column }"></template>

                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'action'">
                        <span>
                            <a-button size="small" :href="`permission/info?id=${record.id}`" style="margin-right:5px;">详情</a-button>
                            <a-popconfirm title="确认删除？" ok-text="确认" ok-type="danger" cancel-text="取消" @confirm="onDelete(record.id)">
                                <a-button danger size="small" v-permission="'ManagePermissionWrite'">删除</a-button>
                            </a-popconfirm>
                        </span>
                    </template>
                </template>
            </a-table>
        </template>
    </ContentPage>
</template>
<script setup lang="ts">
import ContentPage from '@/components/ContentPage.vue';
</script>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue';
import { message, type TableColumnType } from 'ant-design-vue';
import * as api from '@/api'
export default defineComponent({
    data() {
        return {
            columns: [
                {
                    title: "名称",
                    dataIndex: "name"
                },
                {
                    title: "编码",
                    dataIndex: "code"
                },
                {
                    title: 'Action',
                    key: 'action',
                }
            ],
            data: Array<api.server_api_Permission>(),
        }
    },
    created() {
        this.handleTableChange({}, {})
    },
    methods: {
        handleTableChange: function (filters: any, sorter: any) {

            api.PermissionHandlerService.permissionHandlerGetPermissionTree({
                filter: {},
            }).then((response) => {
                if (response.code != 0) {
                    return Promise.reject(response.message)
                }
                return response.data
            }).then((data) => {
                this.data = data?.list ?? []
            }).catch(err => {
                console.log('err:', err)
            })
        },
        onDelete: async function(id: string) {
            api.PermissionHandlerService.permissionHandlerDeletePermission({
                ids: [id]
            }).then((response) => {
                if (response.code != 0) {
                    return Promise.reject(response.message)
                }
                message.success('删除成功')
                this.handleTableChange({}, {})
            }).catch(err => {
                console.log('err:', err)
            })
        }
    }
})
</script> 