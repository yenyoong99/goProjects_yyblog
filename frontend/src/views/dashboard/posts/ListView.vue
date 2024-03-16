<template>
  <div>
    <div>
      <a-breadcrumb>
        <a-breadcrumb-item>Posts</a-breadcrumb-item>
        <a-breadcrumb-item>All Posts</a-breadcrumb-item>
      </a-breadcrumb>
    </div>
    <div class="op">
      <div>
        <a-button type="primary" @click="router.push({name: 'PostsEdit'})">Create Post</a-button>
      </div>
      <div>
        <a-input
            v-model="request.keywords"
            :style="{ width: '320px' }"
            placeholder="search"
            allow-clear
            @change="ListBlog"
        />
      </div>
    </div>
    <div>
      <a-table :loading="isLoading" column-resizable :bordered="{cell:true}" :pagination="false" :data="data.items">
        <template #columns>
<!--          <a-table-column title="Id" data-index="id" align="center"></a-table-column>-->
          <a-table-column title="Title" data-index="title" align="center"></a-table-column>
          <a-table-column title="Author" data-index="author" align="center"></a-table-column>
          <a-table-column title="Category" align="center">
            <template #cell="{ record }">
              <a-tag key="Menu" color="pinkpurple" bordered>{{ record.tags['Menu'] }}</a-tag>
            </template>
          </a-table-column>
          <a-table-column title="Last Modified" data-index="updated_at" align="center">
            <template #cell="{ record }">
              {{ formatTimestamp(record.updated_at) }}
            </template>
          </a-table-column>
          <a-table-column title="Action" align="center">
            <template #cell="{ record }">
              <a-space>
                <a-button @click="router.push({name: 'PostsDetails', params: {id: record.id}})">Preview</a-button>
                <a-button @click="router.push({name: 'PostsEdit', query: {id: record.id}})">Edit</a-button>
                <a-button @click="$modal.info({ title:'Name', content:record.title })">Delete</a-button>
              </a-space>
            </template>
          </a-table-column>
        </template>
      </a-table>
    </div>
    <div class="pagi">
      <a-pagination
          :total="data.total"
          :page-size-options="[2, 10, 20, 50, 100, 200]"
          @page-size-change="handlePageSizeChange"
          @change="hanlePageNumberChange"
          show-total
          show-jumper
          show-page-size
      />
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'
import {LIST_BLOG} from '@/common/api/blog.js'
import {useRouter} from 'vue-router';

const router = useRouter()

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
  create_by: '',
  keywords: '',
})

const handlePageSizeChange = (pageSize) => {
  request.value.page_size = pageSize
  ListBlog()
}

const hanlePageNumberChange = (pageNumber) => {
  request.value.page_number = pageNumber
  ListBlog()
}

const formatTimestamp = (timestamp) => {
  const date = new Date(timestamp * 1000);
  return date.toLocaleString();
}

</script>

<style lang="css" scoped>

.op {
  margin-top: 8px;
  margin-bottom: 8px;
  display: flex;
  justify-content: space-between;
}

.pagi {
  margin-top: 4px;
  display: flex;
  flex-direction: row-reverse;
}

</style>
