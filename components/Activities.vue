<template>
  <v-row>
    <v-col
      v-for="action in activities"
      :key="action.i"
      cols="12"
      sm="8"
      md="6"
      @click="selection"
    >
      <v-card>
        <v-img width="100" src="https://picsum.photos/150/150?random=1"></v-img>
        <NuxtLink :to="action.url">{{ action.title }}</NuxtLink>
      </v-card>
    </v-col>
  </v-row>
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
