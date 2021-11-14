<template>
  <v-col class="col-12 col-sm-6 col-md-4">
    <v-card>
      <v-img
          max-height="150px"
          width="auto"
          v-bind:src="require('../assets/'+image)"/>
      <v-card-title class="card-title">{{title}}</v-card-title>
      <v-card-text>{{desc}}</v-card-text>
      <v-card-actions class="card-actions">
        <v-btn
            color="#1455D9"
            class="card-btn"
            align="right"
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
.card-title{
  padding-bottom: 0;
}
.card-actions{
  justify-content: end;
}
.card-btn{
  color: #F2F2F2;
}
</style>