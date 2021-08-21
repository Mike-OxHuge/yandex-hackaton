<template>
  <v-main>
    <div class="cats">
      <h1>{{ $route.params }}</h1>
      <Item
        :id="randomId"
        width="400"
        title="Урок 1"
        description="Для урока тебе понадобится пенопласт."
      />
    </div>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
  computed: {
    randomId() {
      return Math.floor(Math.random() * 100)
    },
  },
}
</script>

<style></style>
