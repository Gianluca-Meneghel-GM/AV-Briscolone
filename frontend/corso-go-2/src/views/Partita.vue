<template>
    <div class="text-center ma-n3 pa-0" style="background-color: #BFC8AD">
        <v-dialog
                v-model="notYetDialog"
                width="500"
        >
            <v-card style="text-align: center; background-color: #90B494">
                <v-card-title class="headline lighten-2" style="background-color: #DBCFB0">
                    Attendere prego...
                </v-card-title>

                <v-card-text class="mt-4">
                    Mancano ancora {{missingPlayers}} giocatori
                </v-card-text>
                <v-card-actions v-if="me === 0">
                    <v-spacer/>
                    <v-btn @click="mandaMessaggio('addBots')">Gioca coi bot</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-layout style="height:18vh;">
            <box-giocatore :pos="3" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
            <box-giocatore :pos="2" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
        </v-layout>
        <v-layout style="height:45vh">
            <box-giocatore :pos="4" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
            <tavolo :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></tavolo>
            <box-giocatore :pos="1" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
        </v-layout>
        <v-layout style="justify-content: center; height:4.5vh; margin-bottom:5px">
            <div class="ma-3"><b style="font-style: italic; padding-right: 0.5vh;">Prese:</b><b style="color:#4454e3; padding-right: 0.5vh;">{{getPrese(0)}}</b></div>
            <v-btn v-if="sonoProntoBtn" class="ma-3" color="#718F94" style="color: white;" @click="sonoPronto()">Sono
                pronto
            </v-btn>
            <v-btn v-if="sommaPunti < 3" class="ma-3" color="#718F94" style="color: white" @click="giocaCarta">a
                monte
            </v-btn>
            <div v-if="toccaAMe">
                <v-btn class="ma-3" color="#718F94" style="color: white" @click="giocaCarta">Giocala</v-btn>
                <!--v-btn class="ma-3" color="#718F94" style="color: white" @click="tuttoNostro">Tutto nostro</v-btn-->
            </div>
        </v-layout>
        <v-layout style="justify-content: center; height: 24vh">
            <v-flex xs2><h2 style=" font-style: italic; padding-right: 0.5vh;">Tocca a: </h2><h1 style="color: #4454e3">{{getNomeGiocatore(toccaA)}}</h1></v-flex>
            <v-flex xs8 :style="getStyleGiocatore(0)">
                <v-layout style="justify-content: center;">
                    <div v-for="(carta,i) in carte" :key="i">
                        <v-img :class="{transform: selectedCarta === carta}"
                                :style="getStyleCarta(carta)"
                                contain
                                :src="require(`../assets/${$store.state.mazzo}/${carta.Valore + carta.SemeStr}.png`)"
                                @click="toggleSelection(carta)"
                        ></v-img>
                    </div>
                </v-layout>
            </v-flex>
            <v-flex xs2>                
                <v-layout v-if="showCartaSocio" style="justify-content: center; font-style: italic"><b>Carta del socio:</b></v-layout>
                <v-layout v-else style="justify-content: center; font-style: italic"><b>Carta chiamata:</b></v-layout>
                <v-layout style="justify-content: center;">
                    <b style="color: #6e0021">{{getNomeCarta(valChiamato)}} <div v-if="showCartaSocio"> di {{getNomeSeme(briscola)}} </div></b>
                </v-layout>
                <v-layout style="justify-content: center; font-style: italic" class="mt-4"><b>Chiamante:</b></v-layout>
                <v-layout style="justify-content: center;">
                    <h2 style="color: #6e0021">{{getNomeGiocatore(chiamante)}}</h2>
                </v-layout>
                <v-layout style="justify-content: center; font-style: italic" class="mt-4"><b>Punti per vincere:</b></v-layout>
                <v-layout style="justify-content: center;">
                    <h2 style="color: #6e0021">{{puntiVittoria}}</h2>
                </v-layout>
            </v-flex>
        </v-layout>
        <v-dialog
                v-model="chiamaDialog"
                width="400"
                persistent
        >
            <v-card>
                <v-card-title class="headline lighten-2 pa-3" style="background-color: #DBCFB0">
                    Fase di Chiamata
                    <v-spacer/>
                    <div v-if="valChiamato !== ''" class="ma-3"> {{getNomeGiocatore(chiamanteProvvisorio)}} ha chiamato:
                        {{getNomeCarta(valChiamato)}}
                    </div>
                </v-card-title>
                <v-card-text style="text-align: center; background-color: #90B494">
                    <div v-if="chiamante === ''">
                        <v-btn class="ma-3" v-for="val in chiamabili" :disabled="!toccaAMe" :key="val"
                               @click="mandaMessaggio('chiamaValore', [val.toString(), puntiChiamati.toString()])">
                            {{getNomeCarta(val)}}
                        </v-btn>
                        <div v-if="showChiamaAPuntiBox" style="text-align: -webkit-center">
                            <v-text-field label="chiama a punti" type="number" :min="puntiChiamati"
                                          :readonly="!toccaAMe"
                                          v-model="puntiChiamati" style="width: 90px; text-align-last: center"/>
                            <v-btn v-if="toccaAMe"
                                   @click="mandaMessaggio('chiamaValore', ['2', puntiChiamati.toString()])">Chiama
                            </v-btn>
                        </div>
                    </div>
                    <v-layout v-else-if="showSemi">
                        <v-flex xs6>
                            <v-btn @click="chiamaSeme('B')" class="ma-3">Bastù</v-btn>
                            <v-btn @click="chiamaSeme('D')" class="ma-3">Den Arrow</v-btn>
                        </v-flex>
                        <v-flex xs6>
                            <v-btn @click="chiamaSeme('C')" class="ma-3">Cope</v-btn>
                            <v-btn @click="chiamaSeme('S')" class="ma-3">Spade</v-btn>
                        </v-flex>
                    </v-layout>
                    <div v-else>Chiama {{getNomeGiocatore(chiamante)}}</div>
                </v-card-text>
            </v-card>
        </v-dialog>
        <v-dialog
                v-model="vittoriaDialog"
                :width="getLarghezzaDialogVittoria()"
                persistent
        >
            <v-card>
                <v-card-title class="headline lighten-2" style="background-color: #DBCFB0">
                    Finita
                </v-card-title>
                <v-card-text style="text-align: center; background-color: #90B494">
                    <div class="pa-3" style="font-size: 1.3em; font-weight: bold">
                        Han vinto {{getVincitoriStr(datiVittoria.Vincitori)}}{{datiVittoria.PuntiVincitori}} a {{120 -
                        datiVittoria.PuntiVincitori}}
                    </div>
                    <v-simple-table v-if="!showPunteggiTotali" style="background-color: #DBCFB0">
                        <template v-slot:default>
                            <thead>
                            <tr>
                                <th class="text-center" style="font-size: 1.1em">Nome</th>
                                <th class="text-center" style="font-size: 1.1em">Punti</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr v-for="(punti,i) in datiVittoria.PuntiPartita" :key="i">
                                <td>{{ getNomeGiocatore(i) }}</td>
                                <td>{{ punti }}</td>
                            </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                    <v-simple-table v-else style="background-color: #DBCFB0">
                        <template v-slot:default>
                            <thead>
                            <tr>
                                <th class="text-center" style="font-size: 1.1em">Nome</th>
                                <th class="text-center" style="font-size: 1.1em">Giornate</th>
                                <th class="text-center" style="font-size: 1.1em">Vittorie</th>
                                <th class="text-center" style="font-size: 1.1em">% Vittorie</th>
                                <th class="text-center" style="font-size: 1.1em">Punti totali</th>
                                <th class="text-center" style="font-size: 1.1em">Punti/giornata</th>
                                <th class="text-center" style="font-size: 1.1em">Partite</th>
                                <th class="text-center" style="font-size: 1.1em">Punti/partita</th>
                                <th class="text-center" style="font-size: 1.1em">Punti di oggi</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr v-for="(g,i) in giocatori" :key="i">
                                <td>{{g.Nome}}</td>
                                <td>{{g.Partite}}</td>
                                <td>{{g.Vittorie}}</td>
                                <td>{{Math.round(g.Vittorie/g.Partite*100)}}%</td>
                                <td>{{g.PuntiTot}}</td>
                                <td>{{Math.round(g.PuntiTot/g.Partite*100)/100}}</td>
                                <td>{{g.Round}}</td>
                                <td>{{Math.round(g.PuntiTot/g.Round*100)/100}}</td>
                                <td>{{puntiDiOggi[g.Nome]}}</td>
                            </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                </v-card-text>
                <v-card-actions style="text-align: center; background-color: #90B494">
                    <v-spacer/>
                    <v-btn v-if="me === 0" @click="iniziaAltroRound()">GG, un'altra</v-btn>
                    <v-btn v-if="me === 0" @click="mandaMessaggio('bastaCosì')">GG, basta così</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script>
    import partitaApi from "../api/partitaApi";
    import BoxGiocatore from "../components/BoxGiocatore.vue"
    import Tavolo from '../components/Tavolo.vue'

    export default {
        name: "Partita",
        components:{
            BoxGiocatore,
            Tavolo
        },
        data: () => ({
            ws: {},
            me: 0,
            notYetDialog: false,
            chiamaDialog: false,
            vittoriaDialog: false,
            showChiamaAPuntiBox: false,
            showPunteggiTotali: false,
            puntiChiamati: 61,
            sonoProntoBtn: true,
            missingPlayers: '',
            carte: [],
            toccaA: 0,
            chiamabili: [],
            toccaAMe: false,
            chiamante: '',
            puntiVittoria: 61,
            showSemi: false,
            showCartaSocio: false,
            showCarteGirate: false,
            carteInMano: [],
            carteQuestaMano: {},
            selectedCarta: {},
            valChiamato: '',
            chiamanteProvvisorio: '',
            briscola: '',
            cartaChiamata: '',
            datiVittoria: {},
            giocatori: [],
            puntiDiOggi: {},
            mano: undefined
        }),
        mounted() {
            this.me = this.$store.state.giocatore.id
            this.registerWebSocket(this.me)
        },
        computed: {
            sommaPunti() {
                if (this.carte.length !== 8)
                    return 3
                let punti = 0
                for (let i = 0; i < this.carte.length; i++) {
                    let punto = this.carte[i].Punti
                    punti = punti + punto
                }
                return punti
            },
        },
        methods: {
            mandaMessaggio(azione, params) {
                let messaggio = {azione: azione, mittente: this.me, params: params}
                let json = JSON.stringify(messaggio)
                this.ws.send(json)
                //console.log("mandato " + json)
            },
            registerWebSocket(id) {
                var currentLocation = window.location;
                let port = currentLocation.port
                if (process.env.NODE_ENV == "development") {
                    port = "8080"
                }
                let baseURL = `//` + currentLocation.hostname + `:` + port;
                this.ws = new WebSocket("ws://" + baseURL + "/ws/registra/" + id);
                this.ws.onopen = function () {
                    //console.log(id + " Connesso")
                }
                this.ws.onmessage = (event) => {
                    let messaggio = JSON.parse(event.data)
                    //console.log("ricevuto " + event.data)
                    switch (messaggio.Azione) {
                        case "setRegistrati":
                            this.setRegistrati(messaggio)
                            break
                        case "iniziaPartita":
                            this.iniziaPartita(messaggio)
                            break
                        case "setChiamabili":
                            this.setChiamabili(messaggio)
                            break
                        case "iniziaRound":
                            this.iniziaRound(messaggio)
                            break
                        case "setGiocata":
                            this.setGiocata(messaggio)
                            break
                        case "giraCarte":
                            this.giraCarte(messaggio)
                            break
                        case "finePartita":
                            this.finePartita(messaggio)
                            break
                    }
                }
            },
            setRegistrati(resp) {
                this.missingPlayers = 5 - resp.Iscritti
            },
            sonoPronto() {
                this.mandaMessaggio('chiManca')
                this.mandaMessaggio("sonoPronto")
                this.notYetDialog = true
            },
            clearVariabili() {
                this.vittoriaDialog = false
                this.showSemi = false
                this.showPunteggiTotali = false
                this.puntiVittoria = 61
                this.notYetDialog = false
                this.sonoProntoBtn = false
                this.chiamante = ''
                this.briscola = ''
                this.showCartaSocio = false
                this.showCarteGirate = false
                this.valChiamato = ''
                this.puntiChiamati = 61
                this.chiamanteProvvisorio = ''
                this.puntiDiOggi = {},
                this.mano = undefined
            },
            iniziaPartita(resp) {
                this.clearVariabili()
                this.carte = resp.Carte
                this.$store.dispatch('setGiocatori', resp.CurrentGiocatori)
                this.setTurno(resp.ToccaA)
                this.chiamaDialog = true
            },
            setChiamabili(resp) {
                this.showChiamaAPuntiBox = false
                this.vittoriaDialog = false
                this.chiamaDialog = true
                this.chiamabili = resp.Chiamabili
                if (!this.chiamabili) {
                    this.chiamabili = []
                    this.showChiamaAPuntiBox = true
                    this.puntiChiamati = resp.PuntiVittoria + 1
                }
                this.chiamabili.push("passo")
                this.valChiamato = resp.ValChiamato
                this.chiamanteProvvisorio = resp.ChiamanteProvvisorio
                this.setTurno(resp.ToccaA)
                if (resp.Chiamante !== -1) {
                    this.chiamante = resp.Chiamante
                    this.puntiVittoria = resp.PuntiVittoria
                    if (this.chiamante === this.me) {
                        this.showSemi = true
                    }
                } else if (this.me === 0 && this.$store.state.giocatori[this.toccaA].IsBot === 1) {
                    this.mandaMessaggio("chiamaBot")
                }
            },
            iniziaRound(resp) {
                this.chiamaDialog = false
                this.setTurno(resp.ToccaA)
                this.briscola = resp.Briscola
                if (this.me === 0 && this.$store.state.giocatori[this.toccaA].IsBot) {
                    this.giocaCartaBot(this.toccaA)
                }
            },
            giraCarte() {
                this.showCarteGirate = true
            },
            iniziaAltroRound() {
                this.mandaMessaggio('altroRound')
            },
            finePartita(resp) {
                this.giocatori = resp.Giocatori
                this.showPunteggiTotali = true
            },
            setGiocata(resp) {
                this.$store.dispatch("setCarteGiocatori", {
                    inMano: resp.CarteInMano,
                    prese: resp.CartePrese
                })
                this.carteQuestaMano = resp.CarteGiocate
                this.setTurno(resp.ToccaA)
                if(this.mano !== resp.Mano){
                    this.mano = resp.Mano;
                }
                if (resp.Mano === 1) {
                    this.showCartaSocio = true
                }
                if (this.carteQuestaMano && Object.keys(this.carteQuestaMano).length === 5) {
                    let timeout = 3500;
                    if(this.mano === 0){
                        timeout = 5000;
                    }
                    setTimeout(() => this.svuotaTavolo(resp.Mano), timeout)
                } else {
                    if (this.me === 0 && this.$store.state.giocatori[this.toccaA].IsBot === 1) {
                        this.giocaCartaBot(this.toccaA)
                    }
                }
                this.$forceUpdate()
            },
            getStyleCarta(carta) {
                let border = "border: ;"
                if (carta.isSelected) {
                    border += "border: 5px solid #24ccf2;"
                }
                return "width: 9vh; margin-top: 30px; margin-left: 5px; margin-right: 5px; margin-bottom: 5px;" + border
            },
            getStyleGiocatore(pos) {
                return this.$store.getters.getStyleGiocatore(pos, this.toccaA, this.chiamante);
            },
            getNomeCarta(val) {
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
            getNomeSeme(seme) {
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
            toggleSelection(carta) {
                if (this.toccaAMe) {
                    for (let i = 0; i < this.carte.length; i++) {
                        this.carte[i].isSelected = false
                    }
                    carta.isSelected = true
                    this.selectedCarta = carta
                    this.$forceUpdate()
                }
            },
            getNomeGiocatore(id) {
                return this.$store.state.giocatori[id] ? this.$store.state.giocatori[id].Nome : ''
            },
            chiamaSeme(seme) {
                this.mandaMessaggio("chiamaSeme", [seme.toString()])
            },
            svuotaTavolo(mano) {
                this.carteQuestaMano = []
                if (mano === 7) {
                    this.mostraVittoria()
                } else {
                    this.iniziaNuovaMano()
                }
            },
            iniziaNuovaMano() {
                this.mandaMessaggio("iniziaNuovaMano")
            },
            mostraVittoria() {
                partitaApi.getVittoria().then(t => {
                    if (t.status === 200) {
                        this.datiVittoria = t.data
                        this.vittoriaDialog = true
                        for (let i=0; i<this.datiVittoria.PuntiPartita.length; i++){
                            this.puntiDiOggi[this.getNomeGiocatore(i)] = this.datiVittoria.PuntiPartita[i]
                        }
                    }
                })
            },
            setTurno(toccaA) {
                this.toccaA = toccaA
                this.toccaAMe = this.toccaA === this.me;
            },
            getPrese(pos) {
                return this.$store.getters.getPrese(pos);
            },
            giocaCarta() {
                if (this.toccaAMe && this.selectedCarta.Valore !== 0) {
                    this.mandaMessaggio("giocaCarta", [this.selectedCarta.Valore.toString(),
                        this.selectedCarta.SemeStr])
                    let index = this.carte.indexOf(this.selectedCarta);
                    if (index > -1) {
                        this.carte.splice(index, 1);
                    }
                    this.selectedCarta = {}
                }
            },
            tuttoNostro() {

            },
            getGiocatoreInPos(pos) {
                //return ((this.me + pos) % 5)
                return this.$store.getters.getGiocatoreInPos(pos);
            },
            setBots() {
                this.mandaMessaggio("addBots")
            },
            chiamaCartaBot(id) {
                partitaApi.chiamaBot(id)
            },
            giocaCartaBot(id) {
                setTimeout(() => this.mandaMessaggioBot(id), 1000)
            },
            mandaMessaggioBot(id) {
                this.mandaMessaggio("giocaCartaBot", [id.toString()])
            },
            getVincitoriStr(vinc) {
                if (!vinc)
                    return ''
                let vincitori = ''
                for (let i = 0; i < vinc.length; i++) {
                    vincitori += this.getNomeGiocatore(vinc[i]) + ', '
                }
                return vincitori
            },
            getLarghezzaDialogVittoria() {
                if (this.showPunteggiTotali) {
                    return 900
                }
                return 400
            }
        }
    }
</script>

<style scoped>

.transform {
  /*transform: translateX(-150px);*/
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
