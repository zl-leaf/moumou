<template>
    <a-page-header :title="title" style="padding-top: 10px;padding-bottom: 10px;">
        <template #extra>
            <slot name="extra"></slot>
        </template>
    </a-page-header>
    <a-divider style="margin-top: 0;"/>
    <div style="padding: 0 15px;">
        <slot name="content"></slot>
    </div>
</template>
<style scoped>
.ant-page-header-heading-title {
    font-weight: normal !important;
}
</style>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    data() {
        return {
            title: this.$route.meta.title
        }
    }
})
</script>