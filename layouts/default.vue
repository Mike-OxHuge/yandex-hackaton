<template>
  <v-app>
    <div class="text-center mt-1 menu">
      <v-row justify="space-between" align="center">
        <v-col cols="3">
          <NuxtLink to="/">
            <h1 style="color: black" class="text-uppercase">Яндекс Досуг</h1>
          </NuxtLink>
        </v-col>
        <v-col cols="5">
          <Menu />
        </v-col>
        <v-col cols="1">
          <v-avatar class="my-2">
            <v-img src="https://picsum.photos/200/200?random=1"></v-img>
          </v-avatar>
        </v-col>
      </v-row>
    </div>
    <Nuxt />
    <Footer />
  </v-app>
</template>

<script>
import Menu from '~/components/Menu.vue'
import Footer from '~/components/Footer.vue'

export default {
  components: { Menu, Footer },
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

<style scoped>
.menu {
  background: rgba(0, 0, 0, 0);
}
</style>
