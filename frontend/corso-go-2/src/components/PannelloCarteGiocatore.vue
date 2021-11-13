<template>
    <v-flex xs8 :style="getStyleGiocatore()">
        <v-btn class="ma-3" v-if="carte.length > 0" color="#718F94" style="color: white; float: right; margin: 5px" @click="toggleCarte">{{getNomePulsanteToggleCarte()}}</v-btn>
        <v-layout style="justify-content: center; margin-left: 4vw;">
            <div v-for="(carta,i) in carte" :key="i">
                <v-img :class="{transform: selectedCarta === carta}"
                        :style="getStyleCarta(carta)"
                        contain
                        :src="getSpriteCartaGiocatore(carta)"
                        @click="toggleSelection(carta)"
                ></v-img>
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
                border += "border: 5px solid #24ccf2;"
            }
            return "width: 8vh; margin-top: 30px; margin-left: 2px; margin-right: 2px; margin-bottom: 5px;" + border
        },
        toggleSelection(carta) {
            this.$emit('cartaSelected', carta)
            
        },
    }
}
</script>

<style scoped>

.transform {
  animation: slide-scale 0.2s ease-in forwards;
}

@keyframes slide-scale{
  0% {
    transform: translateY(0) scale(1);
  }

  70% {
    transform: translateY(-15px) scale(1.05);
  }

  100% {
    transform: translateY(-25px) scale(1.1);
  }
  
}

</style>