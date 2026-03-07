<template>
  <a-layout-footer :class="['app-footer', { 'app-footer--minimal': minimal }]">
    © {{ new Date().getFullYear() }} Moumou Admin · 内部管理系统
  </a-layout-footer>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ minimal?: boolean }>(), { minimal: false })
</script>

<style scoped>
.app-footer {
  text-align: center;
}

.app-footer--minimal {
  color: #bfbfbf;
  font-size: 12px;
  padding: 16px 0 24px;
  background: transparent;
}
</style>
