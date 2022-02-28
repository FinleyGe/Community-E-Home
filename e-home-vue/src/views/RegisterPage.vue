<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { isPwdValid } from "../utilities/valid";

const UserData = reactive({
  name: "",
  pwd: "",
  phone: "",
  email: "",
  type: 0,
});

const pwdRepeat = ref("");
const isPwdSame = ref(false);
const pwdValid = ref(false);
const router = useRouter();
function registerClick() {
  if (!(isPwdSame.value && pwdValid.value)) {
    alert("请重新检查您的密码！");
  } else {
    // register(UserData);
  }
}
function CheckPwd() {
  isPwdSame.value = UserData.pwd == pwdRepeat.value;
  pwdValid.value = isPwdValid(UserData.pwd);
}
function CheckInfo() {}

function cancelClick() {
  router.push("/index");
}
</script>

<template>
  <div style="width: 70%; margin: 0 auto; padding-top: 20em; padding-bottom: 20em">
    <div class="ui stackable centered column grid">
      <div class="centered row">
        <div class="six wide column">
          <div
            style="padding-bottom: 5em; background-color: rgba(245, 245, 245, 0.5)"
            class="ui segment"
          >
            <form id="regist" class="ui form" onsubmit="return false;">
              <h3 style="color: #1890ff; text-align: center">用户注册</h3>
              <!-- <notice> asd asd </notice> -->
              <div class="field">
                <label for="username">用户名：</label>
                <input
                  v-model="UserData.name"
                  type="text"
                  name="username"
                  id="username"
                  placeholder="请输入用户名"
                />
              </div>
              <div class="field">
                <label for="password">密码：</label>
                <input
                  v-model="UserData.pwd"
                  @input="CheckPwd"
                  type="password"
                  name="password"
                  id="password"
                  placeholder="请输入密码"
                />
                <label class="notice">{{ pwdValid ? "合法" : "密码过于简单" }}</label>
              </div>
              <div class="field">
                <label for="confirm_password">重复密码：</label>
                <input
                  v-model="pwdRepeat"
                  @input="CheckPwd"
                  type="password"
                  name="confirm_password"
                  id="confirm_password"
                  placeholder="重复输入密码"
                />
                <label class="notice">{{ isPwdSame ? "合法" : "两次输入不同" }}</label>
              </div>
              <div class="field">
                <label for="email">邮箱：</label>
                <input
                  v-model="UserData.email"
                  type="email"
                  name="email"
                  id="email"
                  placeholder="请输入邮箱"
                />
              </div>
              <div class="field">
                <label for="phone">手机号码：</label>
                <input
                  v-model="UserData.phone"
                  type="text"
                  name="phone"
                  id="phone"
                  placeholder="请输入手机号码"
                />
              </div>
              <div class="field">
                <label for="identity">选择您的身份</label>
                <select v-model="UserData.type" name="identity" id="identity">
                  <option value="0">志愿者</option>
                  <option value="1">子女</option>
                </select>
              </div>
            </form>
            <!-- <p> {{ UserData }}</p> -->
            <button
              style="float: right"
              class="ui inverted green button"
              @click.native="registerClick"
              :disabled="!(pwdValid && isPwdSame)"
            >
              注册
            </button>
            <button
              style="float: right"
              class="ui inverted blue button"
              @click.native="cancelClick"
            >
              取消
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.notice {
  color: red;
}
</style>
