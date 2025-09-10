<template>
  <h1>{{ t('about') }}</h1>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";

const {t} = useI18n({useScope: 'global'});
</script>

<style scoped lang="scss">

</style>