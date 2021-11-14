<template>
  <div>
    <Hero title="Test title" desc="Test description"/>

    <v-container class="category-cont" id="models">
      <div class="category-title">
        Available models
      </div>
      <v-row>
        <CardElement title="Test title" desc="Test description" image="header.png" url="https://google.com"/>
        <CardElement title="Test title" desc="Test description" image="header.png" url="https://google.com"/>
        <CardElement title="Test title" desc="Test description" image="header.png" url="https://google.com"/>
        <CardElement title="Test title" desc="Test description" image="header.png" url="https://google.com"/>
        <!-- Resize stuff-->
      </v-row>

    </v-container>
  </div>
</template>

<script>


  import Hero from "../components/Hero";
  import CardElement from "../components/CardElement";
  export default {
    name: 'Home',

    components: {
      CardElement,
      Hero

    },
  }
</script>

<style scoped>
  .category-cont{
    padding: 20px 30px;
  }
  .category-title{
    margin-top: 30px;
    margin-bottom: 20px;
    font-weight: 800;
    font-size: 1.5em;
  }
</style>
