<template>
  <v-main>
    <v-container>
      <v-row>
        <v-col v-for="card in cards" :key="card.id">
          <Item :item="card" width="auto" height="400" />
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
export default {
  components: { Item },
  data() {
    return {
      cards: [
        {
          title: 'Игры',
          link: '/games',
          image: '/chillout/games.png',
          id: '1',
        },
        {
          title: 'Кино',
          link: '/movies',
          image: '/chillout/movies.png',
          id: '2',
        },
        {
          title: 'Голосовой Чат',
          link: '/voice-chat',
          image: '/chillout/voice.png',
          id: '3',
        },
      ],
    }
  },
}
</script>

<style scoped>
/* #chillax {
  height: 90vh;
  background-image: url('/chillout-bg.png');
  background-repeat: no-repeat;
  background-size: cover;
} */
</style>
