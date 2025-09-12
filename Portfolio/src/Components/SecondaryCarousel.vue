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
        :slides-per-view="Math.min(projects.length, 4)"
        :centered-slides="false"
        :loop="true"
        :show-arrow="false"
        :show-dots="false"
        draggable
      >
        <n-carousel-item
          v-for="(project) in props.projects"
          :key="project.Slug"
          @click.prevent="() => GoToProjectPage(project.Slug)"
        >
          <remote-image
            preview-disabled
            :src="project.Thumbnail.url"
            :alt="project.Thumbnail.alternativeText"
            object-fit="cover"
          />
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
import {useTemplateRef} from "vue";
import ColorPalette from "@/ColorPalette.ts";
import {GoToProjectPage} from "@/utils.ts";
import RemoteImage from "@Components/RemoteImage.vue";

type Props = {
  projects: ProjectModel[]
}

const props = defineProps<Props>();
const carouselRef = useTemplateRef<CarouselInst>("carousel")

</script>

<style scoped lang="scss">
.carousel {
  .n-image {
    :deep(img) {
      width: 100%;
    }
    cursor: pointer;
    height: 100%;
    width: 100%;
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
  height: 15rem;
  padding-bottom: 30px;
}
</style>