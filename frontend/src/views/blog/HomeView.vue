<template>
  <div class="app">

    <!-- Main Section -->
    <!-- Header -->
    <a-layout-header>
      <header class="header">
        <div class="header-banner">
          <h1 class="banner-title">YY部落站</h1>
          <p class="banner-subtitle">欢迎来到YY部落站，这里分享最新的技术文章和个人经验。</p>
        </div>
      </header>
    </a-layout-header>

    <main class="main-content">
      <section class="articles-preview">

        <div style="text-align: center">
          <h2 class="section-title">最新文章</h2>
          <p class="section-subtitle">希望可以带给你一些帮助和启发</p>

          <a-card hoverable :style="{ width: '100%' }">
            <a-input-search v-model="request.keywords" :style="{width:'100%'}" placeholder="Please enter something" button-text="Search" search-button @search="searchTrigger" @press-enter="searchTrigger"/>
          </a-card>

          <a-layout-header style="margin-bottom: 15px">
            <a-menu mode="horizontal" theme="light" default-selected-keys="['home']">
              <a-menu-item key="home">前端</a-menu-item>
              <a-menu-item key="tech">后端</a-menu-item>
              <a-menu-item key="mobile">移动端</a-menu-item>
              <a-menu-item key="about">逆向</a-menu-item>
<!--              <a-menu-item key="ai">人工智能</a-menu-item>-->
              <a-menu-item key="security">网络与安全</a-menu-item>
              <a-menu-item key="os">操作系统</a-menu-item>
              <a-menu-item key="other">其他</a-menu-item>
            </a-menu>
          </a-layout-header>
        </div>

        <div class="articles-container">
          <div class="article-card" v-for="item in data.items" :key="item">

            <a-skeleton v-if="isLoading" :animation="true">
              <a-space direction="vertical" :style="{width:'100%'}" size="large">
                <a-skeleton-line :rows="3" />
                <a-skeleton-shape />
              </a-space>
            </a-skeleton>

            <a-card v-else class="clickable" hoverable @click="navigatePostTo(item.id)">
              <template #cover>
                <div :style="{ height: '250px', overflow: 'hidden' }">
                  <img :style="{ width: '100%', transform: 'translateY(-20px)' }" :src="`https://picsum.photos/id/${item.id}/300/200`"/>
                </div>
              </template>
              <a-card-meta :title="item.title">
              <template #description>
                {{ truncatedContent(item.summary) }} <br />
                <time class="article-date">{{ `${item.author} on ${formatTimestamp(item.created_at)}` }}</time>
              </template>
            </a-card-meta>
            </a-card>

          </div>
        </div>
      </section>

      <!-- Load More Button -->
      <div class="load-more-container">
        <button @click="loadMoreArticles" class="load-more-btn" v-if="!loadMoreDisabled">加载更多</button>
        <div v-else><a style="color: #7d7d7f">-- no more --</a></div>
      </div>
    </main>
  </div>
</template>

<script setup>
import {LIST_BLOG} from '@/common/api/blog.js'
import {onMounted, ref} from "vue";

const searchCheck = ref(false)
const isLoading = ref(false)
const loadMoreDisabled = ref(false)
const data = ref({items: [], total: 0});

const ListBlog = async () => {
  isLoading.value = true
  try {
    const resp = await LIST_BLOG(request.value)
    data.value = resp
  } finally {
    setTimeout(() => {
      isLoading.value = false;
    }, 500);
  }
  if (request.value.page_size >= data.value.total) {
    loadMoreDisabled.value = true
  }
}

const searchTrigger = () =>{
  ListBlog()
  request.value.page_size = 9
  searchCheck.value = !!request.value.keywords
}

onMounted(() => {
  ListBlog()
})

const request = ref({
  page_size: 3,
  page_number: 1,
  keywords: '',
})

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}

const loadMoreArticles = () => {
  request.value.page_size += 3
  ListBlog()
}

const navigatePostTo = (postId) => {
  window.location.href = '/post/' + postId.toString();
}

const truncatedContent = (content) => {
  const maxLength = 100; // 设置最大字符数
  if (content.length > maxLength) {
    return content.substring(0, maxLength) + '...';
  }
  return content;
}

</script>

<style scoped>
/* General Styles */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body, button {
  font-family: 'Helvetica Neue', Arial, sans-serif;
}

.app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

/* Main Section Styles */
.main-content {
  flex: 1;
  padding: 2rem;
}

.articles-preview {
  max-width: 1200px;
  margin: auto;
}

.section-title {
  font-size: 1.5rem;
  color: #333;
  margin-bottom: 1rem;
}

.section-subtitle {
  color: #666;
  margin-bottom: 2rem;
}

.articles-container {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.article-card {
  width: calc(33.700% - 1rem);
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.article-image img {
  width: 100%;
  height: auto;
  display: block;
}

.article-date {
  color: #999;
  margin-top: 1rem;
  font-size: 0.8rem;
}

.load-more-container {
  text-align: center;
  margin-top: 2rem;
}

.load-more-btn {
  padding: 0.5rem 1.5rem;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

.load-more-btn:hover {
  background-color: #0056b3;
}

/* Responsive Styles */
@media (max-width: 768px) {
  .nav-bar, .user-actions {
    display: none;
  }

  .articles-container {
    flex-direction: column;
  }

  .article-card {
    width: 100%;
  }
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
