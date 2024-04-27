<template>
  <div class="app">
    <!-- Main Section -->
    <main class="main-content">
      <section class="post-preview">

        <a-spin dot v-if="loadding"/>

        <div v-if="!loadding && content" :style="{ background: 'var(--color-fill-2)', padding: '28px' }">
          <a-page-header :style="{ background: 'var(--color-bg-2)' }" :title="title" @back="router.go(-1)">
            <template #breadcrumb>
              <a-breadcrumb>
                <a-breadcrumb-item></a-breadcrumb-item>
              </a-breadcrumb>
            </template>
            <template #extra>
              <a-space>
                <span>{{ `${author} on ${formatTimestamp(created_at)}` }}</span>
                <a-tag color="red" size="small">Admin</a-tag>
              </a-space>
            </template>
            <MdPreview class="parent" :editorId="id" :modelValue="content"></MdPreview>
          </a-page-header>

          <!-- comment box -->
          <a-page-header
              :style="{ background: 'var(--color-bg-2)' }"
              title="Comment Post"
              :show-back="false"
          >
            <template #extra>
              <a-space>
              <span>
                Total comment: 1
              </span>
              </a-space>
            </template>
            <a-card :style="{ width: '100%', marginBottom: '20px' }">
              <div :style="{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }">
              <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }">
                <a-comment
                    author="Admin"
                    content="Comment body content test."
                    datetime="1 hour"
                >
                <template #actions>
                  <span class="action" key="heart" @click="onLikeChange">
                    <span v-if="like">
                      <IconHeartFill :style="{ color: '#f53f3f' }"/>
                    </span>
                    <span v-else>
                      <IconHeart/>
                    </span>
                    {{ 83 + (like ? 1 : 0) }}
                  </span>
                  <span class="action" key="star" @click="onStarChange">
                    <span v-if="star">
                      <IconStarFill style="{ color: '#ffb400' }"/>
                    </span>
                    <span v-else>
                      <IconStar/>
                    </span>
                    {{ 3 + (star ? 1 : 0) }}
                  </span>
                  <span class="action" key="reply">
                    <IconMessage/> Reply
                  </span>
                </template>
                <template #avatar>
                  <a-avatar>
                    <img
                        alt="avatar"
                        src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"
                    />
                  </a-avatar>
                </template>
              </a-comment>
              </span>
                <!--              <a-link>More</a-link>-->
              </div>
            </a-card>
          </a-page-header>
        </div>

        <a-empty v-else>
          <template #image>
            <icon-empty/>
          </template>
          No data, please reload!
        </a-empty>

      </section>

    </main>
  </div>
</template>

<script setup>
import {GET_BLOG} from '@/common/api/blog.js'
import {onMounted, ref} from "vue";
import {MdPreview} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import {useRouter} from "vue-router";

const id = 'preview-only';

const router = useRouter()
const blogId = router.currentRoute.value.params.id

const query = ref({})
const loadding = ref(true)

const title = ref('')
const content = ref('')
const author = ref('')
const created_at = ref('')

const isSmallScreen = window.innerWidth <= 768;

const GetBlog = async () => {
  loadding.value = true
  try {
    const resp = await GET_BLOG(blogId, query.value)
    title.value = truncatedTitle(resp.title)
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

const truncatedTitle = (title) => {
  if (isSmallScreen) {
    const maxLength = 12; // 设置最大字符数
    if (title.length > maxLength) {
      return title.substring(0, maxLength) + '...';
    }
  }
  return title;
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

.action {
  display: inline-block;
  padding: 0 4px;
  color: var(--color-text-1);
  line-height: 24px;
  background: transparent;
  border-radius: 2px;
  cursor: pointer;
  transition: all 0.1s ease;
}

.action:hover {
  background: var(--color-fill-3);
}
</style>
