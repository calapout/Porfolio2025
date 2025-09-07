<template>
  <h1>{{ t('projects') }}</h1>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onMounted} from 'vue'
import {storeToRefs} from 'pinia'
import {useAppStore} from '@Stores/App'

const {t, locale} = useI18n({useScope: 'global'});

const appStore = useAppStore()

const {
  apiBaseUrl
} = storeToRefs(appStore)

onMounted(() => {
  fetch(`${apiBaseUrl.value}/projects?locale=${locale.value}`)
      .then(response => response.json())
      .then(data => {
        console.log(data)
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      })
})

</script>


<style scoped>

</style>