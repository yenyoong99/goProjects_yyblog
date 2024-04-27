import { createRouter, createWebHistory } from 'vue-router'
import DashboardLayout from '../views/dashboard/layout.vue'
import { state } from '@/stores/index'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'BlogLayout',
      component: () => import('../views/blog/layout.vue'),
      children: [
          {
            path: '',
            name: 'BlogView',
            component: () => import('../views/blog/HomeView.vue')
          },
          {
            path: 'post/:id',
            name: 'PostView',
            component: () => import('../views/blog/PostView.vue')
          }
      ]
    },
    {
      path: '/login',
      name: 'LoginView',
      component: () => import('../views/users/LoginView.vue')
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: DashboardLayout,
      children: [
        {
          path: 'posts',
          name: 'PostsList',
          component: () => import('../views/dashboard/posts/ListView.vue')
        },
        {
          path: 'posts/detail/:id',
          name: 'PostsDetails',
          component: () => import('../views/dashboard/posts/DetailView.vue')
        },
        {
          path: 'posts/edit',
          name: 'PostsEdit',
          component: () => import('../views/dashboard/posts/EditView.vue')
        },
        // {
        //   path: 'comments',
        //   name: 'CommentsList',
        //   component: () => import('../views/dashboard/comments/ListView.vue')
        // }
      ]
    },

  ]
})

router.beforeEach((to) => {
  // access dashboard
  if (to.fullPath.startsWith('/dashboard')) {
    // check login state
    if (!state.value.token) {
      // redirect to login page
      return { name: 'LoginView' }
    }
  }
})

export default router
