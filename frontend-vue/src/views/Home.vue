<template>
  <div>
    <Hero title="AI Marketplace" desc="Combining the blockchain (Smart contract, NFT, IPFS), privacy-preserving computation and semantic web technologies, together with AI can take data utilization to a new level, building a decentralized, collaborative, trustworthy and autonomous AI asset marketplace."/>

    <v-container class="category-cont" id="models">
      <div class="category-title">
        Available models
      </div>
      <v-row>
        <CardElement title="Automatic Speech Recognition Model" desc="This model is an automatic speech recognition system implemented in PyTorch, supporting both English and German languages." image="asr.png" url="https://google.com"/>
        <CardElement title="Machine Translation Model" desc="TThis model is a Transformer-based machine translation model implemented in TensorFlow." image="machine_translation.png" url="https://google.com"/>
        <CardElement title="Image Classification Model" desc="This model is a neural network implemented in PyTorch for common animal recognition." image="img_classification.png" url="https://google.com"/>
        <CardElement title="Time Series Prediction Model" desc="This service uses Prophet and Statsmodel to forecast and analyze points of a given time series." image="time_series_prediction.png" url="https://google.com"/>
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
