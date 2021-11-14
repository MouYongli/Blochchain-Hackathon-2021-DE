<template>
  <div class="model-container" v-if="loaded">
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
      <v-card v-if="model.addr">
        <v-card-title style="padding-top: 5px; padding-bottom: 5px">{{getTitle()}}</v-card-title>
        <v-divider/>
        <v-card-subtitle class="model-info-desc">{{getDescription()}}</v-card-subtitle>
        <v-card-text>Address: {{$route.params.id}}</v-card-text>
      </v-card>
    </div>

    <div id="dl-data" class="model-info-cont inter" v-if="purchased">
      <v-card v-if="model.addr">
        <v-card-title style="padding-top: 5px; padding-bottom: 5px">How to use this dataset</v-card-title>
        <v-divider/>
        <v-card-subtitle class="model-info-desc">Download this dataset:</v-card-subtitle>
        <v-col class="token-cont">
          <v-btn class="model-btn" align="center" color="#1455D9">
            Download
          </v-btn>
        </v-col>

        <div class="test-model-cont">
          <v-card class="model-btn-cont" outlined>
            <v-card-title class="text-subtitle-1">View Dataset</v-card-title>
            <v-divider></v-divider>
            <div class="file-upload">
              <div class="send-loader">
                <v-progress-circular class="ma-auto" v-if="sendLoading" color="#0439D9" size="100" indeterminate/>
              </div>
            </div>
            <v-card-actions class="send-cont">
              <v-btn class="send-btn" color="#1455D9" v-on:click="sendDataset(1)">View</v-btn>
            </v-card-actions>

          </v-card>
        </div>

      </v-card>
    </div>

    <v-card class="model-btn-cont col-auto col-sm-4" outlined v-if="purchased">
      <p class="model-btn-text">Thank you for purchasing this dataset!</p>
      <v-btn class="model-btn" align="center" color="#1455D9" href="#dl-data">
        Download
      </v-btn>
    </v-card>

    <!-- Image dialog -->
    <v-dialog v-model="imageDialog">
      <v-card>
        <v-card-title>Dataset</v-card-title>
        <v-divider/>
        <div class="file-upload">
          <v-card-text>This is the dataset information: </v-card-text>
          <v-img max-width="100%" v-bind:src="image"/>
        </div>
        <v-col class="token-cont">
          <v-text-field label="Raw Data" outlined v-bind:value="image" disabled></v-text-field>
        </v-col>
        <v-card-actions class="send-cont">
          <v-btn class="send-btn" color="#1455D9" v-on:click="imageDialog=false">Close</v-btn>
          <v-btn class="model-btn" align="center" color="#1455D9" v-on:click="copyTokenCliboard(image)">
            Copy to clipboard
          </v-btn>
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
import axios from "axios";

export default {
  data() {
    return {
      model:{},
      loaded: false,

      sendLoading:false,

      token: "de123a416b6889319a05758a3013e072b6883027042489ae5729fcfcd3c8ffbc",
      notify: false,
      notifyText:"",
      purchased: false,


      imageDialog: false,
      image: "",
    }
  },
  created() {
    this.getModel(this.$route.params.id);
  },
  name: "Modelview",
  methods:{
    async getModel(addr){
      try {
        let res = await axios.get("http://3.69.255.140:8090/ai-marketplace/data");
        if (res.data !== undefined) {
          let modelList = res.data;
          for(let model of modelList){
            if(model.addr == addr){
              this.model=model;
              this.loaded=true;
              console.log("Modellist: ",res)
            }
          }

        }
      }catch (e){
        console.error(e);
      }
    },

    getTitle(){
      if(this.model.name){
        return this.model.name;
      }
    },
    getSubtitle(){
      return "Model"
    },
    getDescription(){
      if(this.model.info){
        return this.model.info;
      }
    },

    async sendDataset(file){
      this.image="";
      this.sendLoading=true;
      await this.sleep(2000);

      //TODO: send dataset and display image
      let res=await axios.get("http://3.69.255.140:8090/ai-marketplace/contract/"+this.model.addr);
      if(res.data!==undefined){
        this.image=res.data;
      }
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