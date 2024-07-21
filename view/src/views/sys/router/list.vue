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
                    <a-popconfirm title="确认删除？" ok-text="确认" ok-type="danger" cancel-text="取消" @confirm="onDelete(record.id)">
                        <a-button danger size="small">删除</a-button>
                    </a-popconfirm>
                </span>
            </template>
        </template>
    </a-table>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import * as api from '@/api'
import { useRouterStore } from '@/pinia/modules/router';

const routerStore = useRouterStore()


export default defineComponent({
    data() {
        return {
            columns: [
                {title: '名称', dataIndex: 'name', key: 'name'},
                {title: '标题', dataIndex: 'title', key: 'title'},
                {title: '菜单', dataIndex: 'is_menu', key: 'is_menu'},
                {title: '路径', dataIndex: 'path', key: 'path'},
                {title: '操作', key: 'action'}
            ],
            data: Array<api.server_api_Router>(),
        };
    },
    created() {
        this.handleTableChange()
    },
    methods: {
        handleTableChange: function () {
            api.RouterHandlerService.routerHandlerGetRouterList({}).then(response => {
                if (response.code != 0) {
                    return Promise.reject(response.message)
                }
                return response.data
            }).then(data => {
                if (data?.list) {
                    this.data = this.formatData(data.list)
                }
            }).catch(err => {
                console.log('err:', err)
            })
        },
        formatData: function(list:api.server_api_Router[]) {
            // 格式化，如果children为空，就设置为undefined
            list.forEach(item => {
                if (item.children?.length) {
                    item.children = this.formatData(item.children)
                } else {
                    item.children = undefined
                }
            })
            return list
        },
        onDelete: function(id: string) {
            api.RouterHandlerService.routerHandlerDeleteRouter({
                ids: [id]
            }).then(response => {
                if (response.code != 0) {
                    return Promise.reject(response.message)
                }

                // 删除完成，刷新列表
                this.handleTableChange()
                routerStore.updateRouter()

            }).catch(err => {
                console.log('err:',err)
            })
        }
    }
});
</script>