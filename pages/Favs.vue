<template>
  <v-main class="mx-2">
    <h1>Продолжить просмотр</h1>
    <v-row>
      <v-col v-for="item in items" :key="item.id" cols="4">
        <Item :item="item" width="auto" height="200" />
      </v-col>
    </v-row>
    <h1>Рекомендации для вас</h1>
    <Carousel :items="itemsTwo" width="200" />
  </v-main>
</template>

<script>
import Item from '~/components/Item.vue'
import Carousel from '~/components/Carousel.vue'
export default {
  components: { Item, Carousel },

  data() {
    return {
      items: [
        {
          title: 'Лепим динозавра',
          link: '/categories/kids-below-15/lepka/dunnosaurus',
          image: '/categories/kids/types/dunno.png',
          description: "Лепим dunnosaurus'a",
          id: '1',
        },
        {
          title: 'Хакатоны',
          link: '/categories/youth/smart-fella/hackatons',
          image: '/categories/youth/types/hacka.png',
          description: 'Ближашайшие хакатоны',
          id: '4',
        },
      ],
      itemsTwo: [
        {
          title: 'Техники тайм-менеджмента',
          link: '/categories/youth/time-management/tm-techs',
          image: '/categories/youth/types/tm-tech.jpg',
          description: 'Успеть всё!',
          id: '1',
        },
        {
          title: 'Бесплатное просвещение',
          link: 'categories/youth/smart-fella/free-education',
          image: '/categories/youth/types/edu.jpg',
          description: 'Бесплатно. Этим всё сказано',
          id: '2',
        },
        {
          title: 'Прошаренный студент',
          link: '/categories/youth/smart-fella',
          image: '/categories/youth/survive.jpg',
          description: 'Выжить любой ценой',
          id: '3',
        },
        {
          title: 'Скидки по студенческому',
          link: '/categories/youth/smart-fella/discounts',
          image: '/categories/youth/types/discounts.jpg',
          description: 'Дешево и сердито',
          id: '4',
        },
        {
          title: 'Таблица Шульцмана',
          link: '/categories/kids-below-15/developing-attention/shulcmann',
          image: '/categories/kids/types/dev-att.png',
          description: 'Будем ну очень внимательными',
          id: '5',
        },
        {
          title: 'Откуда взять деньги?',
          link: '/categories/ever-developing/investments/where-is-my-money-lebowsky',
          image: '/categories/ever-developing/types/where-is-money-3.jpg',
          description: 'Начать можно даже с 10000 рублей',
          id: '6',
        },
        {
          title: 'Доходность и риски инвестиций',
          link: '/categories/ever-developing/finance-control/risks-and-profits',
          image: '/categories/ever-developing/types/invest.jpg',
          description: 'Финансовая грамотность ',
          id: '7',
        },
      ],
    }
  },
  computed: {
    //
  },
  mounted() {
    //
  },
}
</script>

<style></style>
