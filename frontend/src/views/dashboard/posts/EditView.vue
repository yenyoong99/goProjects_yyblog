<template>
  <div>
    <a-page-header
        :title="MODE"
        @back="router.go(-1)"
    >
    </a-page-header>
    <a-form ref="formRef" layout="vertical" :model="form">
      <a-form-item required field="title" label="Title">
        <a-input
            v-model="form.title"
            placeholder="Add title"
        />
      </a-form-item>
      <a-form-item required field="summary" label="Summary">
        <a-input v-model="form.summary" placeholder="Add summary" />
      </a-form-item>
      <a-form-item required field="content" label="Content">
        <MdEditor
            language="en-US"
            class="editor"
            v-model="form.content"
            @onSave="handleSave"
            @onUploadImg="handleImageUpload"
        >
        </MdEditor>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { GET_BLOG, CRATE_BLOG, UPDATE_BLOG, UPLOAD_BLOG_IMAGES } from '@/common/api/blog.js'
import { state } from '../../../stores/index'
import { Message } from '@arco-design/web-vue';

const router = useRouter()

const MODE = ref('Edit Post')
const id = router.currentRoute.value.query.id
if (!id) {
  MODE.value = 'New Post'
}


const formRef = ref(null)
const form = ref({
  title: '',
  summary: '',
  author: state.value.token.username,
  content: '',
})

const handleImageUpload = async (files, callback) => {
  try {
    const response = await UPLOAD_BLOG_IMAGES(files)
    const { urls } = response
    callback(urls)
  } catch (error) {
    Message.error('Image upload failed')
  }
}

const handleSave = async () => {
  const err = await formRef.value.validate()
  if (!err) {
    switch (MODE.value) {
      case 'Edit Post':
        await UPDATE_BLOG(router.currentRoute.value.query.id, form.value)
        MODE.value = 'Edit Post'
        Message.success('Post Saved')
        break;
      case 'New Post':
        var resp = await CRATE_BLOG(form.value)
        router.replace({query: {id: resp.id}})
        MODE.value = 'Edit Post'
        Message.success('Post Created')
        break;
    }
  }
}

onMounted( async () => {
  if (id) {
    const resp = await GET_BLOG(router.currentRoute.value.query.id)
    form.value = resp
  }
})

</script>

<style lang="css" scoped>
.editor {
  height: calc(100vh - 100px);
}
</style>
