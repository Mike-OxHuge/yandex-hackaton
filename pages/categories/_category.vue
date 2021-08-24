<template>
  <v-main>
    <div class="cats">
      <div class="d-flex">
        <v-icon class="mx-2" large color="header" @click="$router.go(-1)"
          >mdi-arrow-left
        </v-icon>
        <h1>{{ types.title }}</h1>
      </div>
      <v-row>
        <v-col v-for="type in types.types" :key="type.id" cols="4">
          <Item :item="type" width="auto" description="" />
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
    const categories = await $content(route.path.substring(1)).fetch()
    const types = await categories[0]
    return {
      types,
    }
  },
  data() {
    return {
      //
    }
  },
  mounted() {
    //
  },
}
</script>

<style></style>
