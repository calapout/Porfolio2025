import { createApp } from 'vue'
import { createI18n } from 'vue-i18n'
import naive from 'naive-ui'

import './style.scss'
import App from './App.vue'

const app = createApp(App)


const i18n = createI18n({
    // something vue-i18n options here ...
})

app.use(naive)
app.use(i18n)

app.mount('#app')
