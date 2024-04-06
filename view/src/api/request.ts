import axios from "axios";
import { useUserStore } from "@/pinia/modules/user"

export const http = axios.create({
    baseURL: 'http://127.0.0.1:8888',
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
