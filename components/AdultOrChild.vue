<template>
  <v-container>
    <v-chip-group active-class="primary--text" mandatory>
      <v-chip
        v-for="option in options"
        :key="option.i"
        @click="(e) => selection(e)"
      >
        {{ option }}
      </v-chip>
    </v-chip-group>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      options: ['Взрослый', 'Ребенок'],
    }
  },
  methods: {
    selection(e) {
      if (e.target.innerText === 'Взрослый') {
        this.$emit('adult')
      } else {
        this.$emit('child')
      }
    },
  },
}
</script>

<style></style>
