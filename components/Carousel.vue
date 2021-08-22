<template>
  <v-sheet class="mt-5">
    <v-container class="my-2">
      <span class="font-weight-bold">{{ title }}</span>
    </v-container>
    <!-- <v-slide-group show-arrows>
      <v-slide-item v-for="n in 25" :key="n">
        <Item :id="n" url="/categories/foo" />
      </v-slide-item>
    </v-slide-group> -->
  </v-sheet>
</template>

<script>
// import Item from '~/components/Item.vue'
export default {
  // components: { Item },
  props: {
    title: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      //
    }
  },
  methods: {
    //
  },
}
</script>

<style scoped></style>
