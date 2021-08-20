<template>
  <v-container>
    <HeroCenter />
    <Carousel title="Рекомендации для вас" />
    <Carousel title="Категории" />
    <Carousel />
    <Carousel title="Топ - 10" />
  </v-container>
</template>

<script>
import HeroCenter from '~/components/Homepage/HeroCenter.vue'
import Carousel from '~/components/Carousel.vue'
export default {
  components: { HeroCenter, Carousel },
}
</script>

<style></style>
