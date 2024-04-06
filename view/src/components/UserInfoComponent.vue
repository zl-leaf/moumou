<template>

  <a-dropdown :placement="'bottom'">
    <a-avatar style="background-color: #87d068">
      <template #icon>
        <UserOutlined></UserOutlined>
      </template>
    </a-avatar>
    <template #overlay>
      <a-menu>
        <a-menu-item>
          <a href="javascript:;">个人信息</a>
        </a-menu-item>
        <a-menu-item>
          <a href="javascript:;" @click="Logout">注销</a>
        </a-menu-item>
      </a-menu>
    </template>
  </a-dropdown>
</template>
<script lang="ts" setup>
import { UserOutlined } from '@ant-design/icons-vue';
import { defineComponent, ref } from 'vue';
import {useUserStore} from '@/pinia/modules/user'

const Logout = () => {
  const userStore = useUserStore()
  userStore.Logout()
}

</script>