import ProjectsPage from "@Pages/Projects.vue"
import ProjectDetailsPage from "@Pages/Project-Details.vue"
import AboutPage from '@Pages/About.vue'
import type {RouteComponent} from "vue-router";

export type RouteEntry = {
    routes: { "fr": string, "en": string },
    titles: { "fr": string, "en": string}
    component: RouteComponent,
    key: string,
}

export const routeArray: Array<RouteEntry> = [
    {
        "routes": {
            "fr": "/fr",
            "en": "/en"
        },
        titles: {
          fr: "Portfolio - Projets",
          en: "Portfolio - Projects"
        },
        component: ProjectsPage,
        key: "projects"
    },
    {
        "routes": {
            "fr": "/fr/a-propos",
            "en": "/en/about"
        },
        titles: {
            fr: "Portfolio - À propos",
            en: "Portfolio - About me"
        },
        component: AboutPage,
        key: "about"
    },
    {
        "routes": {
            "fr": "/fr/projet/:slug",
            "en": "/en/project/:slug"
        },
        titles: {
            fr: "Portfolio - Projet {slug}",
            en: "Portfolio - Project {slug}"
        },
        component: ProjectDetailsPage,
        key: "project"
    }
];