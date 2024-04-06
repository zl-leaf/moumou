<template>
    <a-breadcrumb :routes="breadItems" style="margin: 16px 0">
        <template #itemRender="{ route, params, routes, paths }">
            <span v-if="routes.indexOf(route) === routes.length - 1">{{ route.breadcrumbName }}</span>
            <router-link v-else :to="route.path">{{ route.breadcrumbName }}</router-link>
        </template>
    </a-breadcrumb>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

interface Route {
  path: string;
  breadcrumbName: string;
}

export default defineComponent({
    data() {
        return {
            breadItems: ref<Route[]>([])
        }
    },
    created() {
        this.updateBreadItems()
    },
    methods: {
        updateBreadItems() {
            this.breadItems = []
            this.$route.matched.filter(item => item.meta?.title).forEach(element => {
                this.breadItems.push({
                    path: element.path,
                    breadcrumbName: String(element.meta?.title),
                })
            })
        }
    },
    watch: {
        $route(to, from) {
            this.updateBreadItems()
        }
    }
})

</script>