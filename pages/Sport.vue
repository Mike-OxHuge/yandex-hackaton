<template>
  <v-container>
    <h1>Спорт</h1>
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
    }
  },
  computed: {
    activities() {
      return this.$store.state.sport.activities
    },
  },
  mounted() {
    //
  },
  methods: {
    //
  },
}
</script>

<style></style>
