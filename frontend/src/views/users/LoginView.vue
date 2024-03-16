<template>
  <div class="login-container">
    <div class="login-content">
      <a-form :rules="formRules" :model="form" @submit="handleSubmit">
        <div class="title">YYBlog Login</div>
        <a-form-item hide-label field="username">
          <a-input v-model="form.username" placeholder="Username">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item hide-label field="password">
          <a-input-password v-model="form.password" placeholder="Password">
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item hide-label field="remember">
          <a-checkbox v-model="form.remember"> Remember </a-checkbox>
        </a-form-item>
        <a-form-item hide-label>
          <a-button type="primary" html-type="submit" style="width: 100%">Login</a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="footer">© 2024 YYBlog. All rights reserved.</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useRoute, useRouter } from 'vue-router'
import { state } from '@/stores/index'
import { LOGIN } from '@/common/api/login'

const form = ref({
  username: state.value.username,
  password: '',
  remember: false
})

const router = useRouter()
const route = useRoute()

const handleSubmit = async (data) => {
  if (!data.errors) {
    console.log(data)
    const form = data.values

    try {
      let resp = await LOGIN(form)
      state.value.isLogin = true
      state.value.username = resp.username
    } catch (error) {
      Message.error(`Login Fail: ${error}`)
    }

    let to = 'Dashboard'
    if (route.query.to) {
      to = route.query.to
    }
    await router.push({name: to})
  }
}

const formRules = {
  username: [
    {
      required: true,
      message: 'Please enter your username'
    }
  ],
  password: [
    {
      required: true,
      message: 'Please enter your password'
    },
    {
      minLength: 6,
      message: 'Password must be at least 6 characters'
    }
  ]
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-image: url('public/login_cover.jpg');
  background-size: cover;
  position: relative;
}

.login-content {
  width: 400px;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 10px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
}

.title {
  text-align: center;
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 20px;
  color: #333;
}

.footer {
  position: fixed;
  bottom: 10px;
  width: 100%;
  text-align: center;
  font-size: 12px;
  color: #666;
}
</style>
