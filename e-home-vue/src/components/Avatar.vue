<script setup lang="ts">
import { NAvatar } from 'naive-ui';

const props = defineProps({
    avatarUrl : String,
    username : String,
})
</script>

<template>
    <n-avatar 
    round
    :size="40"
    :src="avatarUrl"
    />
    <span class="username"> {{username}} </span>
</template>

<style>
.username{
    font-size: 20px;
}
</style>