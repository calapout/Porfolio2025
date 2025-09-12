<template>
  <n-image
    v-bind="props"
    ref="image"
    :src="`${apiBaseUrl}${props.src}`"
    lazy
  />
</template>

<script setup lang="ts">

import {storeToRefs} from "pinia";
import {useAppStore} from "@Stores/App.ts";
import {type ImgHTMLAttributes, ref} from "vue";

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
  onLoad: (e: Event) => void,
}

const props = defineProps<Partial<ImageProps>>()
const {
  apiBaseUrl,
} = storeToRefs(appStore)

const image = ref<HTMLElement | null>(null)

</script>

<style scoped lang="scss">
.n-image {
  :deep(img) {
    width: 100%;
    height: 100%;
  }
  width: 100%;
  height: 100%;
}
</style>