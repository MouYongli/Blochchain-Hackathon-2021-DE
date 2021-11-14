<template>
  <div>
    <Hero title="Test title" desc="Test description"/>

    <v-container>
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
