<script setup lang="ts">
import { NForm, NFormItem, NInput, NSelect, NButton } from 'naive-ui';
import { computed, reactive } from 'vue';

import { isEmail, isPhone } from '../utilities/valid'
const user = reactive({
    email: '',
    pwd: '',
    type: 0,
})

function loginClick() {

}

var emailOrPhone = computed(
    (value: string) => {
        if (isEmail(value)) {
            user.type = 0;
            return true;
        } else if (isPhone(value)) {
            user.type = 1;
            return true;
        } else {
            return false;
        }
    }
)


</script>

<template>
    <n-form>
        <n-form-item label="电子邮件或手机号：" path="user.email">
            <n-input v-model:value="user.email" placeholder="输入电子邮箱或手机号" :feedback="emailOrPhone" />
        </n-form-item>

        <n-form-item label="密码" path="user.pwd">
            <n-input v-model:value="user.pwd" placeholder="输入密码" type="password" />
        </n-form-item>

        <n-button @click="loginClick">登陆</n-button>
    </n-form>
</template>

<style>
</style>