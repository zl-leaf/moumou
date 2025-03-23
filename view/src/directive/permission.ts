import type { App } from 'vue';
import { useUserStore } from '@/pinia/modules/user'

export default {
    install(app: App<any>) {
        app.directive('permission', {
            mounted(el: any, binding: any) {
                const userStore = useUserStore()
                if (!userStore.HasPermission(binding.value)) {
                    el.parentNode && el.parentNode.removeChild(el)
                }
            }
        })
    }
}