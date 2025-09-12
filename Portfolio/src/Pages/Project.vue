<template>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import type {ProjectModel} from "@/index";
import {useRoute} from "vue-router";
import {ApiRequestGet} from "@/utils.ts";

const project = ref<ProjectModel | null>(null);

const {params} = useRoute();
const {slug} = params;

onMounted(() => {
  ApiRequestGet<ProjectModel>(`/projects/${slug}`)
      .then((data) => {
        project.value = data;
      })
      .catch(error => {
        console.error(`Error fetching data: ${error}`);
      })
})
</script>

<style scoped lang="scss">

</style>