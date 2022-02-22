<script setup lang="ts">
import { RouterView, useRouter, useRoute } from 'vue-router';
import {NConfigProvider, zhCN, dateZhCN, NSpace, NLayout, NLayoutHeader,
        NGrid, NGi, NButton, NLayoutContent} from 'naive-ui';
import avatar from './components/Avatar.vue'
const router = useRouter();
const route = useRoute();
function clickAvatar() {
  router.push({path:'/login'})
}
function clickTitle() {
  router.push({
    path : '/test',
  })
}

</script>

<template>
  <!-- <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-space vertical size="large">
      <n-layout>
        <n-layout-header>
          <n-grid :x-gap="12" item-responsive responsive="screen">
            <n-gi span="10 l:4">
            <div class="title" @click.native="clickTitle">
              <h1> 社区 E 家 </h1> 
            </div>
            </n-gi>
            <n-gi span="6 l:15">
              <n-space>
                <menu>
                </menu>
              </n-space>
            </n-gi>
            <n-gi span="5">
            <avatar avatar-url="/favicon.ico" username="未登陆" @click.native="clickAvatar"/>
            </n-gi>
          </n-grid>
        </n-layout-header>
        <n-layout-content content-style="padding: 24px;">
          <router-view/>
        </n-layout-content>
      </n-layout>
    </n-space>
  </n-config-provider> -->

  <router-view/>
  
</template>

<style>
.title {
  background: rgb(121, 121, 243);
  cursor: pointer;
}
</style>
