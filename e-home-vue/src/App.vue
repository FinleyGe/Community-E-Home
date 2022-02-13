<script setup lang="ts">
import { RouterView, useRouter, useRoute } from 'vue-router';
import {NConfigProvider, zhCN, dateZhCN, NSpace, NLayout, NLayoutHeader,
        NGrid, NGi, NButton} from 'naive-ui';
import avatar from './components/Avatar.vue'
const router = useRouter;
const route = useRoute;
function clickAvatar() {
  // todo: unfinished.
}

</script>

<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-space vertical size="large">
      <n-layout>
        <n-layout-header>
          <n-grid :x-gap="12" item-responsive responsive="screen">
            <n-gi span="4">

            </n-gi>
            <n-gi span="16">
              <h1> 社区 E 家 </h1>
            </n-gi>
            <n-gi span="4">
            <avatar avatar-url="/favicon.ico" username="未登陆" @click.native="clickAvatar"/>
            </n-gi>
          </n-grid>
        </n-layout-header>
        <n-layout-content content-style="padding: 24px;">
          <router-view/>
        </n-layout-content>
      </n-layout>
    </n-space>
  </n-config-provider>
</template>

<style>
</style>
