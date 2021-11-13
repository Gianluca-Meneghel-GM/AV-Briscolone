<template>
    <v-flex xs2>                
        <v-layout v-if="showCartaSocio" style="justify-content: center"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Carta del socio:</b></v-layout>
        <v-layout v-else style="justify-content: center"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Carta chiamata:</b></v-layout>
        <v-layout style="justify-content: center;">
            <b style="font-size: 2vh; color: #6e0021">{{cartaChiamata}}</b>
        </v-layout>
        <v-layout style="justify-content: center" class="mt-4"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Chiamante:</b><b style="font-size: 2vh; color: #6e0021">{{chiamante}}</b></v-layout>
        <v-layout style="justify-content: center" class="mt-4"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Punti per vincere:</b><b style="font-size: 2vh; color: #6e0021">{{puntiVittoria}}</b></v-layout>
    </v-flex>
</template>

<script>
export default {
    props:['showCartaSocio','cartaChiamata','puntiVittoria', 'chiamante']
}
</script>