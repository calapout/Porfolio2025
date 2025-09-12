<template>
  <div class="error-404">
    <h2>404 - {{ props.error }}</h2>
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ error: string }>();
</script>

<style scoped lang="scss">
.error-404 {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 1rem;
}
</style>