<template>
    <v-container fluid grid-list-md>
        <v-layout row >
            <v-flex xs4></v-flex>
            <v-flex xs4 class="mt-12">
                <v-card>
                    <v-card-title>Chi sei?</v-card-title>
                    <v-card-text>
                        <v-select label="Utenti già registrati"
                                  :items="nomiGiocatori"
                                  v-model="selectedGiocatore"
                                  @change="selectGiocatore"
                        />
                        <v-divider/>
                        <div class="mt-3 mb-1">
                            Se non ci sei:
                        </div>
                        <v-text-field label="Nome nuovo giocatore"
                                      v-model="newGiocatore"
                                      outlined>
                        </v-text-field>
                    </v-card-text>
                    <v-card-actions>
                        <v-spacer/>
                        <v-btn @click="addGiocatore"
                        >
                            Aggiungi giocatore
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </v-flex>
            <v-flex xs4></v-flex>
        </v-layout>
    </v-container>
</template>

<script>
    /* eslint-disable */
    import giocatoriApi from "../api/giocatoriApi";

    export default {
        name: "Home",
        data: () => ({
            newGiocatore: '',
            selectedGiocatore: '',
            giocatori: [],
            nomiGiocatori: [],
            ws: {}
        }),
        mounted() {
            giocatoriApi.getGiocatori().then(t => {
                if (t.status === 200) {
                    this.giocatori = t.data
                    for (let i= 0; i<this.giocatori.length; i++) {
                        this.nomiGiocatori.push(this.giocatori[i].Nome)
                    }
                }
            })
        },
        computed: {
        },
        methods: {
            addGiocatore () {
                giocatoriApi.addGiocatore(this.newGiocatore).then(t => {
                    if (t.status === 200) {
                        this.setGiocatore(t.data, this.newGiocatore)
                        this.$router.push("PARTITA")
                    }
                })
            },
            selectGiocatore () {
                giocatoriApi.selectGiocatore(this.selectedGiocatore).then(t => {
                    if (t.status === 200) {
                        this.$store.dispatch('setGiocatore', {id: t.data, nome: this.selectedGiocatore, punti: 0})
                        this.$router.push("PARTITA")
                    }
                })
            },
            setGiocatore (id, nome) {
                this.$store.dispatch('setGiocatore', {id: id, nome: nome, punti: 0})
            }
        }
    }
</script>

<style scoped>

</style>
