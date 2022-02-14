<script setup lang="ts">
import { NAvatar, NSpace } from 'naive-ui';

const props = defineProps({
    avatarUrl : String,
    username : String,
})
</script>

<template>
<n-space justify="center">
    <n-avatar 
    round
    :size="40"
    :src="avatarUrl"
    />
    <span class="username"> {{username}} </span>
    </n-space>
</template>

<style>
.username{
    font-size: 16px;
    cursor: pointer;
}
</style>