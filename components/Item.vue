<template>
  <v-card class="mr-2 mb-1 item-card" flat nuxt :to="link">
    <v-img
      :id="item.id"
      :src="item.image"
      :width="width"
      :height="height"
      style="border-radius: 2rem"
      class="item-tile"
    ></v-img>
    <h4 class="text-center" @mouseover="addClass" @mouseout="removeClass">
      {{ item.title }}
    </h4>
    <v-card-text class="text-center">{{ description }}</v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
    width: {
      type: String,
      default: '200',
    },
    height: {
      type: String,
      default: '200',
    },
    description: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      //
    }
  },
  computed: {
    link() {
      return (this.$route.path + this.item.link).replace('//', '/')
    },
  },
  mounted() {
    //
  },
  methods: {
    addClass() {
      document.getElementById(this.item.id).classList.add('highlighted')
    },
    removeClass() {
      document.getElementById(this.item.id).classList.remove('highlighted')
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
