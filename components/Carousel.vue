<template>
  <v-sheet class="mt-5">
    <v-container class="my-2">
      <h2 class="font-weight-bold">{{ title }}</h2>
    </v-container>
    <v-slide-group show-arrows>
      <v-slide-item v-for="item in items" :key="item.id">
        <Item :item="item" :width="width" height="200" />
      </v-slide-item>
    </v-slide-group>
  </v-sheet>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
  props: {
    items: {
      type: Array,
      required: true,
    },
    title: {
      type: String,
      default: '',
    },
    width: {
      type: String,
      default: 'auto',
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
