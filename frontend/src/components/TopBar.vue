<template>
  <div class="top-bar">
    <div>
      <div>YYBlog Dashboard</div>
    </div>
    <div class="top-bar-left">
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
  return state.value.isLogin !== false
})
console.log(state.value)

const Logout = () => {
  state.value.isLogin = false
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
