<template>
    <transition name="fade"
                @before-leave="beforeLeave"
                @leave="leave"
                @after-leave="afterLeave"
    >
        <v-img
                v-if="haGiocato()"
                style="width: 9vh; margin: 5px"
                contain
                v-bind:src="getSpriteCarta()"
        ></v-img>
    </transition>
</template>

<script>
import Velocity from 'velocity-animate'
    export default {
        name: "CartaGiocata",
        props: ['pos', 'mano', 'valChiamato', 'carteQuestaMano', 'coordGiocatore'],
        methods:{
            haGiocato() {
                if (this.carteQuestaMano && this.carteQuestaMano[this.getGiocatoreInPos()]) {
                    return true
                }
                return false
            },
            getSpriteCarta() {
                let carta = this.getCartaQuestaMano(this.pos);
                if (this.isCartaCoperta()) {
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
                return this.$store.getters.getGiocatoreInPos(this.pos);
            },
            isCartaCoperta(){
                let carta = this.getCartaQuestaMano(this.pos)
                return carta.Valore === this.valChiamato && this.mano === 0 && this.carteQuestaMano && Object.keys(this.carteQuestaMano).length < 5;
            },
            beforeLeave(el){
              el.style.opacity = 1;
              el.style.scale = 1
            },
            leave(el, done){
              try{
                const elSize = el.getBoundingClientRect();
                const centerX =  el.offsetLeft + elSize.width / 2;
                const centerY =  el.offsetTop + elSize.height / 2;

                const gX = this.coordGiocatore.offsetLeft + this.coordGiocatore.width / 2;
                const gY = this.coordGiocatore.offsetTop + this.coordGiocatore.height / 2;

                let translateX = (gX - centerX)+"px";
                let translateY = (gY - centerY)+"px";

/*
                console.log("*********************************************************************")
                console.log("pos carta "+ this.pos +": x => " + centerX +" y => "+ centerY);
                console.log("pos giocatore "+ this.coordGiocatore.pos +" x => " + gX +" y => "+ gY);
                console.log("translate: x => " + (translateX) +" y => "+ (translateY));
*/

                Velocity(el,
                {
                  //rotateZ: '360deg',
                  translateY: translateY,
                  translateX: translateX,
                  opacity: 0
                },
                { 
                  duration: 1500,
                  complete: done 
                })
              }
              catch(e){
                //console.log(e)
                let round = 1;
                const interval = setInterval(function(){
                  el.style.opacity = round - round * 0.01;
                  round++;
                  if(round > 100){
                    clearInterval(interval);
                    done();
                  }
                }, 20);
                done();
              }

            },
            afterLeave(el){
              el.style.opacity = 0;
            }
        }
        
    }
</script>

<style scoped>


.fade-enter-active {
  transition: opacity 0.75s;
}
.fade-enter {
  opacity: 0;
}

</style>