<template>
    <a-row>
        <a-col :span="12" style="margin-bottom: 10px;">
            <a-button type="primary" href="add" v-permission="'ManageUserWrite'">添加</a-button>
        </a-col>
    </a-row>
    <a-table :columns="columns" :data-source="data" :pagination="pagination" @change="handleTableChange">
        <template #headerCell="{ column }"></template>

        <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'action'">
                <span>
                    <a-button size="small" :href="`info?id=${record.id}`" style="margin-right:5px;">详情</a-button>
                    <a-button danger size="small">删除</a-button>
                </span>
            </template>
        </template>
    </a-table>
</template>
<script lang="ts">
import { defineComponent, computed, ref } from 'vue';
import type { TableColumnType } from 'ant-design-vue';
import * as api from '@/api'

export default defineComponent({
    data() {
        return {
            columns: [
                {
                    title: "账号",
                    dataIndex: "username"
                },
                {
                    title: 'Action',
                    key: 'action',
                }
            ],
            data: Array<api.server_api_User>(),
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

            api.UserHandlerService.userHandlerGetUserList({
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
        }
    }
})

</script>