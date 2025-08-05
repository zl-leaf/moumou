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
      name: 'admin_layout',
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
          path: '/manage',
          meta: {
            title: '系统管理',
            isMenu: true,
          },
          children: [
            {
              path: '/manage/user',
              meta: {
                title: '账号管理',
                isMenu: true,
                page: 'user_list',
                permission: 'ManageUserRead',
              },
              children: [
                {
                  path: '',
                  name: 'user_list',
                  component: () => import('@/views/manage/user/List.vue'),
                },
                {
                  path: 'add',
                  name: 'user_add',
                  component: () => import('@/views/manage/user/Add.vue'),
                  meta: {
                    title: '添加',
                  },
                },
                {
                  path: 'info',
                  name: 'user_info',
                  component: () => import('@/views/manage/user/Info.vue'),
                  meta: {
                    title: '详情',
                  },
                }
              ]
            },
            {
              path: '/manage/role',
              meta: {
                title: '角色管理',
                isMenu: true,
                page: "role_list",
                permission: 'ManageRoleRead',
              },
              children: [
                {
                  path: '',
                  name: 'role_list',
                  component: () => import('@/views/manage/role/List.vue'),
                },
                {
                  path: 'add',
                  name: 'role_add',
                  component: () => import('@/views/manage/role/Add.vue'),
                  meta: {
                    title: '添加',
                  },
                },
                {
                  path: 'info',
                  name: 'role_info',
                  component: () => import('@/views/manage/role/Info.vue'),
                  meta: {
                    title: '详情',
                  },
                },
                {
                  path: 'permission',
                  name: 'role_permission',
                  component: () => import('@/views/manage/role/Permission.vue'),
                  meta: {
                    title: '角色权限',
                  },
                },
                {
                  path: 'bind_user',
                  name: 'role_bind_user',
                  component: () => import('@/views/manage/role/BindUser.vue'),
                  meta: {
                    title: '绑定用户',
                  },
                }
              ]
            },
            {
              path: '/manage/permission',
              meta: {
                title: '权限管理',
                isMenu: true,
                page: "permission_list",
                permission: 'ManagePermissionRead',
              },
              children: [
                {
                  path: '',
                  name: 'permission_list',
                  component: () => import('@/views/manage/permission/List.vue'),
                },
                {
                  path: 'add',
                  name: 'permission_add',
                  component: () => import('@/views/manage/permission/Add.vue'),
                  meta: {
                    title: '添加',
                  },
                },
                {
                  path: 'info',
                  name: 'permission_info',
                  component: () => import('@/views/manage/permission/Info.vue'),
                  meta: {
                    title: '详情',
                  },
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
