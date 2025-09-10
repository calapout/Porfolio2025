import ProjectsPage from "@Pages/Projects.vue"
import ProjectPage from "@Pages/Project.vue"
import AboutPage from '@Pages/About.vue'
import type {RouteComponent} from "vue-router";

export type RouteEntry = {
    routes: { "fr": string, "en": string },
    component: RouteComponent,
    key: string,
}

export const routeArray: Array<RouteEntry> = [
    {
        "routes": {
            "fr": "/fr",
            "en": "/en"
        },
        component: ProjectsPage,
        key: "projects"
    },
    {
        "routes": {
            "fr": "/fr/a-propos",
            "en": "/en/about"
        },
        component: AboutPage,
        key: "about"
    },
    {
        "routes": {
            "fr": "/fr/projet/:slug",
            "en": "/en/project/:slug"
        },
        component: ProjectPage,
        key: "project"
    }
];