<template>
  <div class="top-bar">
    <div>
      <div>YYBlog {{ currentPath.startsWith('/dashboard') ? "Dashboard" : ""}}</div>
    </div>
    <div class="top-bar-left">
      <a-button type="text" class="full-btn" v-if="isLogin && !currentPath.startsWith('/dashboard')" @click="Dashboard">Dashboard</a-button>
      <a-button type="text" class="full-btn" v-if="isLogin" @click="Logout">Logout</a-button>
      <a-button type="text" class="full-btn" v-else @click="Login">Login</a-button>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()

import { state } from '@/stores/index'
import { computed } from 'vue'
const isLogin = computed(() => {
  return state.value.token !== null
})
// console.log(state.value)

const currentPath = window.location.pathname;

const Dashboard = () => {
  router.push({ name: 'Dashboard' })
}

const Logout = () => {
  state.value.token = null
  router.push({ name: 'LoginView' })
}

const Login = () => {
  router.push({ name: 'LoginView' })
}
</script>

<style lang="css" scoped>
.top-bar {
  padding: 0 12px;
  height: 60px;
  border-bottom: 1px solid rgb(229, 230, 235);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.top-bar-left {
  height: 100%;
  display: flex;
  align-items: center;
}

.full-btn {
  height: 100%;
}
</style>
