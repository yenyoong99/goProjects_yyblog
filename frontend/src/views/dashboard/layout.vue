<template>
  <a-config-provider :locale="enUS">
    <TopBar></TopBar>
    <div class="main">
      <div class="side-bar">
        <a-menu
            @menu-item-click="handleMenuItemClick"
            :style="{ width: '200px', height: '100%' }"
            :default-open-keys="['blog']"
            :default-selected-keys="['blog_list']"
            show-collapse-button
            breakpoint="xl"
        >
          <a-sub-menu key="blog">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>Posts</template>
            <a-menu-item key="PostsList">All Posts</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="comment">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>Comments</template>
            <a-menu-item key="CommentsList">All Comments</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </div>
      <div class="page">
        <RouterView></RouterView>
      </div>
    </div>
  </a-config-provider>
</template>

<script setup>
import TopBar from '../../components/TopBar.vue'
import enUS from '@arco-design/web-vue/es/locale/lang/en-us'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleMenuItemClick = (key) => {
  router.push({ name: key })
}
</script>

<style lang="css" scoped>
.main {
  height: calc(100vh - 60px);
  display: flex;
}

.side-bar {
  height: 100%;
  border-right: 1px solid rgb(229, 230, 235);
}

.page {
  padding: 6px;
  overflow:auto;
  height: calc(100vh - 60px);
  width: 100%
}
</style>
