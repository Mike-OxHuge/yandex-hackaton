<template>
  <v-container>
    <v-row justify="center" dense align="baseline" class="text-center">
      <v-col cols="2"><p>Условия подписки</p></v-col>
      <v-col cols="2"><h2>Trade</h2></v-col>
      <v-col cols="2"><p>Задать вопрос</p></v-col>
    </v-row>
    <v-divider class="mx-15" />
    <v-row justify="center" class="mt-5" no-gutters>
      <v-col v-for="n in 5" :key="n" cols="1" class="d-flex">
        <v-avatar class="mx-auto">
          <v-img src="https://via.placeholder.com/100"></v-img>
        </v-avatar>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {}
</script>

<style></style>
