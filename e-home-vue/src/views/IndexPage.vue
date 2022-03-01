<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "../stores/user";
import { NLayout, NLayoutHeader, NLayoutFooter, NSpace } from "naive-ui";
const store = useUserStore();
document.title = "社区 E 家";
</script>

<template>
  <n-space vertical size="large">
    <n-layout>
      <n-layout-header> </n-layout-header>
      <n-layout-content content-style="padding: 24px;"> {{ store }} </n-layout-content>
      <n-layout-footer>成府路</n-layout-footer>
    </n-layout>
  </n-space>
</template>
<style></style>
