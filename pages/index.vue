<template>
  <v-main>
    <HeroCenter />
    <Carousel title="Рекомендации для вас" :items="itemsTwo" width="200" />
    <v-container>
      <Carousel title="Категории" :items="itemsOne" width="200" />
    </v-container>
    <!-- <Carousel title="Топ - 10" /> -->
  </v-main>
</template>

<script>
import HeroCenter from '~/components/Homepage/HeroCenter.vue'
import Carousel from '~/components/Carousel.vue'
export default {
  components: { HeroCenter, Carousel },
  async asyncData({ $content, route }) {
    const categories = await $content('categories').fetch()
    const cats = categories[0].categories
    return {
      cats,
    }
  },
  data() {
    return {
      itemsOne: [],
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
  mounted() {
    this.repopulate()
  },
  methods: {
    repopulate() {
      if (this.cats) {
        this.itemsOne = this.cats
        for (let i = 0; i < this.itemsOne.length; i++) {
          this.itemsOne[i].link = `/categories${this.itemsOne[i].link}`
        }
      }
    },
  },
}
</script>

<style></style>
