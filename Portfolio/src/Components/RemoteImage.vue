<template>
  <n-image
    v-bind="props"
    :src="`${apiBaseUrl}${props.src}`"
  />
</template>

<script setup lang="ts">

import {storeToRefs} from "pinia";
import {useAppStore} from "@Stores/App.ts";
import type {ImgHTMLAttributes} from "vue";

const appStore = useAppStore()

type ImageProps = {
  onPreviewPrev: () => void,
  onPreviewNext: () => void,
  showToolbar: boolean
  showToolbarTooltip: boolean,
  alt: string,
  height: string | number,
  imgProps: ImgHTMLAttributes,
  previewedImgProps: ImgHTMLAttributes,
  lazy: boolean,
  objectFit: "fill" | "contain" | "cover" | "none" | "scale-down",
  previewSrc: string,
  fallbackSrc: string,
  width: string | number,
  src: string,
  previewDisabled: boolean,
  loadDescription: string,
  onError: (e: Event) => void,
  onLoad: (e: Event) => void
}

const props = defineProps<Partial<ImageProps>>()
const {
  apiBaseUrl,
} = storeToRefs(appStore)

</script>

<style scoped lang="scss">
</style>