<template>
  <v-container>
    <p v-if="$route.params.slug">Здесь будет {{ $route.params.slug }}</p>
  </v-container>
</template>

<script>
export default {
  methods: {
    //
  },
}
</script>

<style></style>
