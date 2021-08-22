<template>
  <div>
    <v-row v-if="!areSearching" no-gutters justify="center" class="text-center">
      <v-col>
        <NuxtLink to="/favs"><span>Мой досуг</span></NuxtLink>
      </v-col>
      <v-col>
        <NuxtLink to="/categories"><span>Категории</span></NuxtLink>
      </v-col>
      <v-col>
        <span>Комната отдыха</span>
      </v-col>
      <v-col>
        <v-icon v-if="!areSearching" color="black" @click="areSearching = true"
          >mdi-magnify</v-icon
        >
      </v-col>
    </v-row>
    <v-row v-else justify="center">
      <v-col cols="8">
        <v-text-field
          v-model="searchQuery"
          autofocus
          label="Поиск"
          single-line
          outlined
          hint="Жми Enter"
          dense
          hide-details
          @keydown.enter="handleSearch"
        >
        </v-text-field>
      </v-col>
      <v-icon color="black" @click="areSearching = false">mdi-close</v-icon>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      areSearching: false,
      searchQuery: '',
    }
  },
  methods: {
    handleSearch() {
      console.log(this.searchQuery)
    },
  },
}
</script>

<style scoped>
a {
  text-decoration: none;
  padding: 0.5rem;
  color: #292929 !important;
}
/* home route and active route will show in bold as it matches / and /about */
/* a.nuxt-link-active {
  color: #00c58e !important;
} */
/* exact link will show the primary color for only the exact matching link */
a.nuxt-link-exact-active {
  background: #ff0000;
  border-radius: 5px;
  color: white !important;
  font-weight: bold;
}
</style>
