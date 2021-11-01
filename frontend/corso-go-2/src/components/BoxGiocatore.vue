
<template>
    <v-flex :style="getStyleGiocatore()">
        <v-layout :style="getStyleNomeGiocatore()">
            <h2>{{getNomeGiocatoreSuTavolo()}}</h2>
        </v-layout>
        <v-layout class="ma-3" style="justify-content: center">
            <b style=" font-style: italic; padding-right: 0.5vh;">Carte:</b>
            <b :style="getStyleSelectedColor()">{{getCarteInMano()}}</b>
        </v-layout>
        <v-layout class="ma-3" style="justify-content: center">
            <b style=" font-style: italic; padding-right: 0.5vh;">Prese:</b>
            <b :style="getStyleSelectedColor()">{{getPrese()}}</b>
        </v-layout>
    </v-flex>
</template>

<script>
/*
function hashCode(str) { // java String#hashCode
    var hash = 0;
    for (var i = 0; i < str.length; i++) {
       hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return 900 * hash;
} 

function intToRGB(i){
    var c = (i & 0x00FFFFFF)
        .toString(16)
        .toUpperCase();

    return "#" + "00000".substring(0, 6 - c.length) + c;
}

function getColor(str){
    return intToRGB(hashCode(str));
}
*/

    export default {
        name: "BoxGiocatore",
        props: ['pos', 'toccaA', 'chiamante'],
        methods:{
            getStyleGiocatore() {
                return this.$store.getters.getStyleGiocatore(this.pos, this.toccaA, this.chiamante);
            },
            getStyleSelectedColor(){
                let style = "color: black;"
                if(this.toccaA === this.getGiocatoreInPos()){
                    style = "color: #4454e3"
                }
                return style
            },
            getStyleNomeGiocatore() {
                let style = "justify-content: center;"
                if(this.toccaA === this.getGiocatoreInPos()){
                    style = "justify-content: center;" + this.getStyleSelectedColor();
                }
                return style
            },
            getNomeGiocatoreSuTavolo() {
                let giocatori = this.$store.state.giocatori
                if (giocatori[this.getGiocatoreInPos()]) {
                    return giocatori[this.getGiocatoreInPos()].Nome
                } else {
                    return ''
                }
            },
            getCarteInMano() {
                let carte = 0
                let giocatori = this.$store.state.giocatori
                if (giocatori[this.getGiocatoreInPos()]) {
                    carte = giocatori[this.getGiocatoreInPos()].carteInMano
                } else {
                    return ''
                }
                return carte
            },
            getPrese() {
                return this.$store.getters.getPrese(this.pos);
            },
            getGiocatoreInPos() {
                return this.$store.getters.getGiocatoreInPos(this.pos);
            },
        }
        
    }
</script>