<template>
  <div class="container">
    <div class="left-column">
      <a-layout class="layout-column">
          <a-list :bordered="false" :pagination-props="paginationProps" :data="data.items">
            <template #item="{ item }">
              <a-list-item>
                <div class="post-image-area">
                  <img src="http://127.0.0.1:5173/assets/media/photos/photo1.jpg" />
                  <a-card :style="{ width: '100%'}" :title="item.title" :bordered="false" hoverable>
                    <a-list-item-meta
                        :title="`${item.author} on ${formatTimestamp(item.updated_at)}`"
                        :description="item.content"
                    >
                    </a-list-item-meta>
                  </a-card>
                </div>
              </a-list-item>
            </template>
          </a-list>
      </a-layout>
    </div>

    <div class="right-column">
      <a-layout class="layout-column">
        <a-card hoverable :style="{ width: '360px' }">
          <a-input-search v-model="request.keywords" :style="{width:'320px'}" placeholder="Please enter something" button-text="Search" search-button @search="ListBlog" @press-enter="ListBlog"/>
        </a-card>
      </a-layout>

      <a-layout class="layout-column">
        <a-card hoverable :style="{ width: '360px' }">
          <template #cover>
            <div :style="{ height: '180px', overflow: 'hidden', }">
              <img :style="{ width: '100%', transform: 'translateY(-20px)' }" src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a20012a2d4d5b9db43dfc6a01fe508c0.png~tplv-uwbnlip3yd-webp.webp"/>
              <div class="avatar-overlay">
                <img src="http://127.0.0.1:5173/assets/media/avatars/avatar1.jpg" alt="Avatar" :style="{ position: 'absolute', top: '95px', right: '120px', width: '120px', height: '120px', border: '2px solid #fff', 'border-radius':'50%' }">
              </div>
            </div>
          </template>
          <a-card-meta title="Don't give up on your dreams.😎" style="margin-top:30px; margin-bottom: 10px" align="center">
            <template #description>
              <icon-email></icon-email> yy@yybloger.com
            </template>
          </a-card-meta>
        </a-card>
      </a-layout>
    </div>
  </div>
</template>

<script setup>

import { reactive } from 'vue'
import {LIST_BLOG} from '@/common/api/blog.js'
import {onMounted, ref} from "vue";

const isLoading = ref(false)
const data = ref({items: [], total: 0});
const ListBlog = async () => {
  isLoading.value = true
  try {
    const resp = await LIST_BLOG(request.value)
    data.value = resp
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  ListBlog()
})

const request = ref({
  page_size: 10,
  page_number: 1,
  keywords: '',
})

const paginationProps = reactive({
  defaultPageSize: 5,
  total: request.value.page_size
})

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}

</script>

<style scoped>

.container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-wrap: wrap;
}

.left-column,
.right-column {
  flex: 1;
  padding: 20px;
}

.post-image-area {
  display: flex;
  height: 150px;
  border-radius: 2px;
}

/* Adjust padding to add spacing between columns */
.left-column {
  padding-top: 20px;
}

.right-column {
  padding-top: 30px;
}

/* Responsive styles */
@media (max-width: 768px) {
  .left-column,
  .right-column {
    flex: none;
    width: 100%;
  }
}

.layout-column {
  margin-bottom: 20px;
}


</style>
