<template>
    <a-row>
        <a-col :span="12" style="margin-bottom: 10px;">
            <a-button type="primary" href="add" v-permission="'ManageRoleWrite'">添加</a-button>
        </a-col>
    </a-row>
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="handleTableChange">
        <template #headerCell="{ column }"></template>

        <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
                <span>
                    <a-button size="small" :href="`info?id=${record.id}`" style="margin-right:5px;">详情</a-button>
                    <a-button size="small" :href="`permission?id=${record.id}`" style="margin-right:5px;" v-permission="'ManageRoleWrite'">权限</a-button>
                    <a-button size="small" :href="`bind_user?id=${record.id}`" style="margin-right:5px;" v-permission="'ManageRoleWrite'">授权</a-button>
                    <a-popconfirm title="确认删除？" ok-text="确认" ok-type="danger" cancel-text="取消" @confirm="onDelete(record.id)">
                        <a-button danger size="small" v-permission="'ManageRoleWrite'">删除</a-button>
                    </a-popconfirm>
                </span>
            </template>
        </template>
    </a-table>
</template>
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
                    title: 'Action',
                    key: 'action',
                }
            ],
            data: Array<api.server_api_Role>(),
            pagination: {
                total: 0,
                current: 1,
                pageSize: 10,
            },
        }
    },
    created() {
        this.handleTableChange(this.pagination, {}, {})
    },
    methods: {
        handleTableChange: function (pag: any, filters: any, sorter: any) {
            this.pagination.current = pag.current
            this.pagination.pageSize = pag.pageSize
            console.log('change', pag, filters, sorter)

            api.RoleHandlerService.roleHandlerGetRoleList({
                currentPage: pag.current,
                pageSize: pag.pageSize,
                filter: {},
            }).then((response) => {
                if (response.code != 0) {
                    return Promise.reject(response.message)
                }
                return response.data
            }).then((data) => {
                this.pagination.total = Number(data?.total ?? 0)
                this.data = data?.list ?? []
            }).catch(err => {
                console.log('err:', err)
            })
        },
        onDelete: async function(id: string) {
            try {
                let response = await api.RoleHandlerService.roleHandlerDeleteRole({
                    ids: [id]
                })
                if(response.code != 0) {
                    throw new Error(response.message)
                }
                this.handleTableChange(this.pagination, {}, {})
            } catch(err) {
                message.error("删除失败")
            }
        }
    }
})

</script>