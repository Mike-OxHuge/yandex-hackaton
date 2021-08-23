<template>
  <v-container
    id="hero-center"
    fluid
    d-flex
    flex-column
    justify-center
    align-center
  >
    <h1>Досуг</h1>
    <div style="width: 30%" class="text-center my-5">
      <p>
        Это место, где вы можете проводить время с близкими, независимо в какой
        части планеты каждый из вас находится
      </p>
    </div>
    <div>
      <v-btn color="button">
        <NuxtLink to="/categories"
          ><span class="black--text">Выбрать занятие</span></NuxtLink
        >
      </v-btn>
    </div>
  </v-container>
</template>

<script>
export default {}
</script>

<style scoped>
#hero-center {
  height: 90vh;
  background-image: url('/bg.jpeg');
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
}
#hero-center::before {
  content: '';
  margin-top: -90px;
}
</style>
