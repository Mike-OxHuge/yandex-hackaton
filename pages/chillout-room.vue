<template>
  <v-container>
    <h1>Комната отдыха</h1>
    <v-row>
      <v-col cols="4">
        <v-row>
          <Activities
            :activities="activities"
            @selected="anythingSelected = true"
          />
        </v-row>
      </v-col>

      <v-col cols="8">
        <NuxtChild v-if="anythingSelected" />
         <p v-else>Выбери что-нибудь слева. </p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Activities from '~/components/Activities.vue'
export default {
  components: { Activities },
  data() {
    return {
      anythingSelected: false,
      isAdult: true,
    }
  },
  computed: {
    activities() {
      return this.$store.state.chillOut.activities
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
