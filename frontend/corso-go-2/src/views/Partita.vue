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

        <v-row class="board">
            <v-col md="10">
                <v-layout style="height:18vh;">
                    <box-giocatore ref="box3" :pos="3" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
                    <box-giocatore ref="box2" :pos="2" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
                </v-layout>
                <v-layout style="height:45vh">
                    <box-giocatore ref="box4" :pos="4" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
                    <tavolo :coordGiocatore="coordGiocatore" :mano="mano" :valChiamato="valChiamato" :carteQuestaMano="carteQuestaMano"></tavolo>
                    <box-giocatore ref="box1" :pos="1" :toccaA="toccaA" :chiamante="chiamante"></box-giocatore>
                </v-layout>
                <v-layout style="justify-content: center; height:4vh">
                    <div class="ma-3"><b style="font-style: italic; padding-right: 0.5vh;">Prese:</b><b style="color:#4454e3; padding-right: 0.5vh;">{{getPrese(0)}}</b></div>
                    <v-btn v-if="sonoProntoBtn" class="ma-3" color="#718F94" style="color: white;" @click="sonoPronto()">Sono
                    pronto
                    </v-btn>
                    <div v-if="toccaAMe">
                        <v-btn color="#718F94" class="ma-3" style="color: white" @click="giocaCarta">Giocala</v-btn>
                        <!--v-btn class="ma-3" color="#718F94" style="color: white" @click="tuttoNostro">Tutto nostro</v-btn-->
                    </div>
                </v-layout>
                <v-layout style="justify-content: center; height: 24vh">
                    <v-flex xs2 style="overflow-y: auto">
                        <section id="log">
                            <b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Tocca a: </b><b style="font-size: 2vh; color: #6e0021">{{getNomeGiocatore(toccaA)}}</b>
                            <the-logs :logMessages="logMessages"></the-logs>
                        </section>
                    </v-flex>
                    <pannello-carte-giocatore ref="box0" :carte="carte" :toccaA="toccaA" :chiamante="chiamante" :selectedCarta="selectedCarta" @cartaSelected="cartaSelected"></pannello-carte-giocatore>
                    <informazioni-chiamante :timerVisible="roundIniziato && toccaAMe" 
                                            :showCartaSocio="showCartaSocio" 
                                            :cartaChiamata="cartaChiamata" 
                                            :puntiVittoria="puntiVittoria" 
                                            :chiamante="getNomeGiocatore(chiamante)"
                                            @timesUp="timesUp">
                    </informazioni-chiamante>
                </v-layout>
            </v-col>
            <v-col md="2">
                <the-chat :enabled="roundIniziato" :myName="getNomeGiocatore(me)" :logChat="logChat" @mandaMessaggioChat="mandaMessaggioChat"></the-chat>
            </v-col>
        </v-row>

        <v-dialog
                v-model="aMonteDialog"
                width="400"
                persistent
        >
            <v-card style="text-align: center; background-color: #90B494">
                <v-card-title class="headline lighten-2" style="background-color: #DBCFB0">
                    A Monte!
                </v-card-title>

                <v-card-text class="mt-4">
                    <b>{{nomeGiocatoreAMonte}} ha mandato a monte!</b>
                </v-card-text>
            </v-card>
        </v-dialog>

        <v-dialog
                v-model="chiamaDialog"
                width="400"
                persistent
        >
            <v-card>
                <v-card-title class="headline lighten-2 pa-3" style="background-color: #DBCFB0">
                    <b>Fase di Chiamata</b>
                    <v-spacer/>
                    <div v-if="getValoreCarta(valChiamato) !== 'niente'" class="ma-3"> 
                        <b style="color: #4454e3">{{getNomeGiocatore(chiamanteProvvisorio)}}</b> ha chiamato: <b style="color: #4454e3">{{getValoreCarta(valChiamato)}}</b>
                    </div>
                </v-card-title>
                <v-card-text style="text-align: center; background-color: #90B494">
                    <div v-if="chiamante === ''">
                        <v-btn class="ma-3" v-for="val in chiamabili" :disabled="!toccaAMe" :key="val" :style="getStyleChiamabili(val)"
                               @click="chiamaCarta(val, puntiChiamati)">
                            {{getValoreCarta(val)}}
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
                        Han vinto <b style="color: #4454e3">{{getVincitoriStr(datiVittoria.Vincitori)}}</b> <b style="color: #064723">{{datiVittoria.PuntiVincitori}}</b> a <b style="color: #a01118">{{120 -
                        datiVittoria.PuntiVincitori}}</b>
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
                            <tr v-for="g in giocatori" :key="g.Nome">
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
    import TheChat from '../components/TheChat.vue'
    import TheLogs from '../components/TheLogs.vue'
    import PannelloCarteGiocatore from "../components/PannelloCarteGiocatore.vue";
    import InformazioniChiamante from "../components/InformazioniChiamante.vue";

    export default {
        name: "Partita",
        components:{
            BoxGiocatore,
            Tavolo,
            TheChat,
            TheLogs,
            PannelloCarteGiocatore,
            InformazioniChiamante
        },
        data: () => ({
            ws: {},
            me: undefined,
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
            datiVittoria: {},
            giocatori: [],
            puntiDiOggi: {},
            mano: undefined,
            coordGiocatore: undefined,
            haIniziatoIlRound: undefined,
            logMessages: [],
            carteChiamate:[],
            logChat: [],
            roundIniziato: false,
            nomeGiocatoreAMonte: '',
            aMonteDialog: false
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
            cartaChiamata(){
                let val = this.getValoreCarta(this.valChiamato);
                if(this.showCartaSocio){
                    val = val + ' di ' + this.getNomeSeme(this.briscola)
                }
                return val
            }
        },
        methods: {
            mandaMessaggio(azione, params) {
                let messaggio = {azione: azione, mittente: this.me, params: params}
                let json = JSON.stringify(messaggio)
                this.ws.send(json)
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
                        case "showMessaggioChat":
                            this.showMessaggioChat(messaggio)
                            break
                        case "aMonte":
                            this.aMonte(messaggio)
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
                this.mano = undefined,
                this.logMessages = [],
                this.carteChiamate = [],
                this.nomeGiocatoreAMonte = '',
                this.aMonteDialog = false
            },
            iniziaPartita(resp) {
                this.clearVariabili()
                this.carte = resp.Carte
                this.$store.dispatch('setGiocatori', resp.CurrentGiocatori)
                this.setTurno(resp.ToccaA)
                this.chiamaDialog = true
                this.roundIniziato = false;
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
                if(this.sommaPunti < 3){
                    this.chiamabili.push("a monte");
                }

                this.valChiamato = resp.ValChiamato
                this.chiamanteProvvisorio = resp.ChiamanteProvvisorio

                let nomeCarta = this.getValoreCarta(this.valChiamato);
                const carteChiamateStr = this.chiamanteProvvisorio + "-"+ this.valChiamato + "-" + this.puntiChiamati;
                if(nomeCarta != undefined && nomeCarta !== 'niente' && this.carteChiamate.indexOf(carteChiamateStr) === -1) {
                    this.carteChiamate.push(carteChiamateStr);  //per evitare che lo stesso log venga spammato più volte in chat
                    let nomeGiocatore = this.getNomeGiocatore(this.chiamanteProvvisorio);
                    if(this.puntiChiamati > 61){
                        const punti = this.puntiChiamati - 1;
                        this.addLogMessage(" ha chiamato ", nomeGiocatore, this.getValoreCarta(this.valChiamato) + " a " + punti);
                    }
                    else{
                        this.addLogMessage(" ha chiamato ", nomeGiocatore, this.getValoreCarta(this.valChiamato));
                    }
                    
                }
                if(resp.HaPassato !== -1 && resp.ValChiamato !==  0) {
                    const nomeGiocatore = this.getNomeGiocatore(resp.HaPassato);
                    this.addLogMessage(" ha passato ", nomeGiocatore);
                }

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
                this.roundIniziato = true
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
            iniziaAltroRound(aMonte = false) {
                this.mandaMessaggio('altroRound', [aMonte.toString()])
            },
            finePartita(resp) {
                this.giocatori = resp.Giocatori
                this.showPunteggiTotali = true
            },
            setGiocata(resp) {
                let carte = {
                    inMano: resp.CarteInMano,
                    prese: resp.CartePrese
                };
                this.$store.dispatch("setCarteGiocatori", carte).then( () => {
                    if(resp.CarteGiocate && Object.keys(resp.CarteGiocate).length == 1) {
                        this.haIniziatoIlRound = Number(Object.keys(resp.CarteGiocate)[0]);
                    }
                    this.carteQuestaMano = resp.CarteGiocate

                    if(this.mano !== resp.Mano){
                        this.mano = resp.Mano;
                        let manoStr = Number(this.mano) + 1;
                        this.addLogMessage("Inizio mano: " + manoStr, undefined, undefined, undefined, undefined, "#c24444");
                    }

                    if(this.toccaA != null && this.toccaA !== -1){
                        let chiHaGiocato = this.$store.state.giocatori[this.toccaA];
                        if(chiHaGiocato != null && this.carteQuestaMano != null && this.carteQuestaMano[chiHaGiocato.Id]){
                            const carta = this.carteQuestaMano[chiHaGiocato.Id];
                            const isCoperta = carta.Valore === this.valChiamato && this.mano === 0;
                            let nomeCarta = this.getValoreCarta(carta.Valore);
                            let nomeSeme = this.getNomeSeme(carta.SemeStr);
                            let soprannome = carta.Soprannome;
                            if(isCoperta){
                                nomeCarta = nomeCarta + " (coperta)"
                                nomeSeme = undefined;
                                soprannome = undefined;
                            }
                            this.addLogMessage(" ha giocato ", chiHaGiocato.Nome, nomeCarta, nomeSeme, soprannome);
                        }
                    }

                    this.setTurno(resp.ToccaA)
                    if (resp.Mano === 1) {
                        this.showCartaSocio = true
                    }

                    if (this.carteQuestaMano && Object.keys(this.carteQuestaMano).length === 5) {
                        let timeout = 3500;
                        if(this.mano === 0){
                            timeout = 5000;
                        }
                        this.aggiornaCoordGiocatoreCheHaPreso(resp.CarteGiocate);
                        setTimeout(() => this.svuotaTavolo(resp.Mano), timeout)
                    }
                    else {
                        if (this.me === 0 && this.$store.state.giocatori[this.toccaA].IsBot === 1) {
                            this.giocaCartaBot(this.toccaA)
                        }
                    }
                    this.$forceUpdate()
                });
            },
            getValoreCarta(val) {
                return this.$store.getters.getValoreCarta(val);
            },
            getNomeSeme(seme) {
                return this.$store.getters.getNomeSeme(seme);
            },
            getNomeGiocatore(id) {
                return this.$store.state.giocatori[id] ? this.$store.state.giocatori[id].Nome : ''
            },
            chiamaCarta(val, puntiChiamati){
                this.mandaMessaggio('chiamaValore', [val.toString(), puntiChiamati.toString()])
            },
            chiamaSeme(seme) {
                this.mandaMessaggio("chiamaSeme", [seme.toString()])
            },
            svuotaTavolo(mano) {
                this.carteQuestaMano = []
                if (mano === 7) {
                    this.mostraVittoria()
                }
                else {
                    this.iniziaNuovaMano()
                }
            },
            aggiornaCoordGiocatoreCheHaPreso(carteGiocate){
                try{
                    let giocatoreCheHaPreso = this.getGiocatoreCheHaVintoLaMano(carteGiocate);
                    if(giocatoreCheHaPreso >= 0) {
                        this.haIniziatoIlRound = undefined;
                        let posizione = this.getPosizioneFromIdGiocatore(giocatoreCheHaPreso);
                        if(posizione >= 0){
                            let el;
                            if(posizione === 0){
                                el = this.$refs.box0.$el;    

                            }
                            else if(posizione === 1){
                                el = this.$refs.box1.$el;    
                            }
                            else if(posizione === 2){
                                el = this.$refs.box2.$el;  
                            }
                            else if(posizione === 3){
                                el = this.$refs.box3.$el;  
                            }
                            else if(posizione === 4){
                                el = this.$refs.box4.$el;  
                            }

                            if(el != undefined){
                                this.coordGiocatore = {
                                    pos: posizione,
                                    offsetLeft: el.offsetLeft,
                                    offsetTop: el.offsetTop,
                                    width: el.getBoundingClientRect().width,
                                    height: el.getBoundingClientRect().height
                                }
                            }
                        }
                    }
/*                     else{
                        console.log("Impossibile calcolare il giocatore che si è preso la mano..");
                    } */
                }
                catch(e){
                    //console.log(e);
                }
            },
            getGiocatoreCheHaVintoLaMano(carteGiocate){
                let giocatoreCheHaPreso = -1;
                let tmpValore = -1

                let briscolaMano = "";
                for (let key in carteGiocate) {
                    if(carteGiocate[key].SemeStr === this.briscola){
                        briscolaMano = this.briscola;
                        break;
                    }

                    let idGioc = Number(key);
                    if(idGioc === this.haIniziatoIlRound){
                        briscolaMano = carteGiocate[key].SemeStr;
                    }
                }

                //console.log("************      MANO => " + this.mano + "      ****************")
                //console.log("briscola di mano: "+ briscolaMano)
                for (let key in this.carteQuestaMano) {
                    let valCarta = this.carteQuestaMano[key].Valore;
                    let semeCarta = this.carteQuestaMano[key].SemeStr
                    let calcValore = briscolaMano === semeCarta ? valCarta + 100 : valCarta;
                    //console.log("giocatore: " + key + " seme: " + semeCarta + " val: " + valCarta + " => valore calcolato => " + calcValore);
                    if (calcValore > tmpValore) {
                        tmpValore = calcValore;
                        giocatoreCheHaPreso = Number(key);
                        //let nome = this.$store.state.giocatori[giocatoreCheHaPreso].Nome
                        //console.log("prende: " + giocatoreCheHaPreso + " => " + nome);
                    }
                }
                //console.log("************      FINE MANO     ****************")
                return giocatoreCheHaPreso;
            },
            getPosizioneFromIdGiocatore(giocatoreCheHaPreso){
                for(let i = 0; i < 5; i++){
                    if(this.getGiocatoreInPos(i) === giocatoreCheHaPreso){
                        //let nome = this.$store.state.giocatori[giocatoreCheHaPreso].Nome
                        //console.log("########## ha preso il giocatore " + nome + " in posizione => " + i);
                        return i;
                    }
                }
                return 0;
            },
            addLogMessage(azione, chi = undefined, valore = undefined, seme = undefined, soprannomeCarta = undefined, colorMessage = "#000000") {
                this.logMessages.unshift({
                    azione: azione,
                    chi: chi,
                    valore: valore,
                    soprannomeCarta: soprannomeCarta,
                    seme: seme,
                    colorMessage: "color: " + colorMessage
                });
            },
            showMessaggioChat(resp){
                let nome = this.getNomeGiocatore(resp.Giocatore);
                let isMe = this.me === resp.Giocatore;
                this.logChat.unshift({
                    id: this.logChat.length,
                    chi: nome,
                    messaggio: resp.Messaggio,
                    me: isMe,
                    timestamp: new Date().toLocaleTimeString()
                });
            },
            mandaMessaggioChat(messaggio) {
                this.mandaMessaggio("mandaMessaggioChat", [messaggio]);
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
            timesUp() {
                if(this.selectedCarta && Object.keys(this.selectedCarta).length === 0 && Object.getPrototypeOf(this.selectedCarta) === Object.prototype) {
                    const index = Math.floor(Math.random() * this.carte.length);
                    this.cartaSelected(this.carte[index])
                }
                this.giocaCarta();
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
            cartaSelected(carta){
                if (this.toccaAMe) {
                    for (let i = 0; i < this.carte.length; i++) {
                        this.carte[i].isSelected = false
                    }
                    carta.isSelected = true
                    this.selectedCarta = carta
                    this.$forceUpdate()
                }
            },
            aMonte(resp){
                this.nomeGiocatoreAMonte = this.getNomeGiocatore(resp.Giocatore);
                this.aMonteDialog = true;
                this.chiamaDialog = false;
                setTimeout(() => {
                    this.iniziaAltroRound(true);
                }, 5000);
            },
            getStyleChiamabili(val){
                let style = "";
                if(val === "a monte" || val === "passo") {
                    style = "background-color: #4454e3; color: white";
                }
                return style;
            },
            tuttoNostro() {

            },
            getGiocatoreInPos(pos) {
                return this.$store.getters.getGiocatoreInPos(pos);
            },
            setBots() {
                this.mandaMessaggio("addBots")
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
