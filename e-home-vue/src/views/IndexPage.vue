<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "../stores/user";
const store = useUserStore();
// const isLogin = ref(store.isLoggedIn);
const jwt = ref(store.jwt);
</script>

<template>
  {{ store }}
  <!-- <p>isLogin: {{ isLogin ? "logged" : "not logged" }}</p> -->
  <!-- <p>jwt: {{ jwt }}</p> -->
</template>
<script></script>
