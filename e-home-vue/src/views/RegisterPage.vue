<script setup lang="ts">

import { NForm, NFormItem, NInput, NSelect, NButton} from 'naive-ui';
import {reactive} from 'vue'
import axios from 'axios';
import config from '../config'
const user = reactive({
    name            : '',
    pwd             : '',
    pwd_confirm     : '',
    email           : '',
    phone           : '',
    type            : 0,
})

const options = [
    {
        label : "志愿者",
        value : 0,
    },
    {
        label : "子女",
        value : 1,
    }
]


function registerClick() {
    // axios.post('localhost:8989/api/register')
    // axios.get('http://0.0.0.0:8989/api/test').then(res => {console.log(res)})
    // TODO: is data valid
    const data = {
        email : user.email,
        pwd : user.pwd,
        phone : user.phone,
        name : user.name,
    }
    axios.post(config.SERVER + '/api/register', data).then(res => {console.log(res)})
}

</script>
<template>
<n-form>
    <n-form-item label="用户名" path="user.name">
        <n-input v-model:value="user.name" placeholder="输入姓名" />
    </n-form-item>

    <n-form-item label="电子邮件" path="user.email">
        <n-input v-model:value="user.email" placeholder="输入电子邮件地质"/>
    </n-form-item>
    
    <n-form-item label="电话号码" path="user.phone">
        <n-input v-model:value="user.phone" placeholder="输入电话号码"/>
    </n-form-item>
    
    <n-form-item label="密码" path="user.pwd">
        <n-input v-model:value="user.pwd" placeholder="输入密码"/>
    </n-form-item>

    <n-form-item label="确认密码" path="user.pwd_confirm">
        <n-input v-model:value="user.pwd_confirm" placeholder="输入确认密码"/>
    </n-form-item>

    <n-form-item label="选择你的类型" path="user.type">
        <n-select v-model:value="user.type" :options="options"/>
    </n-form-item>

    <n-button @click="registerClick"> 注册 </n-button>
</n-form>
{{user.name}}
</template>

<style>

</style>
