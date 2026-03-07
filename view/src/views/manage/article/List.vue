<template>
    <ContentPage>
        <template #extra>
            <a-button
                type="primary"
                href="article/add"
                v-permission="'ManageArticleWrite'"
            >
                添加
            </a-button>
        </template>

        <template #content>
            <a-table
                :columns="columns"
                :data-source="data"
                :pagination="pagination"
                :loading="loading"
                row-key="id"
                @change="handleTableChange"
            >
                <template #bodyCell="{ column, record }">
                    <template v-if="column.key === 'action'">
                        <span>
                            <a-button
                                size="small"
                                :href="`article/info?id=${record.id}`"
                                style="margin-right: 5px"
                            >
                                详情
                            </a-button>
                            <a-popconfirm
                                title="确定删除该文章吗？"
                                ok-text="确定"
                                cancel-text="取消"
                                @confirm="onDelete(record)"
                            >
                                <a-button danger size="small" v-permission="'ManageArticleWrite'">
                                    删除
                                </a-button>
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
import { defineComponent } from 'vue';
import * as api from '@/api';
import { message } from 'ant-design-vue';

export default defineComponent({
    data() {
        return {
            columns: [
                {
                    title: 'ID',
                    dataIndex: 'id',
                },
                {
                    title: '标题',
                    dataIndex: 'title',
                },
                {
                    title: 'Action',
                    key: 'action',
                    width: 180,
                },
            ],
            data: Array<api.server_api_Article>(),
            pagination: {
                total: 0,
                current: 1,
                pageSize: 10,
            },
            title: this.$route.meta.title,
            loading: false,
        };
    },
    created() {
        this.handleTableChange(this.pagination, {}, {});
    },
    methods: {
        handleTableChange(pag: any, filters: any, sorter: any) {
            this.pagination.current = pag.current;
            this.pagination.pageSize = pag.pageSize;
            this.fetchList();
        },
        async fetchList() {
            this.loading = true;
            try {
                const resp = await api.ArticleHandlerService.articleHandlerGetArticleList({
                    currentPage: this.pagination.current,
                    pageSize: this.pagination.pageSize,
                    filter: {},
                });
                if (resp.code !== 0) {
                    throw new Error(resp.message);
                }
                const data = resp.data;
                this.pagination.total = Number(data?.total ?? 0);
                this.data = data?.list ?? [];
            } catch (err) {
                console.error('fetchList error:', err);
                message.error('加载文章列表失败');
            } finally {
                this.loading = false;
            }
        },
        async onDelete(record: api.server_api_Article) {
            if (!record.id) {
                message.error('记录缺少 ID');
                return;
            }
            try {
                const resp = await api.ArticleHandlerService.articleHandlerDeleteArticle({
                    ids: [record.id],
                });
                if (resp.code !== 0) {
                    throw new Error(resp.message);
                }
                message.success('删除成功');
                this.fetchList();
            } catch (err) {
                console.error('delete error:', err);
                message.error('删除失败');
            }
        },
    },
});
</script>

