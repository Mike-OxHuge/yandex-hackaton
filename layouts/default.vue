<template>
  <v-app dark>
    <v-container class="text-center">
      <NuxtLink to="/"> <h1>Yandex.dosug</h1> </NuxtLink>
      <v-row>
        <v-col v-for="link in links" :key="link.i" sm="6" md="3">
          <NuxtLink :to="link.url">{{ link.title }}</NuxtLink>
        </v-col>
      </v-row>
    </v-container>
    <Nuxt />
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      links: [
        {
          title: 'Образование',
          url: '/education',
        },
        {
          title: 'Творчество',
          url: '/art',
        },
        {
          title: 'Спорт',
          url: '/sport',
        },
        {
          title: 'Комната отдыха',
          url: '/chillout-room',
        },
      ],
    }
  },
}
</script>

<style>
a {
  text-decoration: none;
}
/* home route and active route will show in bold as it matches / and /about */
a.nuxt-link-active {
  color: #00c58e !important;
}
/* exact link will show the primary color for only the exact matching link */
a.nuxt-link-exact-active {
  color: #00c58e !important;
}
</style>
