<template>
    <ContentPage>
        <template #content>
            <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
                <a-form-item label="授权用户">
                    <a-select
                        v-model:value="formState.values"
                        mode="multiple"
                        placeholder="请输入用户名"
                        style="width: 100%"
                        :filter-option="false"
                        :not-found-content="formState.fetching ? undefined : null"
                        :options="formState.data"
                        @search="fetchUser"
                    >
                        <template v-if="formState.fetching" #notFoundContent>
                        <a-spin size="small" />
                        </template>
                    </a-select>
                </a-form-item>
                <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
                    <a-button type="primary" @click="onSubmit" :loading="loading" v-permission="'ManageRoleWrite'">提交</a-button>
                    <a-button style="margin-left: 10px" @click="$router.back()">返回</a-button>
                </a-form-item>
            </a-form>

            <div style="margin-top: 24px">
                <a-table
                    :columns="columns"
                    :data-source="tableData"
                    :loading="loading"
                    :pagination="false"
                    row-key="id"
                >
                    <template #bodyCell="{ column, record }">
                        <template v-if="column.key === 'action'">
                        </template>
                    </template>
                </a-table>
            </div>
        </template>
    </ContentPage>
</template>

<script setup lang="ts">
import ContentPage from '@/components/ContentPage.vue';
</script>
<script lang="ts">
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { useUserStore } from '@/pinia/modules/user';
import { debounce } from 'lodash';
import { message } from 'ant-design-vue';
import type { DefaultOptionType } from 'ant-design-vue/es/select';

type FromState = {
    data: DefaultOptionType[],
    values: string[],
    fetching: boolean,
}

export default defineComponent({
    data() {
        return {
            dataId: "",
            formState: ref<FromState>({
                data: [],
                values: [],
                fetching: false,
            }),
            loading: ref<boolean>(false),
            tableData: ref<any[]>([]),
            columns: [
                {
                    title: '用户名',
                    dataIndex: 'username',
                    key: 'username',
                },
                {
                    title: '操作',
                    key: 'action',
                    width: 120,
                },
            ],
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function(dataId: string) {
            this.loading = true
            let _this = this
            try {
                let getBindUserResponse = await api.RoleHandlerService.roleHandlerGetBindUser({
                    roleId: this.dataId
                })
                if (getBindUserResponse.code !== 0) {
                    throw new Error(getBindUserResponse.message)
                }
                const userList = getBindUserResponse.data?.list || []
                _this.tableData = userList
                userList.forEach(user => {
                    _this.formState.data.push({
                        label: user.username,
                        value: user.id,
                    })
                })

                this.loading = false
            } catch (err) {
                message.error("网络错误")
            }
        },
        onSubmit: async function() {
            this.loading = true
            try {
                let response = await api.RoleHandlerService.roleHandlerUpdateBindUser({
                    roleId: this.dataId,
                    userIds: this.formState.values
                })
                if(response.code != 0) {
                    throw new Error(response.message)
                }

                await useUserStore().UpdatePermissions()
                await this.initData(this.dataId)

                message.success('操作完成')
                this.loading = false
            } catch(err) {
                message.error("网络异常")
            }
        },
        fetchUser: debounce(async function(this: any, value: string) {
            this.formState.fetching = true
            try {
                let response = await api.UserHandlerService.userHandlerGetUserList({
                    filter: {
                        usernameLike: value,
                    },
                    currentPage: 1,
                    pageSize: 100,
                })
                if (response.code != 0) {
                    throw new Error(response.message)
                }
                
                this.formState.data = response.data?.list?.map(user => ({
                    label: user.username ?? '',
                    value: user.id ?? '',
                })) ?? []

                this.formState.fetching = false
            } catch(err) {
                message.error('网络错误')
            }
        }, 300)
    }
})
</script>