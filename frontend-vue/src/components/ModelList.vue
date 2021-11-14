<template>
  <v-container class="category-cont" id="models">
    <div class="category-title">
      Available models
    </div>
    <v-row>
      <CardElement title="Automatic Speech Recognition Model" desc="This model is an automatic speech recognition system implemented in PyTorch, supporting both English and German languages." image="asr.png" url="/models/language"/>
      <CardElement title="Machine Translation Model" desc="TThis model is a Transformer-based machine translation model implemented in TensorFlow." image="machine_translation.png" url="/models/translation"/>
      <CardElement title="Image Classification Model" desc="This model is a neural network implemented in PyTorch for common animal recognition." image="img_classification.png" url="/models/classification"/>
      <CardElement title="Time Series Prediction Model" desc="This service uses Prophet and Statsmodel to forecast and analyze points of a given time series." image="time_series_prediction.png" url="/models/time-series"/>
      <CardElement title="Fast RGB to Hyperspectral image Conversion" desc="The goal of this work is providing a cost-efficient solution for hyperspectral imaging that can reconstruct the spectra of a natural scene from a single RGB image captured by a camera with known spectral response." image="hypercube.jpg" url="/models/graphics"/>
      <!-- Resize stuff-->
    </v-row>

  </v-container>
</template>

<script>
import CardElement from "./CardElement";
export default {
  name: "ModelList",
  components: {CardElement}
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