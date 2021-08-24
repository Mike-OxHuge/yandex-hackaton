<template>
  <v-main id="hero" class="d-flex align-center">
    <h2 class="text-center pt-4">Игра “Крокодил”</h2>
    <v-img
      width="900"
      class="ml-4"
      src="/chillout/games/gamefield.jpeg"
    ></v-img>
  </v-main>
</template>

<script>
export default {}
</script>

<style scoped>
#hero {
  height: 90vh;
  background-image: url('/chillout/games/bg.png');
  background-repeat: no-repeat;
  background-size: cover;
}
</style>
