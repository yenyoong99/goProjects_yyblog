<template>
  <a-affix :offsetTop="0">
    <div class="top-bar">
      <div>
        <a-image class="clickable" width="150" src="/public/logo-blue.png" :preview="false" @click="BlogView"/>
<!--        <div>YY部落站 {{ currentPath.startsWith('/dashboard') ? "Dashboard" : "" }}</div>-->
      </div>

      <div class="top-bar-left">
        <!-- 导航链接 -->
        <a-layout-header v-if="!currentPath.startsWith('/dashboard')">
          <a-menu mode="horizontal" theme="light" default-selected-keys="['home']" :collapsed="isSmallScreen" :collapsed-width="50">
            <a-menu-item key="home" @click="BlogView">首页</a-menu-item>
            <a-menu-item key="tech" disabled>技术文章</a-menu-item>
            <a-menu-item key="about" disabled>个人简介</a-menu-item>
            <a-menu-item key="contact" disabled>联系我</a-menu-item>
            <a-menu-item key="freehost" disabled>免费主机</a-menu-item>
            <a-menu-item key="blogapply" disabled>博客申请</a-menu-item>
          </a-menu>
        </a-layout-header>

        <!-- 只在非小屏幕设备上显示登录按钮 -->
        <a-button type="text" class="full-btn" v-if="isLogin && !currentPath.startsWith('/dashboard') && !isSmallScreen"
                  @click="Dashboard">Dashboard
        </a-button>
        <a-button type="text" class="full-btn" v-if="isLogin && currentPath.startsWith('/dashboard')" @click="BlogView">
          Visit Site
        </a-button>
        <a-button type="text" class="full-btn" v-if="isLogin" @click="Logout">Logout</a-button>
        <a-button type="text" class="full-btn" v-else @click="Login">Login</a-button>
      </div>
    </div>
  </a-affix>

  <!-- Header -->
  <a-layout-header v-if="!currentPath.startsWith('/dashboard')">
    <header class="header">
      <div class="header-banner">
        <h1 class="banner-title">YY部落站</h1>
        <p class="banner-subtitle">欢迎来到YY部落站，这里分享最新的技术文章和个人经验。</p>
      </div>
    </header>
  </a-layout-header>

</template>

<script setup>
import {useRouter} from 'vue-router'

const router = useRouter()

import {state} from '@/stores/index'
import {computed} from 'vue'

const isLogin = computed(() => {
  return state.value.token !== null
})

const currentPath = window.location.pathname;

// 判断是否为小屏幕设备
const isSmallScreen = window.innerWidth <= 768;

const BlogView = () => {
  router.push({name: 'BlogView'})
}

const Dashboard = () => {
  router.push({name: 'Dashboard'})
}

const Logout = () => {
  state.value.token = null
  if (currentPath.startsWith('/dashboard')) {
    router.push({name: 'LoginView'})
  } else {
    router.push({name: 'BlogView'})
  }
}

const Login = () => {
  router.push({name: 'LoginView'})
}

</script>

<style lang="css" scoped>

.top-bar {
  position: relative;
  padding: 0 12px;
  height: 60px;
  border-bottom: 1px solid rgb(229, 230, 235);
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: white;
}

.top-bar-left {
  height: 100%;
  display: flex;
  align-items: center;
}

.full-btn {
  height: 100%;
}

/* Header Styles */
.header {
  flex-shrink: 0;
}

.header-banner {
  text-align: center;
  padding: 3rem;
  background-color: #4f5562;
  /*background-image: url('public/banner-bg.jpg');*/
  background-size: cover;
  background-position: center;
}

.banner-title {
  font-size: 2rem;
  font-weight: bold;
  color: #ffffff;
  margin-bottom: 1rem;
}

.banner-subtitle {
  color: #d3d4c5;
  margin-bottom: 20px;
}
</style>
