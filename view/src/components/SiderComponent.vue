<template>
    <a-layout-sider v-model:collapsed="collapsed" collapsible>
        <div class="logo" />
        <a-menu v-model:openKeys="state.openKeys" v-model:selectedKeys="state.selectedKeys" mode="inline" theme="dark"
            :inline-collapsed="state.collapsed" :items="items" @click="onClick"></a-menu>
    </a-layout-sider>
</template>

<script lang="ts">
import { ref, defineComponent, reactive, h } from 'vue';
import type { MenuProps } from 'ant-design-vue';
import type { ItemType } from 'ant-design-vue';
import router from '@/router';
import type { RouteRecord, RouteRecordRaw } from 'vue-router';

interface MenuData {
    key: string
    title: string
    label: string
    children: Array<MenuData> | undefined
}

export default defineComponent({
    setup(props) {
        const formatRouter2Menu = (item: RouteRecordRaw): ItemType[] => {
            let menus: ItemType[] = []
            if (item.children) {
                item.children.forEach((childItem: RouteRecordRaw) => {
                    formatRouter2Menu(childItem).forEach((element: any) => {
                        if (element) {
                            menus.push(element)
                        }
                    });
                })
            }
            if (item.meta?.isMenu) {
                // 菜单
                let menuData: ItemType = {
                    key: String(item.name),
                    title: String(item.meta?.title),
                    label: String(item.meta?.title),
                    children: menus.length > 0 ? menus : undefined,
                }
                menus = [menuData]
            }
            return menus;
        }

        const items: ItemType[] = reactive([])
        router.getRoutes().forEach((item: RouteRecordRaw) => {
            if (item.name != 'layout') {
                // 只需要处理动态获取的menu
                return
            }
            formatRouter2Menu(item).forEach((item: ItemType) => {
                items.push(item)
            })
        })

        const state = reactive({
            collapsed: false,
            selectedKeys: Array<string>(),
            openKeys: Array<string>(),
        });
        const collapsed = ref<boolean>(false);

        return {
            state,
            items,
            collapsed,
        }
    },
    created() {
        this.updateSide()
    },
    methods: {
        onClick(e: any) {
            if (!this.state.selectedKeys.includes(e.key)) {
                this.$emit('openMenu', e)
            }
        },
        updateSide() {
            let currentRoute = router.currentRoute
            if (currentRoute.value.name) {
                this.state.selectedKeys = [String(currentRoute.value.name)]
            } else {
                this.state.selectedKeys = []
            }

            currentRoute.value.matched.forEach(element => {
                if (element.meta?.isMenu && element.children.length > 0) {
                    this.state.openKeys = [String(element.name)]
                }
            })
        }
    },
    watch: {
        $route(to, from) {
            this.updateSide()
        }
    },
    emits: {
        openMenu: (e: any) => { }
    }
})
</script>