<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-form-item
                    label="标题"
                    :rules="[{ required: true, message: '请输入标题' }]"
                >
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
import router from '@/router';

export default defineComponent({
    data() {
        return {
            formState: ref<api.server_api_CreateArticleRequestData>({
                title: '',
                content: '',
            }),
            loading: ref<boolean>(false),
        };
    },
    methods: {
        async onSubmit() {
            if (!this.formState.title || !this.formState.content) {
                message.error('请填写必填项');
                return;
            }

            this.loading = true;
            const reqData: api.server_api_CreateArticleRequestData = {
                title: this.formState.title,
                content: this.formState.content,
            };
            try {
                const resp = await api.ArticleHandlerService.articleHandlerCreateArticle({
                    article: reqData,
                });
                if (resp.code !== 0) {
                    throw new Error(resp.message);
                }
                message.success('创建成功');
                router.back();
            } catch (err) {
                console.error('create error:', err);
                message.error('网络错误或创建失败');
            } finally {
                this.loading = false;
            }
        },
    },
});
</script>

