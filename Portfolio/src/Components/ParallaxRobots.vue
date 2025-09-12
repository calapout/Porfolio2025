<template>
  <div
    class="parallax-container"
    :style="{top: lerp(35, 25, scrollPurcentage) + 'dvh'}"
  >
    <n-image
      preview-disabled
      class="robot-left"
      src="/Robot_Left.webp"
      :alt="t('global.redRobotAlt')"
    />
    <n-image
      preview-disabled
      class="robot-right"
      src="/Robot_Right.webp"
      :alt="t('global.blueRobotAlt')"
    />
  </div>
</template>
<script setup lang="ts">
import {onBeforeUnmount, onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";

const {t} = useI18n({useScope: 'global'});

const scrollPurcentage = ref(0.0)

onMounted(() => {
  document.addEventListener('scroll', onScroll);
})

onBeforeUnmount(() => {
  document.removeEventListener("scroll", onScroll);
})

function lerp(x: number, y: number, a: number) {
  return x * (1 - a) + y * a;
}

function onScroll() {
  const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
  const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight;
  const clientHeight = document.documentElement.clientHeight || document.body.clientHeight;

  scrollPurcentage.value = (scrollTop / (scrollHeight - clientHeight));
}

</script>
<style
    scoped
    lang="scss"
>
.parallax-container {
  .robot-left {
    position: absolute;
    left: 0;
    transform: translateX(-40%);
    height: 100%;
    object-fit: cover;
  }

  .robot-right {
    position: absolute;
    right: 0;
    transform: translateX(40%);
    height: 100%;
    object-fit: cover;
  }

  position: fixed;
  left: 0;
  width: 100%;
  height: 30rem;
  z-index: 1;
  pointer-events: none;
}

</style>