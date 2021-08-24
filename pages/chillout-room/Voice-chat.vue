<template>
  <v-main id="hero">
    <v-container v-for="man in people" :key="man.i" fluid>
      <div :style="man.style" class="d-flex flex-column">
        <span class="text-center my-1 px-1 name">
          {{ man.name }}
        </span>
        <v-avatar size="100">
          <v-img :src="man.pic"></v-img>
        </v-avatar>
      </div>
    </v-container>
  </v-main>
</template>

<script>
export default {
  data() {
    return {
      people: [
        {
          name: 'Boris',
          pic: 'https://picsum.photos/200/200?random=1',
          style: {
            position: 'absolute',
            top: '20%',
            left: '20%',
          },
        },
        {
          name: 'Doris',
          pic: 'https://picsum.photos/200/200?random=2',
          style: {
            position: 'absolute',
            top: '150px',
            left: '150px',
          },
        },
        {
          name: 'Britney',
          pic: 'https://picsum.photos/200/200?random=3',
          style: {
            position: 'absolute',
            top: '50%',
            left: '70%',
          },
        },
      ],
    }
  },
}
</script>

<style scoped>
#hero {
  height: 90vh;
  background-image: url('/chillout/voice-bg.png');
  background-repeat: no-repeat;
  background-size: cover;
  position: relative;
}

.name {
  font-family: Roboto;
  font-style: normal;
  font-weight: bolder;
  font-size: 24px;
  line-height: 28px;
  color: #ffffff;
  -webkit-text-stroke: 1px #ff0000;
}
</style>
