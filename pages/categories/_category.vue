<template>
  <v-main>
    <div class="cats">
      <h1>{{ $route.params }}</h1>
      <v-row>
        <v-col v-for="n in 4" :key="n" cols="4">
          <Item
            :id="n"
            width="auto"
            url="/categories/foo/bar"
            title="Тип категории"
            description="Краткое описание курса."
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
