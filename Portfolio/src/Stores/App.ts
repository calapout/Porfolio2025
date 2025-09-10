import {defineStore, type StoreDefinition} from "pinia";
import {ref} from "vue";

export const useAppStore: StoreDefinition = defineStore("app", () => {
    const routeInOtherLanguage = ref(location.pathname);
    const apiBaseUrl = ref("http://localhost:8080/api/v1");
    const currentRouteKey = ref("");

    return {
        routeInOtherLanguage,
        apiBaseUrl,
        currentRouteKey
    }
});