<template>
  <v-container>
    <h1>Творчество</h1>
    <AdultOrChild @adult="isAdult = true" @child="isAdult = false" />
    <v-row>
      <v-col cols="4">
        <v-row>
          <Activities
            v-if="isAdult"
            :activities="adultActivities"
            @selected="anythingSelected = true"
          />
          <Activities
            v-else
            :activities="childActivities"
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
import AdultOrChild from '~/components/AdultOrChild.vue'
import Activities from '~/components/Activities.vue'
export default {
  components: { AdultOrChild, Activities },
  data() {
    return {
      anythingSelected: false,
      isAdult: true,
    }
  },
  computed: {
    adultActivities() {
      return this.$store.state.art.adultActivities
    },
    childActivities() {
      return this.$store.state.art.childActivities
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
