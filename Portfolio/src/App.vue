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
          <ParallaxRobots />
        </div>
      </div>
    </main>
    <footer>
      <p>Jimmy Tremblay-Bernier © {{ new Date().getFullYear() }}</p>
      <p>Viviane Badea © {{ t('global.veeMention') }}</p>
    </footer>
  </n-config-provider>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, type DefineComponent, h} from "vue";
import {storeToRefs} from "pinia";
import {darkTheme, NConfigProvider, NMenu, NGlobalStyle, NIcon} from 'naive-ui'
import {RouterLink, RouterView} from "vue-router";

import {useAppStore} from "@Stores/App.ts";
import {ThemeOverrides} from "@/ThemeOverrides.ts";
import ColorPalette from "@/ColorPalette.ts";

import ParallaxRobots from "@Components/ParallaxRobots.vue";
import {Envelope, Github, LinkedinIn} from "@vicons/fa";

const {t, locale} = useI18n({useScope: 'global'});
const appStore = useAppStore();

const {routeInOtherLanguage, currentRouteKey} = storeToRefs(appStore)

const oppositeLocale = computed(() => {
  return locale.value == 'fr' ? 'en' : 'fr';
})

const menuOptions = computed(() => {
      const routes = [];

      if (locale.value == 'fr') {
        routes.push(
            generateTextMenuEntry(t('menu.projects'), 'projects', '/fr'),
            generateTextMenuEntry(t('menu.about'), 'about', '/fr/a-propos'),
        )
      } else if (locale.value == 'en') {
        routes.push(
            generateTextMenuEntry(t('menu.projects'), 'projects', '/en'),
            generateTextMenuEntry(t('menu.about'), 'about', '/en/about'),
        );
      }

      routes.push(
          generateTextMenuEntry(t('menu.downloadCV'), 'CV', '/CV.pdf', false),
          generateIconMenuEntry(LinkedinIn, 'linkedin', 'https://www.linkedin.com/in/jimmy-tremblay-bernier/', false),
          generateIconMenuEntry(Github, 'github', 'https://github.com/calapout', false),
          generateIconMenuEntry(Envelope, 'email', 'mailto:tremblaybernierjimmy@protonmail.com', false),
          generateTextMenuEntry(oppositeLocale.value.toUpperCase(), 'opposite-locale', routeInOtherLanguage.value),
      )

      return routes;
    }
)

function generateTextMenuEntry(label: string, key: string, to: string, local: boolean = true) {
  if (local) {
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
  return {
    label: () =>
        h(
            "a",
            {
              href: to,
              target: "_blank",
            },
            {
              default: () => label
            }
        ),
    key,
  }
}

function generateIconMenuEntry(icon: Partial<DefineComponent>, key: string, to: string, local: boolean = true) {
  if (local) {
    return {
      label: () =>
          h(
              RouterLink,
              {
                to
              },
              h(NIcon, null, {default: () => h(icon)})
          ),
      key,
    }
  }
  return {
    label: () =>
        h(
            "a",
            {
              href: to,
              target: "_blank",
            },
            h(NIcon, null, {default: () => h(icon)})
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
  border-radius: 0.75rem;

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
}
</style>