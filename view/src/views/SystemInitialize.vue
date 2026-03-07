<template>
    <a-layout class="login-layout">
        <a-layout-content>
            <div class="login-page">
                <div class="login-slogan">
                    <div class="brand">
                        <div class="brand-logo">M</div>
                        <div class="brand-text">
                            <div class="brand-name">Moumou Admin</div>
                            <div class="brand-subtitle">统一管理 · 高效协作 · 安全可控</div>
                        </div>
                    </div>
                    <div class="slogan-desc">
                        首次使用需完成系统初始化：填写认证信息与管理员账号，提交后将初始化数据表并创建管理员。
                    </div>
                </div>

                <div class="login-card">
                    <div class="login-card-header">
                        <div class="login-title">系统初始化</div>
                        <div class="login-subtitle">填写认证信息与管理员账号后提交</div>
                    </div>
                    <a-spin :spinning="loading">
                        <a-form
                            :model="formState"
                            layout="vertical"
                            @finish="onSubmit"
                        >
                            <a-form-item
                                label="认证账号"
                                name="username"
                                :rules="[{ required: true, message: '请输入认证账号' }]"
                            >
                                <a-input v-model:value="formState.username" placeholder="与 config 中 system.username 一致" />
                            </a-form-item>
                            <a-form-item
                                label="认证密码"
                                name="password"
                                :rules="[{ required: true, message: '请输入认证密码' }]"
                            >
                                <a-input-password v-model:value="formState.password" placeholder="与 config 中 system.password 一致" />
                            </a-form-item>
                            <a-form-item
                                label="管理员账号"
                                name="adminUsername"
                                :rules="[{ required: true, message: '请输入管理员账号' }]"
                            >
                                <a-input v-model:value="formState.adminUsername" placeholder="初始化后用于登录的管理员用户名" />
                            </a-form-item>
                            <a-form-item
                                label="管理员密码"
                                name="adminPassword"
                                :rules="[{ required: true, message: '请输入管理员密码' }]"
                            >
                                <a-input-password v-model:value="formState.adminPassword" placeholder="初始化后用于登录的管理员密码" />
                            </a-form-item>
                            <a-form-item>
                                <a-button type="primary" html-type="submit" :loading="loading" block size="large">
                                    提交初始化
                                </a-button>
                            </a-form-item>
                        </a-form>
                    </a-spin>
                </div>
            </div>
        </a-layout-content>
        <AppFooter minimal />
    </a-layout>
</template>

<script lang="ts">
import AppFooter from '@/components/AppFooter.vue'
import { defineComponent, reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import router from '@/router'
import * as api from '@/api'

export default defineComponent({
    components: { AppFooter },
    data() {
        return {
            formState: reactive<api.server_api_InitializeRequest>({
                username: '',
                password: '',
                adminUsername: '',
                adminPassword: '',
            }),
            loading: ref(false),
        }
    },
    methods: {
        async onSubmit() {
            this.loading = true
            try {
                const res = await api.SystemHandlerService.systemHandlerInitialize(this.formState)
                if (res.code === 0) {
                    message.success(res.message ?? '初始化成功，请使用管理员账号登录')
                    router.replace({ path: '/login' })
                } else {
                    message.error(res.message ?? '初始化失败')
                }
            } catch (err: any) {
                message.error(err?.message ?? err?.body?.message ?? '请求失败')
            } finally {
                this.loading = false
            }
        },
    },
})
</script>

<style scoped>
.login-layout {
    min-height: 100vh;
    background: radial-gradient(circle at 0 0, #f0f5ff 0, #f5f5f5 40%, #fafafa 100%);
}

.login-page {
    min-height: calc(100vh - 80px);
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 80px;
    padding: 40px 80px;
    box-sizing: border-box;
}

.login-slogan {
    max-width: 460px;
}

.brand {
    display: flex;
    align-items: center;
    margin-bottom: 24px;
}

.brand-logo {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    background: linear-gradient(135deg, #1677ff, #36cfc9);
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    font-weight: 700;
    font-size: 24px;
    box-shadow: 0 10px 30px rgba(22, 119, 255, 0.35);
}

.brand-text {
    margin-left: 16px;
}

.brand-name {
    font-size: 22px;
    font-weight: 600;
    color: #1f1f1f;
}

.brand-subtitle {
    margin-top: 4px;
    font-size: 13px;
    color: #8c8c8c;
}

.slogan-desc {
    margin-top: 16px;
    font-size: 14px;
    color: #595959;
    line-height: 1.7;
}

.login-card {
    width: 380px;
    padding: 32px 32px 28px;
    border-radius: 16px;
    background: #fff;
    box-shadow:
        0 18px 45px rgba(0, 0, 0, 0.06),
        0 0 0 1px rgba(0, 0, 0, 0.03);
    box-sizing: border-box;
}

.login-card-header {
    margin-bottom: 24px;
}

.login-title {
    font-size: 22px;
    font-weight: 600;
    color: #1f1f1f;
}

.login-subtitle {
    margin-top: 6px;
    font-size: 13px;
    color: #8c8c8c;
}

@media (max-width: 960px) {
    .login-page {
        flex-direction: column;
        padding: 32px 24px;
        gap: 32px;
        align-items: stretch;
    }

    .login-slogan {
        max-width: none;
        text-align: center;
    }

    .brand {
        justify-content: center;
    }

    .login-card {
        margin: 0 auto;
        width: 100%;
        max-width: 380px;
    }
}
</style>
