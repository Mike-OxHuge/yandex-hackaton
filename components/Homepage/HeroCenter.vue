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
        <span class="black--text">Выбрать занятие</span>
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
}
</style>
