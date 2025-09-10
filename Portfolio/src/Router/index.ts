import {createRouter, createWebHistory, type RouteRecordRaw} from 'vue-router'

import {routeArray} from "@/Routes"
import {useAppStore} from "@Stores/App";
import {i18n} from "@/main";


const supportedLocales: Array<"fr" | "en"> = ['fr', 'en'];
const routes: RouteRecordRaw[] = [];
for (const route of routeArray) {
    for (const locale of supportedLocales) {
        if (locale in route.routes) {
            routes.push({path: route.routes[locale], component: route.component})
        }
    }
}
routes.push({path: '/:pathMatch(.*)*', redirect: '/fr'})

const index = createRouter({
    history: createWebHistory(),
    routes,
})


index.afterEach((to) => {
    const appStore = useAppStore();

    // remove trailing /
    let path = to.path;
    if (path.endsWith("/")) {
        path = path.substring(0, path.length - 1);
    }

    // split the path
    const splittedPaths = path.substring(1).split("/");
    if (splittedPaths.length == 0) {
        throw new Error("Invalid path");
    }

    // locale should be first element
    const newLocale = splittedPaths[0];
    if (newLocale == 'fr' || newLocale == 'en') {
        i18n.global.locale.value = newLocale;
    }

    // find the opposite route
    const oppositeLocale = newLocale == 'fr' ? 'en' : 'fr';
    for (const route of routeArray) {
        if(route.routes[newLocale as 'en'|'fr'] == path && oppositeLocale in route.routes) {
            appStore.currentRouteKey = route.key;
            appStore.routeInOtherLanguage = route.routes[oppositeLocale];
            return;
        }
    }

    // backup to home page in the other locale
    appStore.routeInOtherLanguage = `/${oppositeLocale}`
    appStore.currentRouteKey = "";
})

export default index;