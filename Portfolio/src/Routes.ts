import HomePage from "@Pages/Projects.vue"
import AboutPage from '@Pages/About.vue'
import type {RouteComponent} from "vue-router";

export type RouteEntry = {
    routes: { "fr": string, "en": string },
    component: RouteComponent
}

export const routeArray: Array<RouteEntry> = [
    {
        "routes": {
            "fr": "/fr",
            "en": "/en"
        },
        component: HomePage
    },
    {
        "routes": {
            "fr": "/fr/a-propos",
            "en": "/en/about"
        },
        component: AboutPage
    }
];