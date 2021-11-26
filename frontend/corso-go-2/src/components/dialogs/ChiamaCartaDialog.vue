<template>
    <v-dialog
                v-model="chiamaDialog"
                width="400"
                persistent
        >
            <v-card>
                <v-card-title class="headline lighten-2 pa-3" style="background-color: #DBCFB0">
                    <b>Fase di Chiamata</b>
                    <v-spacer/>
                    <div v-if="getValoreCarta(valChiamato) !== 'niente'" class="ma-3"> 
                        <b style="color: #4454e3">{{getNomeGiocatore(chiamanteProvvisorio)}}</b> ha chiamato: <b style="color: #4454e3">{{getValoreCarta(valChiamato)}}</b>
                    </div>
                </v-card-title>
                <v-card-text style="text-align: center; background-color: #90B494">
                    <div v-if="chiamante === ''">
                        <v-btn class="ma-3" v-for="val in chiamabili" :disabled="!toccaAMe" :key="val" :style="getStyleChiamabili(val)"
                               @click="chiamaCarta(val, puntiChiamati)">
                            {{getValoreCarta(val)}}
                        </v-btn>
                        <div v-if="showChiamaAPuntiBox" style="text-align: -webkit-center">
                            <v-text-field label="chiama a punti" type="number" :min="puntiChiamati"
                                          :readonly="!toccaAMe"
                                          v-model="puntiChiamati" style="width: 90px; text-align-last: center"/>
                            <v-btn v-if="toccaAMe"
                                   @click="mandaMessaggio('chiamaValore', ['2', puntiChiamati.toString()])">Chiama
                            </v-btn>
                        </div>
                    </div>
                    <v-layout v-else-if="showSemi">
                        <v-flex xs6>
                            <v-btn @click="chiamaSeme('B')" class="ma-3">Bastù</v-btn>
                            <v-btn @click="chiamaSeme('D')" class="ma-3">Den Arrow</v-btn>
                        </v-flex>
                        <v-flex xs6>
                            <v-btn @click="chiamaSeme('C')" class="ma-3">Cope</v-btn>
                            <v-btn @click="chiamaSeme('S')" class="ma-3">Spade</v-btn>
                        </v-flex>
                    </v-layout>
                    <div v-else>Chiama {{getNomeGiocatore(chiamante)}}</div>
                </v-card-text>
            </v-card>
        </v-dialog>
</template>

<script>
export default {
    props:['chiamaDialog','valChiamato','chiamanteProvvisorio','chiamante','chiamabili','toccaAMe','puntiChiamati','showChiamaAPuntiBox','showSemi'],
    emits:['cartaChiamata','chiamaSeme'],
    methods:{
        chiamaCarta(val, puntiChiamati){
            this.$emit('cartaChiamata', val, puntiChiamati);
        },
        chiamaSeme(seme){
            this.$emit('chiamaSeme', seme);
        },
        getStyleChiamabili(val){
            let style = "";
            if(val === "a monte" || val === "passo") {
                style = "background-color: #4454e3; color: white";
            }
            return style;
        },
        getNomeGiocatore(id) {
            return this.$store.state.giocatori[id] ? this.$store.state.giocatori[id].Nome : ''
        },
        getValoreCarta(val) {
            return this.$store.getters.getValoreCarta(val);
        }
    }
}
</script>