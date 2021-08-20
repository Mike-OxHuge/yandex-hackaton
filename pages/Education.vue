<template>
  <v-container>
    <h1>Образование</h1>
    <AdultOrChild @adult="isAdult = true" @child="isAdult = false" />
    <v-container>
      <Activities v-if="isAdult" :activities="adultActivities" />
      <Activities v-else :activities="childActivities" />
    </v-container>
  </v-container>
</template>

<script>
import AdultOrChild from '~/components/AdultOrChild.vue'
import Activities from '~/components/Activities.vue'
export default {
  components: { AdultOrChild, Activities },
  data() {
    return {
      isAdult: true,
    }
  },
  computed: {
    adultActivities() {
      return this.$store.state.education.adultActivities
    },
    childActivities() {
      return this.$store.state.education.childActivities
    },
  },
  mounted() {
    // console.log(this.adultActivities)
    // console.log(this.childActivities)
  },
  methods: {
    //
  },
}
</script>

<style></style>
