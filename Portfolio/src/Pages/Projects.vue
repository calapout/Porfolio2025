<template>
  <template v-if="projects.length == 0">
    <div
      style="position: fixed; top: 50%; left: 50%; transform: translate(-50%, -50%); display: flex; flex-direction: column; align-items: center;"
    >
      <h1>{{ t('loadingProjects') }}</h1>
      <n-spin
        size="large"
      />
    </div>
  </template>
  <template v-else>
    <div style="z-index: 2">
      <h1>{{ t('projects') }}</h1>
      <div style="margin-bottom: 3rem;">
        <h2>{{ t('newOrShowoff') }}</h2>
        <primary-carousel :projects="newAndTrendyProjects" />
      </div>

      <n-card
        v-if="projectsWithPrize.length > 0"
        style="margin-bottom: 1rem;"
      >
        <h2>{{ t('mentionReceiver') }}</h2>
        <secondary-carousel :projects="projectsWithPrize" />
      </n-card>

      <n-card
        v-if="projectsMadeInCompany.length > 0"
        style="margin-bottom: 1rem;"
      >
        <h2>{{ t('madeInCompany') }}</h2>
        <secondary-carousel :projects="projectsMadeInCompany" />
      </n-card>

      <div
        v-if="personnalProjects.length > 0"
        style="margin-bottom: 1rem;"
      >
        <h2>{{ t('personnalProjects') }}</h2>
        <n-grid
          :cols="4"
          x-gap="1rem"
          y-gap="1rem"
          responsive="screen"
          style="margin-top: 1rem;"
        >
          <n-grid-item
            v-for="project in personnalProjects"
            :key="project.Slug"
            style="cursor: pointer;"
            @click.prevent="() => router.push(projectUrl + project.Slug)"
          >
            <img
              height="100%"
              :src="apiBaseUrl + project.Thumbnail.url"
              :alt="project.Thumbnail.alternativeText"
            >
          </n-grid-item>
        </n-grid>
      </div>
    </div>
  </template>
  <div
    class="parallax-container"
    :style="{top: lerp(35, 25, scrollPurcentage) + 'dvh'}"
  >
    <img
      class="robot-left"
      src="/Robot_Left.webp"
      :alt="t('redRobotAlt')"
    >
    <img
      class="robot-right"
      src="/Robot_Right.webp"
      :alt="t('blueRobotAlt')"
    >
  </div>
</template>

<script
    setup
    lang="ts"
>
import {useI18n} from "vue-i18n";
import {computed, onBeforeUnmount, onMounted, ref} from 'vue'
import {storeToRefs} from 'pinia'
import {useAppStore} from '@Stores/App'
import type {ProjectModel} from "@/index";
import PrimaryCarousel from "@Components/PrimaryCarousel.vue";
import SecondaryCarousel from "@Components/SecondaryCarousel.vue";
import {useRouter} from "vue-router";

const {t, locale} = useI18n({useScope: 'global'});

const appStore = useAppStore()
const router = useRouter()

const {
  apiBaseUrl
} = storeToRefs(appStore)

const projects = ref<ProjectModel[]>([])
const scrollPurcentage = ref(0.0)

const newAndTrendyProjects = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.IsNewAndTrendy)
})

const projectsWithPrize = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.HasPrizes)
})

const projectsMadeInCompany = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.IsMadeInCompany)
})

const personnalProjects = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => !project.IsMadeInCompany)
})

const projectUrl = computed(() => {
  if (locale.value == "fr") {
    return "/fr/projet/";
  }

  if (locale.value == "en") {
    return "/en/project/";
  }
  return "";
})

onMounted(() => {
  fetch(`${apiBaseUrl.value}/projects?locale=${locale.value}`)
      .then(response => {
        if (response.status != 200) {
          throw new Error(`Status code: ${response.status}`)
        }

        return response.json()
      })
      .then((data: ProjectModel[]) => {
        projects.value = data;
      })
      .catch(error => {
        console.error(`Error fetching data: ${error}`);
      })

  document.addEventListener('scroll', onScroll);
})

onBeforeUnmount(() => {
  document.removeEventListener("scroll", onScroll);
})

function lerp(x: number, y: number, a: number) {
  return x * (1 - a) + y * a;
}

function onScroll() {
  const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
  const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight;
  const clientHeight = document.documentElement.clientHeight || document.body.clientHeight;

  console.log(scrollTop / (scrollHeight - clientHeight))
  scrollPurcentage.value = (scrollTop / (scrollHeight - clientHeight));
}
</script>


<style
    scoped
    lang="scss"
>
.parallax-container {
  .robot-left {
    position: absolute;
    left: 0;
    transform: translateX(-40%);
    height: 100%;
    object-fit: cover;
  }

  .robot-right {
    position: absolute;
    right: 0;
    transform: translateX(40%);
    height: 100%;
    object-fit: cover;
  }

  position: fixed;
  left: 0;
  width: 100%;
  height: 30rem;
  z-index: 1;
}
</style>