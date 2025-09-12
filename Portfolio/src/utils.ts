import {useAppStore} from "@Stores/App.ts";
import {storeToRefs} from "pinia";
import {i18n} from "@/main.ts";
import router from "@/Router";

type CacheEntry = {
    content: string;
    expirationTime: number;
}

const CACHE_DURATION = 1000 * 60 * 60 * 24; // 24 hours
const USE_CACHE = false;

class Cache {
    static get<T>(key: string): T | undefined {
        const entry = localStorage.getItem(key);

        if (entry === null) {
            return undefined;
        }

        return JSON.parse(entry).content as T;
    }

    static set<T>(key: string, value: T): void {
        localStorage.setItem(key, JSON.stringify({content: value, expirationTime: Date.now() + CACHE_DURATION}));
    }

    static has(key: string): boolean {
        const entry = localStorage.getItem(key);

        return entry !== null && (JSON.parse(entry) as CacheEntry).expirationTime > Date.now();
    }
}

type ApiGetResponse<T> = { data: T | null, statusCode: number };

export function ApiRequestGet<T>(endpoint: string): Promise<ApiGetResponse<T>> {
    return new Promise<ApiGetResponse<T>>((resolve, reject) => {
        const locale = i18n.global.locale.value;
        const appStore = useAppStore()

        const {
            apiBaseUrl
        } = storeToRefs(appStore)

        if (USE_CACHE && Cache.has(`${endpoint}?locale=${locale}`)) {
            const content = Cache.get<ApiGetResponse<T>>(`${endpoint}?locale=${locale}`);
            if (content === undefined) {
                return reject(new Error("Content is undefined"));
            }
            return resolve(content);
        }

        fetch(`${apiBaseUrl.value}${endpoint}?locale=${locale}`)
            .then(async response => {
                if (response.status == 200) {
                    const json = await response.json();

                    return {statusCode: response.status, data: json as T};
                } else if (response.status == 404) {
                    return {statusCode: response.status, data: null};
                }

                throw new Error(`Request failed with status ${response.status}`);
            })
            .then(data => {
                if (USE_CACHE) {
                    Cache.set(`${endpoint}?locale=${locale}`, data);
                }
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