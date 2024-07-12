import { useUserStore } from "@/pinia/modules/user"
import {OpenAPI} from "./codegen";
import type { ApiRequestOptions } from "./codegen/core/ApiRequestOptions";

OpenAPI.BASE = import.meta.env.VITE_BASE_API
OpenAPI.HEADERS = async (options: ApiRequestOptions):Promise<Record<string, string>> => {
    const userStore = useUserStore()
    let tokenHeader: Record<string, string> = {
        'x-token': userStore.GetToken().value
    }
    return tokenHeader
}


export * from './codegen'