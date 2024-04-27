<template>
  <div class="app">
    <!-- Main Section -->
    <main class="main-content">
      <section class="post-preview">

        <a-result status="404" subtitle="Whoops, that page is gone.">
          <template #extra>
            <a-space>
              <a-link href="/">back</a-link>
            </a-space>
          </template>
        </a-result>

      </section>

    </main>
  </div>
</template>


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

.post-preview {
  max-width: 1200px;
  margin: auto;
}

.go-back-btn {
  padding: 0.5rem 1.5rem;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}
</style>
