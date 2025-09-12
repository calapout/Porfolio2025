<template>
  <video
    v-bind="props"
    :src="`${apiBaseUrl}${props.src}`"
  />
</template>

<script setup lang="ts">

import {storeToRefs} from "pinia";
import {useAppStore} from "@Stores/App.ts";

const appStore = useAppStore()

type Booleanish = boolean | 'true' | 'false';
type Numberish = number | string;

type VideoProps = {
  autoplay?: Booleanish;
  controls?: Booleanish;
  controlslist?: string;
  crossorigin?: string;
  loop?: Booleanish;
  mediagroup?: string;
  muted?: Booleanish;
  playsinline?: Booleanish;
  preload?: string;
  src?: string;
  height?: Numberish;
  poster?: string;
  width?: Numberish;
  disablePictureInPicture?: Booleanish;
  disableRemotePlayback?: Booleanish;
};

const props = defineProps<VideoProps>()
const {
  apiBaseUrl,
} = storeToRefs(appStore)

</script>

<style scoped lang="scss">
</style>