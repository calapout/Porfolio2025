<template>
  <div
    class="carousel"
    :style="{'position': 'relative'}"
  >
    <div class="row">
      <n-button
        class="carousel-arrow carousel-arrows-left"
        @click="carouselRef?.prev()"
      >
        <n-icon>
          <angle-left />
        </n-icon>
      </n-button>
      <n-carousel
        ref="carousel"
        :space-between="16"
        :slides-per-view="4"
        :centered-slides="true"
        :loop="true"
        :show-arrow="false"
        :show-dots="false"
        draggable
      >
        <n-carousel-item
          v-for="(project) in props.projects"
          :key="project.Slug"
          @click.prevent="() => router.push(projectUrl + project.Slug)"
        >
          <img
            height="100%"
            :src="apiBaseUrl + project.Thumbnail.url"
            :alt="project.Thumbnail.alternativeText"
          >
        </n-carousel-item>
        <template #dots="{currentIndex, total, to}">
          <div class="carousel-dots">
            <span
              v-for="n in total"
              :key="n"
              class="carousel-dot"
              :class="{active: n-1 == currentIndex}"
              @click="to(n-1)"
            />
          </div>
        </template>
      </n-carousel>
      <n-button
        class="carousel-arrow carousel-arrows-right"
        @click="carouselRef?.next()"
      >
        <n-icon>
          <angle-right />
        </n-icon>
      </n-button>
    </div>
    <div class="carousel-dots">
      <span
        v-for="n in projects.length"
        :key="n"
        class="carousel-dot"
        :class="{active: n-1 == carouselRef?.getCurrentIndex()}"
        @click="carouselRef?.to(n-1)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">

import {AngleLeft, AngleRight} from "@vicons/fa";
import {type CarouselInst, NButton, NCarousel, NCarouselItem, NIcon} from "naive-ui";
import type {ProjectModel} from "@/index";
import {computed, useTemplateRef} from "vue";
import {useAppStore} from "@Stores/App.ts";
import {storeToRefs} from "pinia";
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import ColorPalette from "@/ColorPalette.ts";

type Props = {
  projects: ProjectModel[]
}

const appStore = useAppStore()

const props = defineProps<Props>();

const {locale} = useI18n({useScope: 'global'})

const router = useRouter()

const {
  apiBaseUrl
} = storeToRefs(appStore)
const carouselRef = useTemplateRef<CarouselInst>("carousel")

const projectUrl = computed(() => {
  if (locale.value == "fr") {
    return "/fr/projet/";
  }

  if (locale.value == "en") {
    return "/en/project/";
  }
  return "";
})

</script>

<style scoped lang="scss">
.carousel {
  img {
    cursor: pointer;
    display: flex;
    max-height: 100%;
    max-width: 100%;
    width: 100%;
    height: auto;
  }

  .row {
    display: flex;
    flex-direction: row;
    align-items: center;
    flex-grow: 1;
    height: 100%;
  }

  .carousel-arrow {
    color: white;
    height: 5rem;
  }

  .carousel-arrows-left {
    background: rgba(20, 20, 24, 1);
    background: linear-gradient(90deg, rgba(16, 16, 20, 1) 0%, rgb(20, 20, 24) 100%);

  }

  .carousel-arrows-right {
    background: rgba(16, 16, 20, 1);
    background: linear-gradient(90deg, rgb(20, 20, 24) 0%, rgba(16, 16, 20, 1) 100%);
  }

  .carousel-dots {
    .carousel-dot {
      width: 20px;
      height: 10px;
      border-radius: 4px;
      background-color: grey;
      content: "";
      display: block;
      transition: background-color 0.3s ease;
      cursor: pointer;

      &.active {
        background-color: v-bind('ColorPalette.primary["20"]');
      }
    }

    display: flex;
    justify-content: center;
    gap: 10px;
    margin-top: 10px;
    margin-bottom: 10px;
  }

  display: flex;
  flex-direction: column;
  height: 10rem;
}
</style>