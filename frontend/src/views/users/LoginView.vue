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
        <a-form-item hide-label field="is_member">
          <a-checkbox v-model="form.is_member"> Remember </a-checkbox>
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

const form = ref({
  username: '',
  password: '',
  is_member: false
})

const handleSubmit = (data) => {
  console.log(data)
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

.login-form {
  display: flex;
  flex-direction: column;
}

.login-form .ant-form-item {
  margin-bottom: 20px;
}

.login-form .ant-form-item-control-input-content {
  border-radius: 5px;
}

.login-form .ant-input-affix-wrapper {
  border-color: #ccc;
}

.login-form .ant-input-affix-wrapper-focused {
  border-color: #007bff;
  box-shadow: none;
}

.login-form .ant-checkbox-wrapper {
  color: #333;
}

.login-form .ant-btn-primary {
  background-color: #007bff;
  border-color: #007bff;
}

.login-form .ant-btn-primary:hover {
  background-color: #0056b3;
  border-color: #0056b3;
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
