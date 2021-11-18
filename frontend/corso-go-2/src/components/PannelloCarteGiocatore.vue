<template>
    <v-flex xs8 :style="getStyleGiocatore()">
        <v-btn class="ma-3" v-if="carte.length > 0" color="#718F94" style="color: white; float: right; margin: 5px" @click="toggleCarte">{{getNomePulsanteToggleCarte()}}</v-btn>
        <v-layout style="justify-content: center; margin-left: 4vw;">
            <div v-for="(carta,i) in carte" :key="i">
                <v-tooltip bottom>
                    <template v-slot:activator="{ on, attrs }">
                        <v-img :class="{selezionata: selectedCarta === carta}"
                            :style="getStyleCarta(carta)"
                            contain
                            :src="getSpriteCartaGiocatore(carta)"
                            @click="toggleSelection(carta)"
                            v-bind="attrs"
                            v-on="on"
                        ></v-img>
                    </template>
                    <span>{{getNomeCarta(carta)}}</span>
                </v-tooltip>
            </div>
        </v-layout>
    </v-flex>
</template>

<script>
export default {
    emits:['cartaSelected'],
    props:['carte', 'toccaA', 'chiamante','selectedCarta'],
    data(){
        return {
            carteVisibili: true
        }
    },
    methods:{
        getNomeCarta(carta){
            return this.$store.getters.getNomeCarta(carta);
        },
        getStyleGiocatore() {
            return this.$store.getters.getStyleGiocatore(0, this.toccaA, this.chiamante);
        },
        toggleCarte(){
                this.carteVisibili = !this.carteVisibili
        },
        getNomePulsanteToggleCarte(){
            if(this.carteVisibili){
                return 'Nascondi';
            }
            else{
                return 'Mostra';
            }
        },
        getSpriteCartaGiocatore(carta){
            if(this.carteVisibili){
                return require(`../assets/${this.$store.state.mazzo}/${carta.Valore + carta.SemeStr}.png`);
            }
            else{
                return require(`../assets/${this.$store.state.mazzo}/retro.png`);
            }
        },
        getStyleCarta(carta) {
            let border = "border: ;"
            if (carta.isSelected) {
                border += "border: 3px solid #24ccf2;"
            }
            return "width: 8.5vh; margin-top: 50px; margin-left: 2px; margin-right: 2px; margin-bottom: 5px; border-radius: 3%; box-shadow: 0 10px 10px 0 rgba(0, 0, 0, 0.2), 10px 10px 20px 0 rgba(0, 0, 0, 0.19); cursor: pointer;" + border
        },
        toggleSelection(carta) {
            this.$emit('cartaSelected', carta)
            
        },
    }
}
</script>

<style scoped>

.selezionata {
  animation: slide-scale 0.2s ease-in forwards;
}

@keyframes slide-scale{
  0% {
    transform: translateY(0) scale(1);
  }

  70% {
    transform: translateY(-15px) scale(1.1);
  }

  100% {
    transform: translateY(-25px) scale(1.2);
  }
  
}

</style>