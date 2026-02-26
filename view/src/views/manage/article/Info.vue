<template>
    <ContentPage>
        <template #content>
            <a-spin :spinning="loading">
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-form-item label="ID">
                    <a-input v-model:value="formState.id" disabled />
                </a-form-item>

                <a-form-item label="标题">
                    <a-input v-model:value="formState.title" />
                </a-form-item>

                <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                    <a-button
                        type="primary"
                        @click="onSubmit"
                        :loading="loading"
                        v-permission="'ManageArticleWrite'"
                    >
                        提交
                    </a-button>
                    <a-button style="margin-left: 10px" @click="$router.back()">
                        返回
                    </a-button>
                </a-form-item>

                <a-form-item :wrapper-col="{ span: 20, offset: 2 }">
                    <MarkdownEditor
                        v-model:value="formState.content"
                        :rows="20"
                    />
                </a-form-item>
            </a-form>
            </a-spin>
        </template>
    </ContentPage>
</template>

<script setup lang="ts">
import ContentPage from '@/components/ContentPage.vue';
import MarkdownEditor from '@/components/MarkdownEditor.vue';
</script>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { message } from 'ant-design-vue';
import * as api from '@/api';

export default defineComponent({
    data() {
        return {
            dataId: '',
            formState: ref<api.server_api_Article>({}),
            loading: ref<boolean>(false),
        };
    },
    created() {
        this.dataId = String(this.$route.query.id || '');
        if (!this.dataId) {
            message.error('缺少文章 ID');
            return;
        }
        this.initData(this.dataId);
    },
    methods: {
        async initData(id: string) {
            this.loading = true;
            try {
                const resp = await api.ArticleHandlerService.articleHandlerGetArticleInfo({
                    id,
                });
                if (resp.code !== 0) {
                    throw new Error(resp.message);
                }
                this.formState = resp.data ?? {};
            } catch (err) {
                console.error('info error:', err);
                message.error('加载详情失败');
            } finally {
                this.loading = false;
            }
        },
        async onSubmit() {
            if (!this.formState.id) {
                message.error('缺少文章 ID');
                return;
            }
            this.loading = true;
            const reqData: api.server_api_UpdateArticleRequestData = {
                id: this.formState.id,
                title: this.formState.title,
                content: this.formState.content,
            };
            try {
                const resp = await api.ArticleHandlerService.articleHandlerUpdateArticle({
                    article: reqData,
                });
                if (resp.code !== 0) {
                    throw new Error(resp.message);
                }
                message.success('保存成功');
            } catch (err) {
                console.error('update error:', err);
                message.error('保存失败');
            } finally {
                this.loading = false;
            }
        },
    },
});
</script>

