<template>
  <a-config-provider :locale="enUS">
    <TopBar></TopBar>
    <div class="page">
      <router-view></router-view>
    </div>
    <div class="footer">© 2024 YYBlog. All rights reserved.</div>
  </a-config-provider>
</template>

<script setup>
import TopBar from '../../components/TopBar.vue'
import enUS from '@arco-design/web-vue/es/locale/lang/en-us'
</script>

<style lang="css" scoped>

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: Arial, sans-serif;
}

.page {
  background-color: #eceef2;
}

.footer {
  bottom: 10px;
  width: 100%;
  text-align: center;
  font-size: 12px;
  color: #666;
  border-top: 1px solid #e1e1e1;
  padding: 15px 0;
  background-color: #f6f7f9;
}
</style>
