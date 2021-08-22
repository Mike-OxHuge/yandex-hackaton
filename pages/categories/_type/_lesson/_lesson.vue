<template>
  <v-main>
    <div class="cats">
      <h3>
        {{ types.description }}
      </h3>
      <iframe
        :title="types.description"
        allowfullscreen="true"
        :src="types.link"
        width="100%"
        height="400px"
      ></iframe>
    </div>
  </v-main>
</template>

<script>
export default {
  async asyncData({ $content, route }) {
    const categories = await $content(route.path).fetch()
    const types = await categories[0]
    return {
      types,
    }
  },
  computed: {
    randomId() {
      return Math.floor(Math.random() * 100)
    },
  },
  mounted() {
    //
  },
}
</script>

<style></style>
