import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        giocatore: {},
        giocatori: {},
        mazzo: 'brescia'
    },
    mutations: {
        SetGiocatore(state, giocatore) {
            state.giocatore = giocatore;
        },
        SetGiocatori(state, giocatori) {
            for (let i = 0; i < giocatori.length; i++) {
                state.giocatori[giocatori[i].Id] = giocatori[i]
            }
        },
        SetCarteGiocatori(state, carte){
            for (let i = 0; i < carte.inMano.length; i++){
                state.giocatori[i].carteInMano = carte.inMano[i]
                state.giocatori[i].cartePrese = carte.prese[i]
            }
        },
        SetMazzo(state, mazzo){
            state.mazzo = mazzo
        }
    },
    actions: {
        setGiocatore(context, giocatore) {
            context.commit('SetGiocatore', giocatore);
        },
        setGiocatori(context, giocatori) {
            context.commit('SetGiocatori', giocatori);
        },
        setCarteGiocatori(context, carte) {
            context.commit('SetCarteGiocatori', carte);
        },
        setMazzo(context, mazzo) {
            context.commit('SetMazzo', mazzo)
        }
    }
})
