import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        giocatore: {},
        giocatori: {},
        mazzo: 'brescia'
    },
    getters: {
        getStyleGiocatore: (_, getters) => (pos, toccaA, chiamante) => {
            let border = "border: 3px solid #718F94;"
            let color = "#DBCFB0"
            let giocatoreInPos = getters.getGiocatoreInPos(pos);
            if (giocatoreInPos === toccaA) {
                color = "#DEAE31"
            }
            if (giocatoreInPos === chiamante) {
                color = "#C37D92"
                if(toccaA === chiamante){
                    border += "border: 10px solid #DEAE31;"
                }
            }
            
            return "background-color: " + color + " ; border-radius: 6px; " + border
        },
        getPrese: (state, getters) => (pos) => {
            let carte = 0
            let giocatori = state.giocatori
            if (giocatori[getters.getGiocatoreInPos(pos)]) {
                carte = giocatori[getters.getGiocatoreInPos(pos)].cartePrese
            }

            if (carte !== undefined) {
                return carte / 5
            }

            return '-'
        },
        getGiocatoreInPos: (state) => (pos) => {
            return ((state.giocatore.id + pos) % 5)
        },
        getValoreCarta: () => (val) =>  {
            switch (val) {
                case 30:
                case 0:
                    return 'niente'
                case 21:
                    return 'Asso';
                case 13:
                    return 'Tre';
                case 10:
                    return 'Re';
                case 9:
                    return 'Cavallo';
                case 8:
                    return 'Fante';
                default:
                    return val;
            }
        },
        getNomeSeme: () => (seme) => {
            switch (seme) {
                case "B":
                    return "Bastoni"
                case "C":
                    return "Coppe"
                case "D":
                    return "Denari"
                case "S":
                    return "Spade"
            }
        },
        getNomeCarta: (_, getters) => (carta) => {
            return getters.getValoreCarta(carta.Valore) + ' di ' + getters.getNomeSeme(carta.SemeStr);
        },
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
