<template>
  <v-main>
    <HeroCenter />
    <Carousel title="Рекомендации для вас" />
    <Carousel title="Категории" />
    <!-- <Carousel title="Топ - 10" /> -->
  </v-main>
</template>

<script>
import HeroCenter from '~/components/Homepage/HeroCenter.vue'
import Carousel from '~/components/Carousel.vue'
export default {
  components: { HeroCenter, Carousel },
  mounted() {
    // console.log(this.firstChar('/shit'))
  },
  methods: {
    // firstChar(str) {
    //   return str.substring(0, 2)
    // },
  },
}
</script>

<style></style>
