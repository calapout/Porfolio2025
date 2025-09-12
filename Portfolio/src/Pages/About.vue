<template>
  <h1>{{ t('about') }}</h1>
  <n-card class="about-me-card">
    <figure>
      <n-image
        src="/jimmy_visage.jpg"
        width="200"
        preview-disabled
        :alt="t('portraitAlt')"
        style="border-radius: 1rem"
      />
      <figcaption>
        <h2>Jimmy Tremblay-Bernier</h2>
        <h3>Diplomé en génie logiciel</h3>
      </figcaption>
    </figure>
    <div>
      <template v-for="(text, i) in texts">
        <p
          v-if="typeof text == 'string'"
          :key="'p-' + i"
        >
          {{ text }}
        </p>
        <ul
          v-else-if="typeof text == 'object'"
          :key="'ul-' + i"
        >
          <li
            v-for="(item, j) in text"
            :key="'li-' + i + '-' + j"
          >
            {{ item }}
          </li>
        </ul>
      </template>
    </div>
  </n-card>
  <template v-if="projectsIReallyLike.length > 0">
    <h2>{{ t("projectIReallyLike") }}</h2>
    <n-card
      v-for="project in projectsIReallyLike"
      :key="project.Slug"
      class="favorite-project"
      style="margin-bottom: 1rem;"
    >
      <h3>{{ project.Title }}</h3>
      <div>
        <RemoteImage
          :src="project.Thumbnail.url"
          :alt="project.Thumbnail.alternativeText"
        />
        <p>{{ project.FavoriteText }}</p>
      </div>
      <router-link
        :to="'/fr/projet/' + project.Slug"
        class="see-more"
      >
        {{ t('seeMoreDetails') }}
        <n-icon>
          <ArrowRight />
        </n-icon>
      </router-link>
    </n-card>
  </template>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onMounted, ref, watch} from "vue";
import {ApiRequestGet} from "@/utils.ts";
import type {ProjectModel} from "@/index";
import RemoteImage from "@Components/RemoteImage.vue";
import {ArrowRight} from "@vicons/fa"
import ColorPalette from "@/ColorPalette.ts";

const {t, locale} = useI18n({useScope: 'global'});

const projectsIReallyLike = ref<ProjectModel[]>([]);

const texts = computed(() => {
  return JSON.parse(t('aboutMeRTF')) as (string | string[])[];
})

onMounted(() => {
  getProjects()
})

watch(
    () => locale.value,
    () => {
      getProjects()
    }
)

function getProjects() {
  ApiRequestGet<ProjectModel[]>(`/projects/favorites`)
      .then((newProjects) => {
        projectsIReallyLike.value = newProjects;
      })
      .catch(error => {
        console.error(`Error fetching data: ${error}`);
      })
}
</script>

<!--MOBILE-->
<style scoped lang="scss">
h1 {
  align-self: start;
}

.about-me-card {
  :deep(.n-card__content) {
    div {
      display: flex;
      flex-direction: column;
      gap: 1rem;
      flex-basis: 70%;
    }

    figure {
      figcaption {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        align-items: center;
      }

      align-self: center;
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 1rem;
      flex-basis: 30%;
    }

    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  margin-bottom: 1rem;
}

.favorite-project {
  :deep(.n-card__content) {
    .n-image {
      flex-basis: 40%;
    }

    p {
      flex-basis: 60%;
    }

    .see-more {
      color: v-bind('ColorPalette.text.light');
      align-self: flex-end;
    }

    & > div {
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }

    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
}

</style>

<!-- DESKTOP -->
<style scoped lang="scss">
@media (min-width: 768px) {
  .about-me-card {
    :deep(.n-card__content) {
      flex-direction: row;
    }
  }

  .favorite-project {
    &:nth-of-type(odd) {
      p {
        order: -1;
      }
    }

    :deep(.n-card__content) {
      & > div {
        flex-direction: row;
      }
    }
  }
}
</style>