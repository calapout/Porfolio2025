import './css/style.scss'
import {createApp} from 'vue'
import {createI18n} from 'vue-i18n'
import { createPinia } from 'pinia'

import naive from 'naive-ui'

import localeEn from "@Locales/en.json"
import localeFr from "@Locales/fr.json"

import Router from './Router'
import App from './App.vue'

const app = createApp(App)
const Pinia = createPinia();

export const i18n = createI18n({
    locale: 'fr',
    fallbackLocale: 'fr',
    messages: {
        fr: localeFr,
        en: localeEn
    },
    legacy: false,
})

app.use(naive);
app.use(i18n);
app.use(Pinia);
app.use(Router);

app.mount('#app');
