<script setup lang="ts">
import { provinces, cities } from "../utilities/address";
import { computed, onMounted, reactive, watch } from "vue";
import { IUserInfo } from "../types/user";
import { useUserStore } from "../stores/user";
import { useRouter } from "vue-router";
import {UpdateUserInfoApi} from "../apis/user"
document.title = "您的个人信息";
const store = useUserStore();
const router = useRouter();
// var userInfo: IUserInfo = reactive({
//   email: "",
//   name: "",
//   type: 0,
//   gender: 0,
//   province: "",
//   city: "",
//   phone: "",
//   profile: "",
// });
var userInfo: IUserInfo = store.data;
const pros = computed(() => {
    return provinces.indexOf(userInfo.province);
});

onMounted(() => {
    if (!store.isLogin) {
        alert("您还未登录");
        // router.push("/login");
    } else {
        alert(store.jwt);
    }
});

async function submitClick() {
    var data = await UpdateUserInfoApi(userInfo, store.jwt);
    if(data.message == "ok"){
        alert("更新成功")
        router.push("/index")
    }
}

function cancelClick() {
    router.push("/cancel");
}
</script>

<template>
    {{ pros }}
    {{ userInfo }}
    <div style="width: 70%; margin: 0 auto; padding-top: 20em; padding-bottom: 20em">
        <div class="ui stackable centered column grid">
            <div class="centered row">
                <div class="six wide column">
                    <div
                        style="padding-bottom: 5em; background-color: rgba(245, 245, 245, 0.5)"
                        class="ui segment"
                    >
                        <form id="regist" class="ui form" onsubmit="return false;">
                            <h3 style="color: #1890ff; text-align: center">修改信息</h3>
                            <div class="field">
                                <label for="name">姓名：</label>
                                <input
                                    type="text"
                                    name="name"
                                    id="name"
                                    placeholder="请输入姓名"
                                    v-model="userInfo.name"
                                />
                            </div>
                            <div class="field">
                                <label for="gender">性别：</label>
                                <select name="gender" id="gender" v-model="userInfo.gender">
                                    <option value="0">男</option>
                                    <option value="1">女</option>
                                    <option value="2">其他</option>
                                </select>
                            </div>
                            <form class="ui form" onsubmit="return false;">
                                <p>
                                    省份:
                                    <select
                                        name="province"
                                        id="province"
                                        v-model="userInfo.province"
                                    >
                                        <option>--请选择省份--</option>
                                        <option v-for="pro in provinces" :value="pro">{{ pro }}</option>
                                    </select>
                                </p>
                                <p>
                                    城市:
                                    <select name="city" id="city" v-model="userInfo.city">
                                        <option v-for="city in cities[pros]" :value="city">{{ city }}</option>
                                    </select>
                                </p>
                            </form>
                        </form>
                        <form class="ui form" onsubmit="return false;">
                            <div class="field">
                                <label for="phone">电话号：</label>
                                <input
                                    type="number"
                                    name="phone"
                                    id="phone"
                                    v-model="userInfo.phone"
                                />
                            </div>
                            <div class="field">
                                <label for="email">Email：</label>
                                <input type="text" name="email" id="email" v-model="userInfo.email" />
                            </div>
                            <div class="field">
                                <label for="profile">简介：</label>
                                <input
                                    type="text"
                                    name="profile"
                                    id="profile"
                                    v-model="userInfo.profile"
                                />
                            </div>
                            <button style="float: right" class="ui inverted green button" @click.native="submitClick">提交修改</button>
                            <button style="float: left" class="ui inverted red button">返回</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style></style>
