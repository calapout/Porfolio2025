import {defineStore, type StoreDefinition} from "pinia";
import {ref} from "vue";
import Config from '../../config'

export const useAppStore: StoreDefinition = defineStore("app", () => {
    const routeInOtherLanguage = ref(location.pathname);
    const apiBaseUrl = ref(Config.API_URL);
    const currentRouteKey = ref("");

    return {
        routeInOtherLanguage,
        apiBaseUrl,
        currentRouteKey
    }
});