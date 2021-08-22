<template>
  <v-main>
    <h1>Избранное...</h1>
    <!-- <h1>{{ categories.title }}</h1>
    <h2>{{ comps.title }}</h2> -->
  </v-main>
</template>

<script>
export default {
  async asyncData({ $content }) {
    const categories = await $content('categories').fetch()

    return {
      categories,
    }
  },
  data() {
    return {
      //
    }
  },
  computed: {
    //
  },
  mounted() {
    //
  },
}
</script>

<style></style>
