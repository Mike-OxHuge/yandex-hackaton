<template>
  <v-main>
    <div class="cats">
      <h1>Категории</h1>
      <v-row>
        <v-col v-for="n in 6" :key="n" cols="4">
          <Item :id="n" width="auto" description="" url="/categories/foo" />
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
}
</script>

<style scoped></style>
