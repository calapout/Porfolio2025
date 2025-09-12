<template>
  <n-config-provider
    :theme="darkTheme"
    :theme-overrides="ThemeOverrides"
  >
    <header>
      <nav class="nav-bar">
        <n-menu
          :value="currentRouteKey"
          :options="menuOptions"
          mode="horizontal"
          responsive
        />
      </nav>
    </header>
    <n-global-style />
    <main>
      <div class="outter-wrapper">
        <div class="inner-wrapper">
          <router-view />
        </div>
      </div>
    </main>
    <footer>
      <p>Jimmy Tremblay-Bernier © {{ new Date().getFullYear() }}</p>
      <p>Viviane Badea © {{ t('veeMention') }}</p>
    </footer>
  </n-config-provider>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h} from "vue";
import {storeToRefs} from "pinia";
import {darkTheme, NConfigProvider, NMenu, NGlobalStyle} from 'naive-ui'
import {RouterLink, RouterView} from "vue-router";

import {useAppStore} from "@Stores/App.ts";
import {ThemeOverrides} from "@/ThemeOverrides.ts";
import ColorPalette from "@/ColorPalette.ts";

const {t, locale} = useI18n({useScope: 'global'});
const appStore = useAppStore();

const {routeInOtherLanguage, currentRouteKey} = storeToRefs(appStore)

const oppositeLocale = computed(() => {
  return locale.value == 'fr' ? 'en' : 'fr';
})

const menuOptions = computed(() => {
      if (locale.value == 'fr') {
        return [
          generateMenuEntry(t('projects'), 'projects', '/fr'),
          generateMenuEntry(t('about'), 'about', '/fr/a-propos'),
          {
            type: 'divider',
            key: 'divider',
            show: true,
          },
          generateMenuEntry(oppositeLocale.value.toUpperCase(), 'opposite-locale', routeInOtherLanguage.value),
        ]
      } else if (locale.value == 'en') {
        return [
          generateMenuEntry(t('projects'), 'projects', '/en'),
          generateMenuEntry(t('about'), 'about', '/en/about'),
          {
            type: 'divider',
            key: 'divider',
            show: true,
          },
          generateMenuEntry(oppositeLocale.value.toUpperCase(), 'opposite-locale', routeInOtherLanguage.value),
        ]
      }
      return []
    }
)

function generateMenuEntry(label: string, key: string, to: string) {
  return {
    label: () =>
        h(
            RouterLink,
            {
              to
            },
            {
              default: () => label
            }
        ),
    key,
  }


}
</script>

<style scoped lang="scss">
.n-config-provider {
  & header {
    & .nav-bar {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0.5rem 1rem;
    }

    position: sticky;
    left: 0;
    top: 0;
    background-color: v-bind('ColorPalette.surface["0"]');
    z-index: 1000;
  }

  main {
    .outter-wrapper {
      .inner-wrapper {
        width: 100%;
        max-width: 1000px;
        display: flex;
        flex-direction: column;
      }

      margin: 0 auto;
      display: flex;
      justify-content: center;
      padding: 1rem;
      width: 100%;
      max-width: 1200px;
      flex-grow: 1;
    }

    background: linear-gradient(90deg, v-bind('ColorPalette.surface["0"]') 0%, v-bind('ColorPalette.primary["10"]') 50%, v-bind('ColorPalette.surface["0"]') 100%);
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  footer {
    position: sticky;
    left: 0;
    bottom: 0;
    background-color: v-bind('ColorPalette.surface["0"]');
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 1rem 1rem;
    z-index: 1000;
  }

  flex-grow: 1;
  display: flex;
  flex-direction: column;
}
</style>

<style lang="scss">
#app {
  min-height: 100dvh;
  display: flex;
  flex-direction: column;
}

.n-config-provider {
  & header {
    & .nav-bar {
      & .n-menu {
        .v-overflow {
          .n-menu-item:nth-of-type(3) {
            margin-left: auto;
          }

          display: flex;
          flex-direction: row;
          gap: 1rem;
        }
      }
    }
  }
}

.n-menu-item-content {
  &:hover {
    background: var(--n-item-color-hover);
  }

  &--selected {
    &:hover {
      background: var(--n-item-color-active-hover);
    }

    background: var(--n-item-color-active);
  }

  a {
    padding: 0.75rem 0.5rem;
    border-radius: 0.5rem;
  }

  border-radius: 0.75rem;
}
</style>