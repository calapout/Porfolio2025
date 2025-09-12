<template>
  <error-404
    v-if="is404"
    :error="t('project-details.projectNotfound')"
  />
  <template v-else-if="project">
    <h1>
      {{ project.Title }}
      (
      {{ project.RealizationPeriod.StartingYear }}
      <template v-if="project.RealizationPeriod.EndingYear != 0">
        - {{ project.RealizationPeriod.EndingYear }}
      </template>
      <template v-else>
        - {{ t('project-details.inProgress') }}
      </template>
      )
    </h1>
    <div
      v-if="project.HasPrizes"
      class="banner"
    >
      <h2>{{ t('project-details.prizeOrMention') }}:</h2>
      <p
        v-for="prize in project.Prizes"
        :key="prize.Name"
      >
        {{ prize.Name }}
      </p>
    </div>
    <div class="project-details">
      <n-card class="project-details-header">
        <div class="project-details-header-row">
          <carousel-project-details :medias="project.Medias || []" />
          <div class="project-details-header-description">
            <p v-html="description" />
          </div>
        </div>
        <n-divider />
        <div class="project-details-header-footer">
          <div
            class="technologies"
          >
            <p>{{ t('project-details.technologies') }}:</p>
            <n-tag
              v-for="technology in project.Technologies"
              :key="technology"
              round
              type="primary"
              :bordered="false"
            >
              {{ technology }}
            </n-tag>
          </div>
          <div
            v-if="project.IsMadeInCompany"
            class="company"
          >
            <p>{{ t('project-details.realizedFor') }}:</p>
            <n-tag
              round
              type="primary"
              :bordered="false"
            >
              <a
                target="_blank"
                :href="project.Company.Website"
              >
                {{ project.Company.Name }}
              </a>
            </n-tag>
            <n-tag
              v-if="project.AdditionnalCompanyWebsite"
              round
              type="primary"
              :bordered="false"
            >
              <a
                target="_blank"
                :href="project.AdditionnalCompanyWebsite"
              >
                {{ t('project-details.projectWebsite') }}
              </a>
            </n-tag>
          </div>
        </div>
      </n-card>
      <n-card>
        <h2>{{ t('project-details.taskRealized') }}</h2>
        <ul>
          <li
            v-for="task in tasks"
            :key="task"
          >
            {{ task }}
          </li>
        </ul>
      </n-card>
    </div>
  </template>
  <WaitingForProjects
    v-else
    :label="t('project-details.loadingProject')"
  />
</template>

<script setup lang="ts">
import {computed, onMounted, ref, watch} from "vue";
import type {ProjectModel} from "@/index";
import {useRoute} from "vue-router";
import {useI18n} from "vue-i18n";
import ColorPalette from "@/ColorPalette.ts";

import {ApiRequestGet} from "@/utils.ts";

import WaitingForProjects from "@Components/WaitingForProjects.vue";
import CarouselProjectDetails from "@Components/CarouselProjectDetails.vue";
import Error404 from "@Components/Error404.vue";

const {t, locale} = useI18n({useScope: 'global'});
const {params} = useRoute();
const {slug} = params;

const project = ref<ProjectModel | null>(null);
const is404 = ref<boolean>(false);

const description = computed(() => {
  if (project.value == null) return "";

  return project.value.Description.replaceAll("\n", "<br />");
});

const tasks = computed(() => {
  if (project.value == null) return "";

  return project.value.TaskRealized.split("\n");
})

onMounted(() => {
  getProjects();
})

watch(
    () => locale.value,
    () => {
      getProjects()
    })

function getProjects() {
  is404.value = false;

  ApiRequestGet<ProjectModel>(`/projects/${slug}`)
      .then((data) => {
        if (data.statusCode == 200 && data.data != null) {
          project.value = data.data;
        } else if (data.statusCode == 404) {
          is404.value = true;
        }
      })
      .catch(error => {
        console.error(`Error fetching data: ${error}`);
      })
}

</script>

<style scoped lang="scss">
a {
  color: v-bind('ColorPalette.text.light');
}

.banner {
  background-color: v-bind('ColorPalette.primary["40"]');
  padding: 1rem;
  display: flex;
  gap: 1rem;
  align-items: center;
}

.project-details {
  .project-details-header {
    .project-details-header-row {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      gap: 1rem;
    }

    .project-details-header-footer {
      .technologies, .company {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
        align-items: center;
      }

      .n-tag {
        :deep(.n-tag__content) {
          overflow-wrap: break-word;
          text-wrap: wrap;
          height: fit-content;
        }

        height: fit-content;
        padding: 0.5rem;
      }

      display: flex;
      flex-direction: column;
      align-items: start;
      gap: 1rem;
    }
  }

  display: flex;
  flex-direction: column;
  gap: 1rem;
  position: relative;
  z-index: 2;
}
</style>

<style scoped lang="scss">
@media (min-width: 768px) {
  .project-details {
    .project-details-header {
      .project-details-header-row {
        .carousel {
          flex-basis: 60%;
        }

        .project-details-header-description {
          flex-basis: 40%;
        }

        flex-direction: row;
      }

      .project-details-header-footer {


        flex-direction: row;
        align-items: center;
        gap: 3rem;
      }
    }
  }
}
</style>