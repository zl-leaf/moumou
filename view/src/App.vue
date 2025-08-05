<template>
  <a-config-provider
    :theme="{
      token: {
        borderRadius: 2,
      },
    }"
  >
  <RouterView />
  </a-config-provider>
</template>
<script lang="ts" setup>
import { RouterView } from 'vue-router'
</script>