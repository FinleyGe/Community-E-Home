<script setup lang="ts">
import { computed, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import {key} from '../store'
import {IUserLoginInfo, login} from '../utilities/apis'
import { isEmailOrPhone } from '../utilities/valid';

const store = useStore(key)
const router = useRouter()
const loginInfo: IUserLoginInfo = reactive({
    emailPhone : '',
    method : computed(() => isEmailOrPhone(loginInfo.emailPhone)),
    pwd : '',
})

function RegisterClicked() {
    router.push('/register');
}

function LoginClicked() {
    if (loginInfo.method == -1){
        alert("请检查您的手机号或电子邮箱！")
    }else{
        login(loginInfo)
    }
}
</script>

<template>
    <div style="width: 70%; margin: 0 auto; padding-top: 20em;padding-bottom: 20em;">
        <div class="ui stackable centered column grid">
            <div class="centered row">
                <div class="six wide column">
                    <div
                        style="padding-bottom: 5em; background-color: rgba(245, 245, 245, 0.5);"
                        class="ui segment"
                    >
                        <form class="ui form" onsubmit="return false;">
                            <h3 style="color:  #1890ff; text-align:center;">用户登录</h3>
                            <div class="field">
                                <label for="username">账号：</label>
                                <input
                                    type="text"
                                    name="username"
                                    id="username"
                                    placeholder="请输入手机号或邮箱"
                                    v-model="loginInfo.emailPhone"
                                />
                            </div>
                            <div class="field">
                                <label for="password">密码：</label>
                                <input
                                    type="password"
                                    name="password"
                                    id="password"
                                    placeholder="请输入密码"
                                    v-model="loginInfo.pwd"
                                />
                            </div>
                            <button 
                            class="ui inverted green button"
                            style="float: right;"
                            @click.native="RegisterClicked">注册</button>
                            <button 
                            class="ui inverted blue button"
                            style="float: right;"
                            @click.native="LoginClicked"
                            :disabled="loginInfo.method == -1 || loginInfo.pwd == ''">登录</button>
                        </form>
                        {{loginInfo}}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
</style>