<template>
  <div class="markdown-editor">
    <MdEditor
      v-model="innerValue"
      :preview="preview"
      :theme="theme"
      :language="language"
      :preview-theme="previewTheme"
      :style="{ height }"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const props = withDefaults(
  defineProps<{
    value?: string;
    height?: string;
    preview?: boolean;
    theme?: 'light' | 'dark';
    previewTheme?: 'default' | 'github' | 'vuepress' | 'mk-cute';
    language?: 'zh-CN' | 'en-US';
  }>(),
  {
    value: '',
    height: '400px',
    preview: true,
    theme: 'light',
    previewTheme: 'default',
    language: 'zh-CN',
  }
);

const emit = defineEmits<{
  (e: 'update:value', value: string): void;
}>();

const innerValue = ref(props.value ?? '');

const { preview, theme, previewTheme, language, height } = props;

watch(
  () => props.value,
  (val) => {
    innerValue.value = val ?? '';
  }
);

watch(innerValue, (val) => {
  emit('update:value', val);
});
</script>

<style scoped>
.markdown-editor {
  border: 1px solid #f0f0f0;
  border-radius: 4px;
  overflow: hidden;
}
</style>

