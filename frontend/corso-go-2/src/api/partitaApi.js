import Api from './api'


export default {

    async tryStart(id) {
        return Api().get(`/partita/start/${id}`);
    },

    async getChiamabili() {
        return Api().get(`/partita/chiamabili`);
    },

    async getFineChiamata() {
        return Api().get(`/partita/finechiamata`);
    },

    async getGiocata() {
        return Api().get(`/partita/giocata`);
    },

    async chiama(id, valore) {
        return Api().get(`/partita/chiama/${id}/carta/${valore}`);
    },

    async chiamaBot(id) {
        return Api().get(`/partita/chiamabot/${id}`);
    },

    async chiamaSeme(seme) {
        return Api().get(`/partita/chiama/seme/${seme}`);
    },

    async giocaCarta(id, carta) {
        return Api().get(`/partita/gioca/${id}/seme/${carta.SemeStr}/valore/${carta.Valore}`);
    },

    async giocaCartaBot(id) {
        return Api().get(`/partita/giocabot/${id}`);
    },

    async finiscila() {
        return Api().get(`/partita/end`);
    },

    async getVittoria() {
        return Api().get(`/partita/mostravittoria`);
    },

}
