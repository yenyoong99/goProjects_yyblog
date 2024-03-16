<template>
  <div>
    <a-page-header
        title="Post Preview"
        @back="router.go(-1)"
    ></a-page-header>
    <a-spin v-if="loadding" />
    <div v-if="!loadding && content">
      <MdPreview class="parent" :editorId="id" :modelValue="content"></MdPreview>
      <MdCatalog class="child" :editorId="id" :scrollElement="scrollElement" />
    </div>


  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { GET_BLOG } from '@/common/api/blog.js'
import { onMounted, ref } from 'vue';
import { MdPreview, MdCatalog } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const id = 'preview-only';
const scrollElement = document.documentElement;

const router = useRouter()
const blogId=router.currentRoute.value.params.id

const query = ref({})
const loadding = ref(true)
const content = ref('')
const GetBlog = async () => {
  loadding.value = true
  try {
    const resp = await GET_BLOG(blogId, query.value)
    content.value = resp.content
  } finally {
    loadding.value = false
  }
}

onMounted(() => {
  GetBlog()
})

</script>

<style lang="css" scoped>
.parent {
  position: relative;
}

.child {
  position: absolute;
  top: 80px;
  right: 0;
  background-color: white;
}
</style>

