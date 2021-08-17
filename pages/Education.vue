<template>
  <v-container>
    <h1>Образование</h1>
    <AdultOrChild @adult="isAdult = true" @child="isAdult = false" />
    <v-row>
      <v-col cols="4">
        <v-row v-if="isAdult">
          <v-col
            v-for="action in adultActivities"
            :key="action.i"
            cols="12"
            @click="anythingSelected = true"
          >
            <NuxtLink :to="action.url">{{ action.title }}</NuxtLink>
          </v-col>
        </v-row>
        <v-row v-else>
          <v-col
            v-for="action in childActivities"
            :key="action.i"
            cols="12"
            @click="anythingSelected = true"
          >
            <NuxtLink :to="action.url">{{ action.title }}</NuxtLink>
          </v-col>
        </v-row>
      </v-col>

      <v-col cols="8">
        <NuxtChild v-if="anythingSelected" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import AdultOrChild from '~/components/AdultOrChild.vue'
export default {
  components: { AdultOrChild },
  data() {
    return {
      anythingSelected: false,
      isAdult: true,
      adultActivities: [
        {
          title: 'Дизайн',
          url: '/education/design',
        },
        {
          title: 'Маркетинг',
          url: '/education/marketing',
        },
        {
          title: 'Программирование',
          url: '/education/programming',
        },
        {
          title: 'Финансовая грамотность',
          url: '/education/finance',
        },
      ],
      childActivities: [
        {
          title: 'Робототехника',
          url: '/education/robotics',
        },
        {
          title: 'Программирование',
          url: '/education/programming',
        },
        {
          title: 'Проектная деятельность',
          url: '/education/project',
        },
        {
          title: 'Животный мир',
          url: '/education/animal',
        },
        {
          title: 'История искусств',
          url: '/education/art',
        },
      ],
    }
  },
  methods: {
    //
  },
}
</script>

<style></style>
