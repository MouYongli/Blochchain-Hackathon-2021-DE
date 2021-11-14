<template>
  <v-col
  class="hero-cont">
    <div class="logo-title">
      AISpeace
    </div>
    <div
    class="titlename">
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
  color: #fff;
  font-family: 'Inter', sans-serif;
}

.logo-title{
  text-align: center;
  font-size: 3em;
}

.titlename{
  text-align: center;
  font-weight: 800;
  font-size: 2em;
  font-family: 'Inter', sans-serif;
}
.description{
  font-weight: 500;
  font-family: 'Inter', sans-serif;
  text-align: center;
  padding: 80px;
  max-width: 1000px;
  margin: auto;
}
/*TODO: FONT and styles

 */
</style>