<script setup lang="ts">
// TODO ?
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
                        <form action="/login/" method="POST" class="ui form">
                            <h3 style="color:  #1890ff; text-align:center;">用户登录</h3>
                            <div class="field">
                                <label for="username">账号：</label>
                                <input
                                    type="text"
                                    name="username"
                                    id="username"
                                    placeholder="请输入手机号或邮箱"
                                />
                            </div>
                            <div class="field">
                                <label for="password">密码：</label>
                                <input
                                    type="password"
                                    name="password"
                                    id="password"
                                    placeholder="请输入密码"
                                />
                            </div>
                            <input
                                style="float: right;"
                                class="ui inverted green button"
                                type="submit"
                                value="登陆"
                            />
                            <input
                                style="float: right;"
                                class="ui inverted gray button"
                                type="submit"
                                value="取消"
                            />
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
</style>