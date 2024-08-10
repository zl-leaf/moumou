<template>
  <a-layout style="min-height: 100vh">
    <SiderComponent @openMenu="openMenu" />
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0" >
        <a-row>
          <a-col :span="23"></a-col>
          <a-col :span="1">
            <UserInfoComponent></UserInfoComponent>
          </a-col>
        </a-row>
      </a-layout-header>
      <a-layout-content style="margin: 0 16px">
        <BreadcrumbComponent></BreadcrumbComponent>
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

import { defineComponent, ref } from 'vue';
import { useUserStore } from "@/pinia/modules/user"
import router from '@/router';

const openMenu = (e: any) => {
  console.log(e)
  router.push({
    name: e.key,
    replace: true,
  })
}

const userStore = useUserStore()
</script>