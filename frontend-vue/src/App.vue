<template>
  <v-app>
    <v-app-bar
        app
        color="#0439D9"
        dark
    >
      <div className="d-flex align-center">
        <v-img
            alt="Vuetify Logo"
            className="shrink mr-2"
            contain
            src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
            transition="scale-transition"
            width="40"
        />

        <v-img
            alt="Vuetify Name"
            className="shrink mt-1 hidden-sm-and-down"
            contain
            min-width="100"
            src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
            width="100"
        />
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
      <router-view/>
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

.inter{
  font-family: 'Inter', sans-serif;
}

body {
  font-family: 'Inter', sans-serif;
}

.vicon {
  padding-right: 10px;
}
</style>
