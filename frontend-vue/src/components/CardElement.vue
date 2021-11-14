<template>
  <v-col
  cols="4">
    <v-card>
      <v-img
          max-height="200px"
          width="auto"
          v-bind:src="require('../assets/'+image)"/>
      <v-card-title>{{title}}</v-card-title>
      <v-card-text>{{desc}}</v-card-text>
      <v-card-actions>
        <v-btn
            target="_blank"
        v-bind:href="url">
           Open Link
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-col>
</template>

<script>
export default {
  name: "CardElement",
  props:["title","image","desc","url"]
}
</script>

<style scoped>

</style>