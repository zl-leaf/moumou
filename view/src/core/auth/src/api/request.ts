import axios from "axios";

export const http = axios.create({
    baseURL: import.meta.env.VITE_BASE_API,
    timeout: 99999,
})


http.interceptors.request.use(
    config => {
        config.headers["Content-Type"] = 'application/json'
        return config
    }
)
