import {useAppStore} from "@Stores/App.ts";
import {storeToRefs} from "pinia";
import {i18n} from "@/main.ts";
import router from "@/Router";

const cache = new Map<string, unknown>();

export function ApiRequestGet<T>(endpoint: string): Promise<T> {
    return new Promise<T>((resolve, reject) => {
        const locale = i18n.global.locale.value;
        const appStore = useAppStore()

        const {
            apiBaseUrl
        } = storeToRefs(appStore)

        if (cache.has(`${endpoint}?locale=${locale}`)) {
            return resolve(cache.get(`${endpoint}?locale=${locale}`) as T);
        }

        fetch(`${apiBaseUrl.value}${endpoint}?locale=${locale}`)
            .then(response => {
                if (response.status != 200) {
                    reject(new Error(`Status code: ${response.status}`));
                }

                return response.json();
            })
            .then(data => {
                cache.set(`${endpoint}?locale=${locale}`, data);
                resolve(data);
            })
            .catch(error => {
                reject(error);
            })
    })
}

export function GoToProjectPage(slug: string) {
    const locale = i18n.global.locale.value;

    const projectUrl = locale == "fr" ? "/fr/projet/" : "/en/project/";

    return router.push(projectUrl + slug)
}