<template>
  <n-config-provider :theme="darkTheme">
    <nav class="nav-bar">
      <n-menu :options="menuOptions" mode="horizontal" responsive/>
    </nav>
    <n-global-style />
  </n-config-provider>
  <router-view/>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h} from "vue";
import {storeToRefs} from "pinia";
import {darkTheme, NConfigProvider, NMenu, NGlobalStyle} from 'naive-ui'
import {RouterLink, RouterView} from "vue-router";


import {useAppStore} from "@Stores/App.ts";

const {locale} = useI18n({useScope: 'global'});
const appStore = useAppStore();

const {routeInOtherLanguage} = storeToRefs(appStore)
const {t} = useI18n({useScope: 'global'});

const oppositeLocale = computed(() => {
  return locale.value == 'fr' ? 'en' : 'fr';
})

const menuOptions = computed(() => {
      if (locale.value == 'fr') {
        return [
          {
            label: () => h(
                RouterLink,
                {
                  to: "/fr"
                },
                {
                  default: () => t('projects')
                }
            ),
            key: 'projects',
          },
          {
            label: () => h(
                RouterLink,
                {
                  to: "/fr/a-propos"
                },
                {
                  default: () => t('about')
                }
            ),
            key: 'about',
          },
          {
            label: () => h(
                RouterLink,
                {
                  to: routeInOtherLanguage.value
                },
                {
                  default: () => oppositeLocale.value.toUpperCase()
                }
            ),
            key: 'opposite-locale',
          },
        ]
      } else if (locale.value == 'en') {
        return [
          {
            label: () => h(
                RouterLink,
                {
                  to: "/en"
                },
                {
                  default: () => t('projects')
                }
            ),
            key: 'projects',
          },
          {
            label: () => h(
                RouterLink,
                {
                  to: "/en/about"
                },
                {
                  default: () => t('about')
                }
            ),
            key: 'about',
          },
          {
            label: () => h(
                RouterLink,
                {
                  to: routeInOtherLanguage.value
                },
                {
                  default: () => oppositeLocale.value.toUpperCase()
                }
            ),
            key: 'opposite-locale',
          },
        ]
      }
      return []
    }
)
</script>

<style scoped>
.nav-bar {
  background-color: rgba(0, 106, 15, 0.44);
  display: flex; justify-content: space-between;
  align-items: center;
  padding: 0 1rem 0 0;
}

:deep(.n-menu .n-menu-item:nth-of-type(3)) {
margin-left: auto;
}
</style>
