<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { isPwdValid } from "../utilities/valid";
import { IRegisterInfo, IVerifyInfo } from "../types/user";
import { RegisterApi, VerifySendApi } from "../apis/user";
// const UserData = reactive({
//   name: "",
//   pwd: "",
//   phone: "",
//   email: "",
//   type: 0,
//   code: "",
// });

const UserData: IRegisterInfo = reactive({
  name: "",
  pwd: "",
  phone: "",
  email: "",
  type: 0,
  code: "",
});
const verifyCode: IVerifyInfo = reactive({
  email: "",
});

const pwdRepeat = ref("");
const isPwdSame = ref(false);
const pwdValid = ref(false);

const router = useRouter();
document.title = "注册";
async function registerClick() {
  if (!(isPwdSame.value && pwdValid.value)) {
    alert("请重新检查您的密码！");
  } else {
    var { data } = await RegisterApi(UserData);
    if (data.message == "ok") {
      alert("注册成功，跳转到登录页面");
      router.push("/login");
    }
  }
}
function CheckPwd() {
  isPwdSame.value = UserData.pwd == pwdRepeat.value;
  pwdValid.value = isPwdValid(UserData.pwd);
}
// function CheckInfo() {}

function cancelClick() {
  router.push("/index");
}

function loginClicked() {
  router.push("/login");
}

async function vertifyCodeClick() {
  verifyCode.email = UserData.email;
  console.log(verifyCode);
  var { data } = await VerifySendApi(verifyCode);
  if (data.message == "ok") {
    alert("已经请求发送邮件");
  }
}
</script>

<template>
  <!-- {{ UserData }}
  {{ verifyCode }} -->
  <div style="width: 70%; margin: 0 auto; padding-top: 20em; padding-bottom: 20em">
    <div class="ui stackable centered column grid">
      <div class="centered row">
        <div class="eight wide column">
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
                <label for="vertify-code">验证码：</label>
                <input
                  v-model="UserData.code"
                  type="text"
                  name="vertify-code"
                  id="vertify-code"
                  placeholder="请输入验证码"
                />
                <button
                  style="float: right"
                  class="ui inverted blue button"
                  @click.native="vertifyCodeClick"
                >
                  发送邮箱验证码
                </button>
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
              @click.native="loginClicked"
            >
              已有帐号
            </button>
            <button
              style="float: left"
              class="ui inverted red button"
              @click.native="cancelClick"
            >
              返回
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
