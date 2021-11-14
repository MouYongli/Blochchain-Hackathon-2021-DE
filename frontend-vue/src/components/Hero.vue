<template>
  <v-col
  class="hero-cont">
    <div
    class="title">
      {{title}}
    </div>
    <div
    class="description">
      {{desc}}
    </div>
  </v-col>
</template>

<script>
export default {
  name: "Hero",
  props:["title","desc"]
}
</script>

<style scoped>
.hero-cont{
  background:linear-gradient(
      rgba(0, 0, 0, 0.7),
      rgba(0, 0, 0, 0.7)
  ), url("../assets/header1.jpg");
  background-size: cover;

  padding: 100px 0px;
  text-align: center;
  color: #fff;
  font-family: 'Inter', sans-serif;
}
.title{
  font-weight: 700;
  font-family: 'Inter', sans-serif;
}
.description{
  font-weight: 500;
  font-family: 'Inter', sans-serif;
}
/*TODO: FONT and styles

 */
</style>