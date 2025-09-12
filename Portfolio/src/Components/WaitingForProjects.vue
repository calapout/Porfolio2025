<template>
  <div class="error-404">
    <h1>{{ props.label }}</h1>
    <n-spin
      size="large"
    />
  </div>
</template>
<script setup lang="ts">
const props = defineProps<{ label: string }>();
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