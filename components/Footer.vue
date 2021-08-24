<template>
  <v-container>
    <v-row justify="center" dense align="baseline" class="text-center">
      <v-col cols="2"><p>Условия подписки</p></v-col>
      <v-col cols="2"><h2>Trade</h2></v-col>
      <v-col cols="2"><p>Задать вопрос</p></v-col>
    </v-row>
    <v-divider class="mx-15" />
    <v-row justify="center" class="mt-5" no-gutters>
      <v-col v-for="n in 5" :key="n" cols="1" class="d-flex justify-center">
        <div style="border-radius: 50%; border: 2px solid black">
          <v-avatar class="mx-auto" color="button">
            <!-- <v-img src="https://via.placeholder.com/100"></v-img> -->
          </v-avatar>
        </div>
      </v-col>
    </v-row>
    <p class="text-center mt-5">
      2021 © Team Forward for ООО «Яндекс.Медиасервисы» Hackaton
    </p>
  </v-container>
</template>

<script>
export default {}
</script>

<style scoped></style>
