<template>
  <v-main>
    <SlugContent />
  </v-main>
</template>

<script>
import SlugContent from '~/components/SlugContent.vue'
export default {
  components: {
    SlugContent,
  },
  props: {
    //
  },
  mounted() {
    console.log('slug is mounted')
    console.log('type is: ', this.$route.params.type)
    console.log('slug is: ', this.$route.params.slug)
  },
  methods: {
    //
  },
}
</script>

<style></style>
