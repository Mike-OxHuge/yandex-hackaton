<template>
  <v-sheet class="mt-5">
    <div class="my-2">
      <span class="font-weight-bold">{{ title }}</span>
    </div>
    <v-slide-group show-arrows>
      <v-slide-item v-for="n in 25" :key="n">
        <v-card class="mr-2 mb-1" flat>
          <v-img
            src="/1black.png"
            width="200"
            style="border-radius: 2rem"
          ></v-img>
          <v-card-subtitle
            class="white--text ml-1"
            style="position: absolute; top: 20%; z-index: 9"
            >Название категории</v-card-subtitle
          >
          <v-card-text class="pl-0">Описание категории</v-card-text>
        </v-card>
      </v-slide-item>
    </v-slide-group>
  </v-sheet>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      //   required: true,
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
