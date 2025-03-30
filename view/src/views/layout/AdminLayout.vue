<template>
  <a-layout style="min-height: 100vh">
    <SiderComponent v-model:renderRouters="renderRouters"/>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0 16px" >
        <a-row>
          <a-col :span="23">
            <BreadcrumbComponent></BreadcrumbComponent>
          </a-col>
          <a-col :span="1">
            <UserInfoComponent></UserInfoComponent>
          </a-col>
        </a-row>
      </a-layout-header>
      <a-layout-content style="background-color: #fff;margin: 10px 16px">
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Ant Design ©2018 Created by Ant UED
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import SiderComponent from '@/components/SiderComponent.vue'
import BreadcrumbComponent from '@/components/BreadcrumbComponent.vue'
import UserInfoComponent from '@/components/UserInfoComponent.vue'
import router from '@/router'

let layoutRouter = router.getRoutes().find(item => item.name == 'admin_layout')
const renderRouters = layoutRouter?.children || []
</script>