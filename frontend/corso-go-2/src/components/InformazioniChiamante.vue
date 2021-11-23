<template>
    <v-flex xs2>                
        <v-layout v-if="showCartaSocio" style="justify-content: center"><b style="font-size: 1.5vh; font-style: italic; padding-right: 0.5vh;">Carta del socio:</b></v-layout>
        <v-layout v-else style="justify-content: center"><b style="font-size: 1.5vh; font-style: italic; padding-right: 0.5vh;">Carta chiamata:</b></v-layout>
        <v-layout style="justify-content: center;">
            <b style="font-size: 1.5vh; color: #6e0021">{{cartaChiamata}}</b>
        </v-layout>
        <v-layout style="justify-content: center" class="mt-1"><b style="font-size: 1.5vh; font-style: italic; padding-right: 0.5vh;">Chiamante:</b><b style="font-size: 1.5vh; color: #6e0021">{{chiamante}}</b></v-layout>
        <v-layout style="justify-content: center" class="mt-1"><b style="font-size: 1.5vh; font-style: italic; padding-right: 0.5vh;">Punti per vincere:</b><b style="font-size: 1.5vh; color: #6e0021">{{puntiVittoria}}</b></v-layout>
        <v-layout style="justify-content: center" class="mt-2">
<!--             <transition name="fade"> -->
                <base-timer v-if="timerVisible" @timesUp="timesUp"></base-timer>
<!--             </transition> -->
            <transition name="bounce" mode="out-in">
                <b v-if="!timerVisible && desedetEnabled" style="font-size: 4vh; color: #db0000">Desèdet!!!</b>
            </transition>
        </v-layout>
    </v-flex>
</template>

<script>
import BaseTimer from "./BaseTimer.vue"
export default {
    data(){
        return {
            desedetEnabled: false
        };
    },
    components: {
        BaseTimer
    },
    emits: ['timesUp'],
    props:['showCartaSocio','cartaChiamata','puntiVittoria', 'chiamante', 'timerVisible'],
    methods: {
        timesUp(){
            this.$emit('timesUp');
            this.desedetEnabled = true;
            setTimeout(() => this.desedetEnabled = false, 5000)
            
        }
    }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.bounce-enter-active {
  animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>