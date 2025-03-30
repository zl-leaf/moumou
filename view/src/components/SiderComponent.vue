<template>
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
        <div class="logo" />
        <a-menu v-model:openKeys="state.openKeys" v-model:selectedKeys="state.selectedKeys" mode="inline" theme="dark"
            :inline-collapsed="state.collapsed" :items="items" @click="onClick"></a-menu>
    </a-layout-sider>
</template>

<script lang="ts">
import { ref, defineComponent, reactive, watch } from 'vue';
import type { ItemType } from 'ant-design-vue';
import router from '@/router';
import type { RouteRecordRaw } from 'vue-router';
import { useUserStore } from '@/pinia/modules/user';
import format from '@/utils/format'

const userStore = useUserStore()

export default defineComponent({
    props: {
        renderRouters: {
            type: Array as () => RouteRecordRaw[],
            required: true
        }
    },
    data() {
        return {
            state: reactive({
                collapsed: false,
                selectedKeys: Array<string>(),
                openKeys: Array<string>(),
            }),
            items: ref<ItemType[]>(),
            collapsed: ref<boolean>(false),
        }
    },
    created() {
        this.updateSide()
        this.updateSeletedMenu()
    },
    methods: {
        onClick(e: any) {
            if (!this.state.selectedKeys.includes(e.key)) {
                router.push({
                    path: e.key,
                    replace: true,
                })
            }
        },
        updateSeletedMenu() {
            let currentRoute = router.currentRoute
            this.state.selectedKeys = [String(currentRoute.value.path)]

            let openKeys:string[] = []
            currentRoute.value.matched.forEach(element => {
                if (element.meta?.isMenu && element.children.length > 0) {
                    openKeys.push(element.path)
                }
            })
            this.state.openKeys = openKeys
        },
        updateSide() {
            const formatRouter2Menu = (parentPath: string, item: RouteRecordRaw): ItemType[] => {
                let menus: ItemType[] = []
                if (item.meta?.permission) {
                    let needPermission = String(item.meta.permission)
                    if (!userStore.HasPermission(needPermission)) {
                        return menus
                    }
                }

                if (item.children?.length) {
                    item.children.forEach((childItem: RouteRecordRaw) => {
                        formatRouter2Menu(item.path, childItem).forEach((element: ItemType) => {
                            menus.push(element)
                        });
                    })
                }
                if (item.meta?.isMenu) {
                    let key = item.path
                    if (key.charAt(0) !== '/') {
                        key = parentPath + '/' + key
                    }

                    // 菜单
                    let menuData: ItemType = {
                        key: format.formatPath(parentPath, item.path),
                        title: String(item.meta?.title),
                        label: String(item.meta?.title),
                        children: menus.length > 0 ? menus : undefined,
                    }
                    menus = [menuData]
                }
                return menus;
            }

            let items:ItemType[] = []
            this.renderRouters.forEach((item: RouteRecordRaw) => {
                formatRouter2Menu('/', item).forEach((item: ItemType) => {
                    items.push(item)
                })
            })
            this.items = items

        }
    },
    watch: {
        $route(to, from) {
            this.updateSeletedMenu()
        },
    },
})
</script>