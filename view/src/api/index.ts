import { useUserStore } from "@/pinia/modules/user"
import {OpenAPI} from "./codegen";
import type { ApiRequestOptions } from "./codegen/core/ApiRequestOptions";

OpenAPI.BASE = import.meta.env.VITE_BASE_API
OpenAPI.TOKEN = async (options: ApiRequestOptions):Promise<string> => {
    const userStore = useUserStore()
    return userStore.GetToken().value
}


export * from './codegen'