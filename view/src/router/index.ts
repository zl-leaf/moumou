import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/layout',
      name: 'layout',
      component: () => import('@/views/layout/AdminLayout.vue'),
      children: [
        {
          path: '/',
          name: 'dashboard',
          component:() => import('@/views/AboutView.vue'),
          meta: {
            title: '仪表盘',
            isMenu: true,
          }
        },
        {
          path: '',
          meta: {
            title: '系统管理',
            isMenu: true,
          },
          children: [
            {
              path: '/manage/user',
              children: [
                {
                  path: 'list',
                  name: 'user_list',
                  component: () => import('@/views/sys/user/list.vue'),
                  meta: {
                    title: '账号管理',
                    isMenu: true,
                  }
                },
                {
                  path: 'info',
                  name: 'user_info',
                  component: () => import('@/views/sys/user/info.vue'),
                }
              ]
            },
            {
              path: '/manage/role',
              children: [
                {
                  path: 'list',
                  name: 'role_list',
                  component: () => import('@/views/manage/role/List.vue'),
                  meta: {
                    title: '角色管理',
                    isMenu: true,
                  }
                },
                {
                  path: 'add',
                  name: 'role_add',
                  component: () => import('@/views/manage/role/Add.vue'),
                },
                {
                  path: 'info',
                  name: 'role_info',
                  component: () => import('@/views/manage/role/Info.vue'),
                },
                {
                  path: 'permission',
                  name: 'role_permission',
                  component: () => import('@/views/manage/role/Permission.vue'),
                }
              ]
            }
          ]
        },
        
      ]
    }
  ]
})

export default router
