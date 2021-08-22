<template>
  <v-main>
    <div class="cats">
      <h1>Категории</h1>
      <v-row>
        <v-col v-for="cat in cats" :key="cat.id" cols="4">
          <Item :item="cat" width="auto" description="" />
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
  async asyncData({ $content }) {
    const categories = await $content('categories').fetch()
    const cats = categories[0].categories
    return {
      cats,
    }
  },
  computed: {
    //
  },
  mounted() {
    //
  },
}
</script>

<style scoped></style>
