<template>
  <div class="login-page">
    <div class="login-left">
      <div class="login-content">
        <img src="/public/logo-blue.png" alt="Logo" class="logo"/>
        <div class="title">Login to YYBlog</div>
        <a-form :rules="formRules" :model="form" @submit="handleSubmit">
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
            <a-button type="primary" html-type="submit" style="width: 100%">Sign In</a-button>
          </a-form-item>
          <div class="forgot-password">
            <a href="#">Forgot Password?</a>
          </div>
        </a-form>
      </div>
    </div>
    <div class="login-right"></div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { state } from '@/stores/index';
import { LOGIN } from '@/common/api/login';

const form = ref({
  username: '',
  password: '',
  remember: false
});

const router = useRouter();

const handleSubmit = (data) => {
  if (!data.errors) {
    LOGIN(form.value).then((response) => {
      state.value.token = response;
      router.push({ name: 'Dashboard' });
    });
  }
};

onMounted(() => {
  if (state.value.token !== null) {
    router.push({ name: 'Dashboard' });
  }
});

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
};
</script>

<style scoped>
.login-page {
  display: flex;
  height: 100vh;
}

.login-left {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f7f9fc;
}

.login-content {
  width: 450px;
  padding: 60px;
  text-align: center;
}

.logo {
  width: 100px;
  margin-bottom: 20px;
}

.title {
  font-size: 18px;
  color: #333;
  margin-bottom: 20px;
}

.login-right {
  flex: 2;
  background: url('public/login_cover.jpg') no-repeat center center;
  background-size: cover;
}

.forgot-password {
  font-size: 12px;
}
</style>
