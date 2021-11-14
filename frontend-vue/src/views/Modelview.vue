<template>
  <div class="model-container">
    <v-row style="margin: 0;">
      <v-col class="image-col col-12 col-md-4">
        <v-img  class="model-image" v-bind:src="require('../assets/header1.jpg')"/>
      </v-col>
      <v-col>
        <p class="model-title inter">{{getTitle()}}</p>
        <p class="model-subtitle inter">{{getSubtitle()}}</p>
      </v-col>
      <v-card class="model-btn-cont col-auto col-sm-4" outlined>
        <p class="model-btn-text">Demo Text</p>
        <v-btn class="model-btn" align="center" color="#1455D9" v-on:click="sendDialog=true; sendLoading=false">
          DEMO
        </v-btn>
      </v-card>
    </v-row>

    <div class="model-info-cont inter">
      <v-card v-if="modelInfo">
        <v-card-title style="padding-top: 5px; padding-bottom: 5px">{{modelInfo.title}}</v-card-title>
        <v-divider/>
        <v-card-subtitle class="model-info-desc">{{modelInfo.content}}</v-card-subtitle>
        <v-card-text>{{$route.params.id}}</v-card-text>
      </v-card>
    </div>

    <!-- Sending dialog-->
    <v-dialog v-model="sendDialog">
      <v-card>
        <v-card-title>Upload Dataset</v-card-title>
        <v-divider/>
        <div class="file-upload">
          <v-card-text>Please upload you dataset as a .jpg/.png file</v-card-text>
          <v-file-input outlined label="Dataset" />
          <div class="send-loader">
            <v-progress-circular class="ma-auto" v-if="sendLoading" color="#0439D9" size="100" indeterminate/>
          </div>
        </div>
        <v-card-actions class="send-cont">
          <v-btn class="send-btn" color="#1455D9" v-on:click="sendDataset(1)">Send</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <!-- Image dialog -->
    <v-dialog v-model="imageDialog">
      <v-card>
        <v-card-title>View Result</v-card-title>
        <v-divider/>
        <div class="file-upload">
          <v-card-text>This is the dataset result: </v-card-text>
          <v-img max-width="100%" v-bind:src="image"/>
        </div>
        <v-card-actions class="send-cont">
          <v-btn class="send-btn" color="#1455D9" v-on:click="imageDialog=false">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      sendDialog: false,
      sendLoading:false,

      imageDialog: false,
      image: "",

      modelInfo: {
        title: "Modelinfo Title",
        content: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
      },
    }
  },
  name: "Modelview",
  methods:{
    getTitle(){
      //TODO: get data from backend
      return "Test Title";
    },
    getSubtitle(){
      //TODO: get data from backend
      return "Test Subtitle";
    },

    async sendDataset(file){
      this.image="";
      this.sendLoading=true;
      await this.sleep(2000);
      this.sendDialog=false;

      //TODO: send dataset and display image
      this.image=require("../assets/header1.jpg");
      await this.sleep(1000);
      this.imageDialog=true;
    },



    sleep(milliseconds) {
      return new Promise(resolve => setTimeout(resolve, milliseconds));
    }
  }
}
</script>

<style scoped>
  .send-loader{
    width: 100%;
    justify-content: center;
    justify-items: center;
  }
  .send-cont{
    justify-content: end;
  }
  .send-btn{
    color: #F2F2F2;
  }
  .file-upload{
    padding: 20px;
  }

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
    margin-bottom: 0;
    font-size: 1.5em;
  }
  .model-subtitle{
    font-weight: 450;
    margin-bottom: 0;
    font-size: 1em;
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