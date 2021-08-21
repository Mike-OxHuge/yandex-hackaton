<template>
  <v-card class="mr-2 mb-1 item-card" flat nuxt :to="url">
    <v-img
      :id="id"
      :src="image"
      :width="width"
      style="border-radius: 2rem"
      class="item-tile"
    ></v-img>
    <v-card-subtitle
      class="white--text ml-1"
      style="position: absolute; top: 20%"
      @mouseover="addClass"
      @mouseout="removeClass"
      >{{ title }}</v-card-subtitle
    >
    <v-card-text class="pl-0">{{ description }}</v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    id: {
      type: Number,
      required: true,
    },
    width: {
      type: String,
      default: '200',
    },
    title: {
      type: String,
      default: 'Название категории',
    },
    description: {
      type: String,
      default: 'Описание категории',
    },
    image: {
      type: String,
      default: '/1black.png',
    },
    url: {
      type: String,
      default: '#',
      // required: true,
    },
  },
  data() {
    return {
      //
    }
  },
  computed: {
    //
  },
  mounted() {
    //
  },
  methods: {
    addClass() {
      document.getElementById(this.id).classList.add('highlighted')
    },
    removeClass() {
      document.getElementById(this.id).classList.remove('highlighted')
    },
  },
}
</script>

<style scoped>
.item-tile:hover {
  border: 3px #ffdf3a solid;
}
.item-tile {
  border: 3px rgba(0, 0, 0, 0) solid;
}
.highlighted {
  border: 3px #ffdf3a solid;
}

.item-card {
  background: #ffff0000 !important;
}
</style>
