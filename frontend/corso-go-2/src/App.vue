<template>
  <v-app id="inspire" style="overflow: hidden">
    <v-app-bar
            app
            clipped-right
            color="#718F94"
    >
      <v-toolbar-title><div style="color: #ffffff">{{header}}</div></v-toolbar-title>
      <v-spacer></v-spacer>
      <div class="text-center">
        <v-menu offset-y class="ma-3">
          <template v-slot:activator="{ on, attrs }">
            <v-btn
                    v-bind="attrs"
                    v-on="on"
            >
              Cambia Mazzo
            </v-btn>
          </template>
          <v-list>
            <v-list-item
                    v-for="(mazzo, index) in mazzi"
                    :key="index"
                    @click="setMazzo(mazzo)"
            >
              <v-list-item-title>{{ mazzo }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
      <v-btn v-if="showFiniscila()" @click="finisciPartita">Finiscila</v-btn>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>

  </v-app>
</template>

<script>
  import partitaApi from "./api/partitaApi";

  export default {
    props: {
      source: String,
    },
    data: () => ({
      selectedMazzo: 'brescia',
      //'Bergamo', 'Milano' temporaneamente rimosse in attesa di nuove sprite
      mazzi: ['Brescia', 'Napoli', 'Jacovitti']
    }),
    computed: {
      header () {
        return this.$store.state.giocatore.nome ? this.$store.state.giocatore.nome : 'Salve'
      }
    },
    methods: {
      finisciPartita(){
        partitaApi.finiscila()
        /*.then(t => {
          if (t.status === 200) {
          }
        })
        */
      },
      setMazzo(mazzo){
        this.$store.dispatch('setMazzo', mazzo.toLowerCase())
      },
      showFiniscila(){
        //console.log("showFiniscila", this.$store.state.giocatore)
        return this.$store.state.giocatore.id == undefined || this.$store.state.giocatore.id === 0;
      }
    }
  }
</script>