<template>
  <div class="app">
    <!-- Main Section -->
    <main class="main-content">
      <section class="post-preview">

        <a-spin v-if="loadding" />

        <div v-if="!loadding && content" :style="{ background: 'var(--color-fill-2)', padding: '28px' }" >
          <a-page-header :style="{ background: 'var(--color-bg-2)' }" :title="title" @back="router.go(-1)">
            <template #breadcrumb>
              <a-breadcrumb>
                <a-breadcrumb-item></a-breadcrumb-item>
              </a-breadcrumb>
            </template>
            <template #subtitle>
              <a-space>
                <span>{{ `${author} on ${formatTimestamp(created_at)}` }}</span>
                <a-tag color="red" size="small">Admin</a-tag>
              </a-space>
            </template>
<!--            <template #extra>-->
<!--              <a-space>-->
<!--              </a-space>-->
<!--            </template>-->
            <MdPreview class="parent" :editorId="id" :modelValue="content"></MdPreview>
          </a-page-header>
        </div>
      </section>

    </main>
  </div>
</template>

<script setup>
import {GET_BLOG} from '@/common/api/blog.js'
import {onMounted, ref} from "vue";
import { MdPreview } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import {useRouter} from "vue-router";

const id = 'preview-only';

const router = useRouter()
const blogId=router.currentRoute.value.params.id

const query = ref({})
const loadding = ref(true)

const title = ref('')
const content = ref('')
const author = ref('')
const created_at = ref('')

const GetBlog = async () => {
  loadding.value = true
  try {
    const resp = await GET_BLOG(blogId, query.value)
    title.value = resp.title
    content.value = resp.content
    author.value = resp.author
    created_at.value = resp.created_at

  } finally {
    loadding.value = false
  }
}

onMounted(() => {
  GetBlog()
})

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
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
  background-color: white;
}

.post-preview {
  max-width: 1200px;
  margin: auto;
}
</style>
