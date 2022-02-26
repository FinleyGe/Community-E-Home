<script setup lang="ts">

</script>

<template>
<div class="notice">
    <div class="noticeContext">
        <slot></slot>
    </div>
</div>
</template>

<style>
.notice {
    width: 100%;
    height: 100%;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 999;
}

.noticeContext {
    position: absolute;
    top: 120px;
    left: 10px;
    padding: 5px;
    background-color: #ffffff;
    box-shadow: 3px 2px 5px #7777;
}
</style>