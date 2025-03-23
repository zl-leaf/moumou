import { useUserStore } from '@/pinia/modules/user'

export default {
    install(app) {
        app.directive('permission', {
            mounted(el, binding) {
                const userStore = useUserStore()
                if (!userStore.HasPermission(binding.value)) {
                    el.parentNode && el.parentNode.removeChild(el)
                }
            }
        })
    }
}