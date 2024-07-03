<template>
    <a-row>
        <a-col :span="12" style="margin-bottom: 10px;">
            <a-button type="primary" href="add">添加</a-button>
        </a-col>
    </a-row>
    <a-table :columns="columns" :data-source="data" rowKey="id">
        <template #bodyCell="{ column, record }">
            <template v-if="column.key == 'is_menu'">
                {{ record.is_menu ? "是":"否" }}
            </template>

            <template v-if="column.key === 'action'">
                <span>
                    <a-button type="primary" size="small" :href="`info?id=${record.id}`" style="margin-right:5px;">详情</a-button>
                    <a-button danger size="small">删除</a-button>
                </span>
            </template>
        </template>
    </a-table>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { getSysRouterListAPI, type RouterModel, } from '@/api/manage/router';

export default defineComponent({
    data() {
        return {
            columns: [
                {title: '名称', dataIndex: 'name', key: 'name'},
                {title: '标题', dataIndex: 'title', key: 'title'},
                {title: '菜单', dataIndex: 'is_menu', key: 'is_menu'},
                {title: '路径', dataIndex: 'path', key: 'path'},
                {title: 'Action', key: 'action'}
            ],
            data: Array<RouterModel>(),
        };
    },
    created() {
        this.handleTableChange()
    },
    methods: {
        handleTableChange: function () {
            getSysRouterListAPI().then(response => {
                if (response.code != '0') {
                    return Promise.reject(response.message)
                }
                return response.data
            }).then(data => {
                this.data = data.list
            }).catch(err => {
                console.log('err:', err)
            })
        }
    }
});
</script>