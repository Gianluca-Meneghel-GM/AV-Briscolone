<template>
    <v-img
            v-if="haGiocato()"
            style="width: 9vh; margin: 5px"
            contain
            v-bind:src="getSpriteCarta()"
    ></v-img>
</template>

<script>
    export default {
        name: "CartaGiocata",
        props: ['pos', 'mano', 'valChiamato', 'carteQuestaMano'],
        methods:{
            haGiocato() {
                if (this.carteQuestaMano && this.carteQuestaMano[this.getGiocatoreInPos()]) {
                    return true
                }
                return false
            },
            getSpriteCarta() {
                let carta = this.getCartaQuestaMano(this.pos)
                if (carta.Valore === this.valChiamato && this.mano === 0 && this.carteQuestaMano && Object.keys(this.carteQuestaMano).length < 5) {
                    return require(`../assets/${this.$store.state.mazzo}/retro.png`)
                }
                return require(`../assets/${this.$store.state.mazzo}/${carta.Valore + carta.SemeStr}.png`)
            },
            getCartaQuestaMano() {
                let carta = '2S'
                if (this.carteQuestaMano[this.getGiocatoreInPos()]) {
                    carta = this.carteQuestaMano[this.getGiocatoreInPos()]
                }
                return carta
            },
            getGiocatoreInPos() {
                //return ((this.$store.state.giocatore.id + this.pos) % 5)
                return this.$store.getters.getGiocatoreInPos(this.pos);
            },
        }
        
    }
</script>