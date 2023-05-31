<!--登录-->
<script setup lang="ts">
import {onMounted} from "vue"
import {useRouter} from "vue-router";
import {login} from "@/api";
import {useMessage} from 'naive-ui'

const message = useMessage()
const router = useRouter();

onMounted(() => {
  let btn: any = document.querySelector(".login-content .login-content-submit .btn")
  btn.addEventListener("click", () => {
    let user: any = document.querySelector("#inputAccountExampleUser")
    let pass: any = document.querySelector("#inputPasswordExamplePass")
    login({username: user.value, password: pass.value}).then(res => {
      if (res.data.code === 1000) {
        message.success('登录成功')
        localStorage.setItem('token', res.data.data)
        router.push("/")
      } else {
        message.warning('登录失败')
      }
    })
  })
})
</script>

<template>
  <div class="login">
    <div class="container">
      <div class="login-content">
        <div class="login-content-logo ">
          <img src="https://blog.dbsgw.com/logo.png" alt="前端笔记">
        </div>
        <div class="login-content-user">
          <div class="input-control has-icon-left">
            <input id="inputAccountExampleUser" type="text" class="form-control" placeholder="用户名"
                   value="">
            <label for="inputAccountExampleUser" class="input-control-icon-left"><i
                class="icon icon-user "></i></label>
          </div>
        </div>
        <div class="login-content-pass">
          <div class="input-control has-icon-left">
            <input id="inputPasswordExamplePass" type="password" class="form-control" placeholder="密码"
                   value="">
            <label for="inputPasswordExamplePass" class="input-control-icon-left"><i
                class="icon icon-key"></i></label>
          </div>
        </div>
        <div class="login-content-submit">
          <button class="btn btn-block btn-primary" type="button">登录</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background-color: #f7f8fa;
}

.container {
  width: 100%;
  background-color: var(--bag)
}

.login-content {
  background-color: #fff;
  padding: 20px;
  border-radius: 5px;
  max-width: 390px;
  width: 100%;
  margin: 0 auto;
  box-shadow: 0 10px 30px -5px rgb(0 0 0 / 30%);
}

.login-content > div {
  padding: 8px;
  width: 100%;
}

.login-content-logo {
  text-align: center;
}

.login-content-logo img {
  width: 150px;
}
</style>