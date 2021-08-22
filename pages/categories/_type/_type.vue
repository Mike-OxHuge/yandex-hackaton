<template>
  <v-main>
    <div class="cats">
      <h1>{{ types.title }}</h1>
      <v-row>
        <v-col v-for="type in types.lessons" :key="type.id" cols="4">
          <Item :item="type" width="auto" :description="type.description" />
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
  async asyncData({ $content, route }) {
    const categories = await $content(route.path).fetch()
    const types = await categories[0]
    return {
      types,
    }
  },
  mounted() {
    //
  },
}
</script>

<style></style>
