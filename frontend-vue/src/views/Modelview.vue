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
        <p class="model-btn-text">Purchase Text</p>
        <v-btn class="model-btn" align="center" color="#1455D9" v-on:click=" purchased=true; notifyText='You have successfully purchased the model!'; notify=true; sendLoading=false">
          PURCHASE
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

    <div id="use-model" class="model-info-cont inter" v-if="purchased">
      <v-card v-if="modelInfo">
        <v-card-title style="padding-top: 5px; padding-bottom: 5px">How to use this model</v-card-title>
        <v-divider/>
        <v-card-subtitle class="model-info-desc">This is your API-Token:</v-card-subtitle>
        <v-col class="token-cont">
          <v-text-field label="Token" outlined v-bind:value="token" disabled></v-text-field>
          <v-btn class="model-btn" align="center" color="#1455D9" v-on:click="copyTokenCliboard(token)">
            Copy
          </v-btn>
        </v-col>

        <div class="test-model-cont">
          <v-card class="model-btn-cont" outlined>
            <v-card-title class="text-subtitle-1">Test model</v-card-title>
            <v-divider></v-divider>
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
        </div>

      </v-card>
    </div>

    <v-card class="model-btn-cont col-auto col-sm-4" outlined v-if="purchased">
      <p class="model-btn-text">Thank you for purchasing this model!</p>
      <v-btn class="model-btn" align="center" color="#1455D9" href="#use-model">
        Use Model
      </v-btn>
    </v-card>

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


    <!-- copy notifier -->
    <v-snackbar
        v-model="notify"
    >
      {{notifyText}}
      <template v-slot:action="{ attrs }">
        <v-btn
            color="#1455D9"
            text
            v-bind="attrs"
            @click="notify = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>

  </div>
</template>

<script>
export default {
  data() {
    return {
      sendLoading:false,

      token: "de123a416b6889319a05758a3013e072b6883027042489ae5729fcfcd3c8ffbc",
      notify: false,
      notifyText:"",
      purchased: false,


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

      //TODO: send dataset and display image
      this.image=require("../assets/header1.jpg");
      this.sendLoading=false;
      await this.sleep(1000);
      this.imageDialog=true;
    },

    copyTokenCliboard(token){
      const clipboardData =
          event.clipboardData ||
          window.clipboardData ||
          event.originalEvent?.clipboardData ||
          navigator.clipboard;

      clipboardData.writeText(token);
      this.notifyText="The Token was copied to your clipboard!";
      this.notify=true;
    },



    sleep(milliseconds) {
      return new Promise(resolve => setTimeout(resolve, milliseconds));
    }
  }
}
</script>

<style scoped>

  .test-model-cont{
    padding: 30px;
  }

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