import Api from './api'


export default {

    async getGiocatori() {
        return Api().get(`/giocatori/all`);
    },
    async addGiocatore(nome) {
        return Api().get(`/giocatori/add/${nome}`);
    },
    async selectGiocatore(nome) {
        return Api().get(`/giocatori/select/${nome}`);
    },
    async addBot(nome) {
        return Api().get(`/giocatori/addbot/${nome}`);
    },
}
