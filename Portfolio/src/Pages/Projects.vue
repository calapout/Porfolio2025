<template>
  <template v-if="projects.length > 0">
    <div
      class="projects"
      style="z-index: 2"
    >
      <h1>{{ t('projects.projects') }}</h1>
      <div>
        <h2 style="margin-bottom: 0.5rem">
          {{ t('projects.newOrShowoff') }}
        </h2>
        <primary-carousel-projects :projects="newAndTrendyProjects" />
      </div>

      <n-card
        v-if="projectsWithPrize.length > 0"
      >
        <h2>{{ t('projects.mentionReceiver') }}</h2>
        <secondary-carousel :projects="projectsWithPrize">
          <template #tag="{project}">
            <n-tag
              v-for="prize in project.Prizes"
              :key="prize"
              class="carousel-tag"
              round
              type="primary"
              :bordered="false"
              @click.prevent
            >
              {{ prize.Name }}
            </n-tag>
          </template>
        </secondary-carousel>
      </n-card>

      <n-card
        v-if="projectsMadeInCompany.length > 0"
      >
        <h2>{{ t('projects.madeInCompany') }}</h2>
        <secondary-carousel :projects="projectsMadeInCompany">
          <template #tag="{project}">
            <n-tag
              class="carousel-tag"
              round
              type="primary"
              :bordered="false"
            >
              {{ project.Company.Name }}
            </n-tag>
          </template>
        </secondary-carousel>
      </n-card>

      <div
        v-if="projects.length > 0"
      >
        <h2 class="all-projects-title">
          {{ t('projects.allProjects') }}
        </h2>
        <n-grid
          :cols="4"
          x-gap="16"
          y-gap="16"
          responsive="screen"
          style="margin-top: 1rem;"
        >
          <n-grid-item
            v-for="project in projects"
            :key="project.Slug"
            style="cursor: pointer;"
            @click.prevent="() => router.push(projectUrl + project.Slug)"
          >
            <RemoteImage
              height="100%"
              :src="project.Thumbnail.url"
              :alt="project.Thumbnail.alternativeText"
              object-fit="cover"
            />
          </n-grid-item>
        </n-grid>
      </div>
    </div>
  </template>
  <WaitingForProjects
    v-else
    :label="t('projects.loadingProjects')"
  />
</template>

<script
    setup
    lang="ts"
>
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref, watch} from 'vue'
import type {ProjectModel} from "@/index";
import PrimaryCarouselProjects from "@Components/PrimaryCarouselProjects.vue";
import SecondaryCarousel from "@Components/SecondaryCarousel.vue";
import {useRouter} from "vue-router";
import {ApiRequestGet} from "@/utils.ts";
import RemoteImage from "@Components/RemoteImage.vue";
import WaitingForProjects from "@Components/WaitingForProjects.vue";

const {t, locale} = useI18n({useScope: 'global'});

const router = useRouter()

const projects = ref<ProjectModel[]>([])

const newAndTrendyProjects = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.IsNewAndTrendy)
})

const projectsWithPrize = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.HasPrizes)
})

const projectsMadeInCompany = computed<ProjectModel[]>(() => {
  return projects.value.filter(project => project.IsMadeInCompany)
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
  getProjects();
})

watch(
    () => locale.value,
    () => {
      getProjects()
    }
)

function getProjects() {
  ApiRequestGet<ProjectModel[]>(`/projects`)
      .then((newProjects) => {
        if (newProjects.statusCode == 200 && newProjects.data != null) {
          projects.value = newProjects.data;
        }
      })
      .catch(error => {
        console.error(`Error fetching data: ${error}`);
      })
}
</script>


<style
    scoped
    lang="scss"
>

.projects {
  h1 {
    margin-left: 42px;
  }

  h2 {
    margin-left: 42px;
  }

  .all-projects-title {
    margin-left: 0px;
  }

  :deep(.n-card__content) {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.carousel-tag {
  position: absolute;
  bottom: 0.5rem;
  left: 0.5rem;
}

</style>