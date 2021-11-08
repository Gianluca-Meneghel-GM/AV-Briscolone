<template>
    <v-flex xs8 style="background-color: #90B494">
        <v-layout style="height: 22vh">
            <v-flex xs6 style="text-align: -webkit-center">
                <carta-giocata :coordGiocatore="coordGiocatore" :pos="3" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></carta-giocata>
            </v-flex>
            <v-flex xs6 style="text-align: -webkit-center">
                <carta-giocata :coordGiocatore="coordGiocatore" :pos="2" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></carta-giocata>
            </v-flex>
        </v-layout>
        <v-layout style="height: 22vh">
            <v-flex xs3>
                <carta-giocata :coordGiocatore="coordGiocatore" :pos="4" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></carta-giocata>
            </v-flex>
            <v-flex xs6 style="text-align: -webkit-center">
                <carta-giocata :coordGiocatore="coordGiocatore" :pos="0" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></carta-giocata>
            </v-flex>
            <v-flex xs3 style="text-align: -webkit-right">
                <carta-giocata :coordGiocatore="coordGiocatore" :pos="1" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></carta-giocata>
            </v-flex>
        </v-layout>
    </v-flex>
</template>

<script>
import CartaGiocata from "./CartaGiocata.vue"
    export default {
        name: "Tavolo",
        components:{
            CartaGiocata
        },
        props: ['mano', 'valChiamato' ,'carteQuestaMano', 'coordGiocatore']
    }
</script>