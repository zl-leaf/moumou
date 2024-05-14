<template>
    <a-table 
        :columns="columns"
        :data-source="data"
        :pagination="pagination"
        @change="handleTableChange"
    >
        <template #headerCell="{ column }"></template>

        <template #bodyCell="{ column, record }"></template>
    </a-table>
</template>
<script lang="ts">
import { defineComponent, computed, ref } from 'vue';
import {getPageAPI, getPageDataListAPI, type Page, type PageSchemaAttribute} from '../api/page';
import type { TableColumnType } from 'ant-design-vue';


export default defineComponent({
    data() {
        return {
            columns: Array<TableColumnType>(),
            data: Array<any>(),
            pagination: {
                total: 0,
                current: 1,
                pageSize: 10,
            },
        }
    },
    created() {
        this.updateView(1)
        this.handleTableChange(this.pagination, {}, {})
    },
    methods: {
        updateView: function(pageId: any) {
            this.fetchPageSchema(pageId).then((page:Page) => {
                page.schema.attributes.forEach((attribute:PageSchemaAttribute) => {
                    this.columns.push({
                        title: attribute.field_attribute.label,
                        key: attribute.field_attribute.key,
                        dataIndex: attribute.field_attribute.key,
                    })
                })
                
            }).catch(err => {
                console.log(err)
            })
        },
        fetchPageSchema: async function(pageId: any) {
            try {
                const getPageResponse = await getPageAPI(pageId)
                if (getPageResponse.code != '0') {
                    return Promise.reject(getPageResponse.message);
                }
                return getPageResponse.data.page
            } catch (err) {
                return Promise.reject("网络错误")
            }
        },
        handleTableChange: function(pag: any, filters: any, sorter: any) {
            this.pagination.current = pag.current
            this.pagination.pageSize = pag.pageSize
            console.log('change', pag, filters, sorter)

            getPageDataListAPI(1, pag.current, pag.pageSize).then((response) => {
                if (response.code != '0') {
                    return Promise.reject(response.message)
                }
                return response.data
            }).then((data) => {
                this.pagination.total = data.total
                this.data = data.list
            }).catch(err => {
                console.log('err:', err)
            })
        }
    }
})

</script>