<template>
  <v-main>
    <div class="cats">
      <h1>{{ $route.params }}</h1>
      <v-row>
        <v-col v-for="n in 2" :key="n" cols="4">
          <Item
            :id="n"
            width="auto"
            url="/categories/foo/bar/buzz"
            title="Урок"
            description="Краткое описание урока."
          />
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default { components: { Item } }
</script>

<style></style>
