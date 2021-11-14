<template>
  <v-app>
    <v-app-bar
        app
        color="#0439D9"
        dark
    >
      <div className="d-flex align-center">

        <div class="logo-title">
          AISpace
        </div>

      </div>

      <v-spacer></v-spacer>

      <v-btn
          href="/"
          text
      >
        <v-icon class="vicon">mdi-home</v-icon>
        <span class="mr-2 inter">Home</span>
      </v-btn>
      <v-btn
          href="/models"
          text
      >
        <v-icon class="vicon">mdi-file-document-outline</v-icon>
        <span class="mr-2 inter">Models</span>
      </v-btn>
      <v-btn
          href="/datasets"
          text
      >
        <v-icon class="vicon">mdi-harddisk</v-icon>
        <span class="mr-2 inter">Datasets</span>
      </v-btn>
      <v-btn
          href="/resources"
          text
      >
        <v-icon class="vicon">mdi-monitor-screenshot</v-icon>
        <span class="mr-2 inter">Resources</span>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view class="r-view"/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter&display=swap');

.logo-title{
  font-weight: 800;
  font-size: 1.4em;
}
.inter{
  font-family: 'Inter', sans-serif;
}

body {
  font-family: 'Inter', sans-serif;
}
.v-application{
  background-color: #F2F2F2 !important;
}

.vicon {
  padding-right: 10px;
}
.r-view{
  background-color: #F2F2F2;
  max-width: 100%!important;
}

</style>
