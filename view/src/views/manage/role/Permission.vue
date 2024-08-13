<template>
    <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col="{ span: 8 }">
        <a-checkbox-group v-model:value="formState.values" name="checkboxgroup" :options="permissionOptions" />
        <a-form-item :wrapper-col="{ span: 8, offset: 6 }">
            <a-button type="primary" @click="onSubmit" :loading="loading">提交</a-button>
            <a-button style="margin-left: 10px" @click="$router.back()">返回</a-button>
        </a-form-item>
    </a-form>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import * as api from '@/api'
import router from '@/router';
import { message } from 'ant-design-vue';
import { useRouterStore } from '@/pinia/modules/router';
import { useUserStore } from '@/pinia/modules/user';

type FormState = {
    values: string[]
}

type PermissionOption = {
    label: string
    value: string
}

const routerStore = useRouterStore()
export default defineComponent({
    data() {
        return {
            dataId: "",
            permissionOptions: ref<PermissionOption[]>([]),
            formState: ref<FormState>({
                values: []
            }),
            loading: ref<boolean>(false),
        }
    },
    created() {
        this.dataId = String(this.$route.query.id)
        this.initData(this.dataId)
    },
    methods: {
        initData: async function(dataId: string) {
            let _this = this
            this.loading = true
            try {
                let permissionResponse = await api.PermissionHandlerService.permissionHandlerGetPermissionList({})
                if (permissionResponse.code != 0) {
                    throw new Error(permissionResponse.message)
                }
                let that = this
                permissionResponse.data?.list?.forEach(function(permission) {
                    that.permissionOptions.push({
                        label: permission.name ?? "",
                        value: permission.id ?? "",
                    })
                })

                // 加载详情数据
                let infoResponse = await api.RoleHandlerService.roleHandlerGetRoleInfo({
                    id: this.dataId,
                    field: {permission: true}
                })
                if (infoResponse.code != 0) {
                    throw new Error(infoResponse.message)
                }
                infoResponse.data?.permissions?.forEach(function(permission) {
                    _this.formState.values.push(permission.id ?? "")
                })

                this.loading = false
            } catch (err) {
                message.error("网络错误")
            }

        },
        onSubmit: async function() {
            this.loading = true
            try {
                let response = await api.RoleHandlerService.roleHandlerUpdateRolePermission({
                    id: this.dataId,
                    permission_ids: this.formState.values
                })
                if(response.code != 0) {
                    throw new Error(response.message)
                }

                await useUserStore().UpdatePermissions()

                message.success('操作完成')
                this.loading = false
            } catch(err) {
                message.error("网络异常")
            }
        }
    }
})
</script>