<template>
  <div class="app">
    <!-- Main Section -->
    <header class="post-header">
      <h1>{{ title }}</h1>
      <div class="meta-info">
        <span class="author">Author: {{ author }}</span>
        <span class="created-at">Published: {{ formatTimestamp(created_at) }}</span>
      </div>
    </header>
    <main class="main-content">
      <a-affix :offsetTop="80">
        <div class="child">
          <h2 class="catalog-title">文章目录</h2>
          <MdCatalog :editorId="id" :scrollElement="scrollElement" />
        </div>
      </a-affix>
      <div class="parent">
        <MdPreview :editorId="id" :modelValue="content"></MdPreview>
<!--        <div class="navigation-buttons">-->
<!--          <button @click="navigateToPrevious">Previous</button>-->
<!--          <button @click="navigateToNext">Next</button>-->
<!--        </div>-->
      </div>
    </main>
  </div>
</template>

<script setup>
import { GET_BLOG } from '@/common/api/blog.js'
import { onMounted, ref } from "vue";
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { useRouter } from "vue-router";

const id = 'preview-only';

const router = useRouter()
const blogId = router.currentRoute.value.params.id

const scrollElement = document.documentElement;

const query = ref({})
const isLoading = ref(true)

const title = ref('')
const content = ref('')
const author = ref('')
const created_at = ref('')

const isSmallScreen = window.innerWidth <= 768;

const GetBlog = async () => {
  isLoading.value = true
  try {
    const resp = await GET_BLOG(blogId, query.value)
    title.value = truncatedTitle(resp.title)
    content.value = resp.content
    author.value = resp.author
    created_at.value = resp.created_at
  } finally {
    setTimeout(() => {
      isLoading.value = false;
    }, 500);
  }
}

onMounted(() => {
  GetBlog()
})

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}

const truncatedTitle = (title) => {
  if (isSmallScreen) {
    const maxLength = 12; // 设置最大字符数
    if (title.length > maxLength) {
      return title.substring(0, maxLength) + '...';
    }
  }
  return title;
}

// const navigateToPrevious = () => {
//   console.log('Navigate to previous article');
// }
//
// const navigateToNext = () => {
//   console.log('Navigate to next article');
// }
</script>

<style lang="css" scoped>
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
  background-color: #f4f4f9;
}

/* Header Styles */
.post-header {
  background-color: #4f5562;
  opacity: 0.9;
  background-image: url('/login_cover.jpg');
  /*color: #434a53;*/
  color: #ffffff;
  padding: 60px;
  text-align: center;
  border-bottom: 1px solid #e0e0e0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.post-header h1 {
  font-size: 2em;
  margin-bottom: 15px;
}

.meta-info {
  font-size: 0.9em;
  display: flex;
  justify-content: center;
  gap: 10px;
  color: #434a53;
}

.meta-info span {
  background-color: #f8f9fa;
  padding: 5px 10px;
  border-radius: 5px;
}

/* Main Section Styles */
.main-content {
  display: flex;
  flex-direction: row;
  max-width: 100%;
  overflow: hidden;
  padding: 20px;
}

.child {
  background-color: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
  padding: 20px;
  margin-right: 20px;
  max-width: 300px;
  min-width: 200px;
  flex-shrink: 0;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

.catalog-title {
  font-size: 1.2em;
  font-weight: bold;
  margin-bottom: 10px;
  color: #333;
  text-align: center;
}

.parent {
  flex-grow: 1;
  padding: 20px;
  background-color: #ffffff;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

/* Navigation Buttons */
/*.navigation-buttons {*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  margin: 20px 0;*/
/*}*/

.navigation-buttons button {
  background-color: #434a53;
  color: #ffffff;
  border: none;
  padding: 10px 20px;
  margin: 0 10px;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s;
}

.navigation-buttons button:hover {
  background-color: #2d3138;
}

@media (max-width: 768px) {
  .main-content {
    flex-direction: column;
    padding: 0;
  }

  .child {
    position: fixed;
    top: 0;
    left: -300px;
    height: 100%;
    width: 250px;
    max-width: none;
    transition: left 0.3s ease;
    z-index: 1000;
    padding: 20px;
  }
}
</style>
