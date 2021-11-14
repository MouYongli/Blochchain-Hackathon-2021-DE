<template>
  <div class="model-container">
    <v-row style="margin: 0;">
      <v-col class="image-col col-12 col-md-4">
        <v-img  class="model-image" v-bind:src="require('../assets/header1.jpg')"/>
      </v-col>
      <v-col>
          <p class="model-title inter">{{getTitle()}}</p>
      </v-col>
      <v-card class="model-btn-cont col-auto col-sm-4" outlined>
        <p class="model-btn-text">Demo Text</p>
        <v-btn class="model-btn" align="center" color="#1455D9">
          DEMO
        </v-btn>
      </v-card>
    </v-row>

    <div class="model-info-cont">
      <v-card>
        <v-card-title style="padding-top: 5px; padding-bottom: 5px">Title</v-card-title>
        <v-divider/>
        <v-card-subtitle class="model-info-desc">Content</v-card-subtitle>
        <v-card-text>{{$route.params.id}}</v-card-text>
      </v-card>
    </div>

  </div>
</template>

<script>
export default {
  name: "Modelview",
  methods:{
    getTitle(){
      //TODO: get data from backend
      return "Test Title";
    }
  }
}
</script>

<style scoped>
  .image-col{
    padding-left: 0;
    padding-right: 0;
  }
  .model-container{
    margin: 40px;
  }
  .model-image{
    width: auto;
    max-height: 100px;
    border-radius: 5px;
  }
  .model-title{
    font-weight: 900;
    font-size: 1.5em;
  }
  .model-btn-cont{
    width: 100%;
    padding: 20px;
    align-content: center;
    text-align: center;
    border-color: #0439D9;
    background-color: #F2F2F2;
  }
  .model-btn{
    color: #fff;
  }
  .model-btn-text{
    font-weight: 300;
    font-size: 1.2em;
  }

  /*
  model info
   */

  .model-info-cont{
    padding: 20px 0px;

  }
  .model-info-desc{
    font-weight: 400;
    font-size: 1em;
  }

</style>