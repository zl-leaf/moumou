import axios from "axios";
import { useUserStore } from "@/pinia/modules/user"

export const http = axios.create({
    baseURL: import.meta.env.VITE_BASE_API,
    timeout: 99999,
})


http.interceptors.request.use(
    config => {
        const userStore = useUserStore()
        config.headers["Content-Type"] = 'application/json'
        config.headers['x-user-id'] = userStore.GetUserInfo().value.userID
        config.headers["x-token"] = userStore.GetToken().value
        return config
    }
)
