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
                        专为内部运营与管理场景设计的 B 端系统，沉淀通用能力，助力快速搭建业务后台。
                    </div>
                </div>

                <div class="login-card">
                    <div class="login-card-header">
                        <div class="login-title">欢迎登录</div>
                        <div class="login-subtitle">请使用公司账号登录后台管理系统</div>
                    </div>
                    <LoginComponent @onLoginSuccess="onLoginSuccess" />
                </div>
            </div>
        </a-layout-content>
        <a-layout-footer class="login-footer">
            © {{ new Date().getFullYear() }} Moumou Admin · 内部管理系统
        </a-layout-footer>
    </a-layout>
</template>

<script setup lang="ts">
import LoginComponent from '@/components/LoginComponent.vue'
import { useUserStore } from '@/pinia/modules/user'
import router from '@/router/index'

const userStore = useUserStore()

const redirectToDefaultPage = () => {
    let redirect = router.currentRoute.value.query.redirect?.toString() ?? ''
    if (redirect === '') {
        redirect = '/'
    }
    router.replace({ path: redirect })
}

const onLoginSuccess = function (ret: any) {
    console.log(ret)
    userStore.SetToken(ret.token)
    redirectToDefaultPage()
}
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

.login-footer {
    text-align: center;
    color: #bfbfbf;
    font-size: 12px;
    padding: 16px 0 24px;
    background: transparent;
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