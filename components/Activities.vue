<template>
  <v-container>
    <v-col
      v-for="action in activities"
      :key="action.i"
      cols="12"
      @click="selection"
    >
      <NuxtLink :to="action.url">{{ action.title }}</NuxtLink>
    </v-col>
  </v-container>
</template>

<script>
export default {
  props: {
    activities: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      //
    }
  },
  methods: {
    selection() {
      this.$emit('selected')
    },
  },
}
</script>

<style></style>
