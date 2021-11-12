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
                    <v-btn v-if="sonoProntoBtn" color="#718F94" style="color: white;" @click="sonoPronto()">Sono
                    pronto
                    </v-btn>
                    <v-btn v-if="sommaPunti < 3" color="#718F94" style="color: white" @click="aMonte">a
                        monte
                    </v-btn>
                    <div v-if="toccaAMe">
                        <v-btn color="#718F94" style="color: white" @click="giocaCarta">Giocala</v-btn>
                        <!--v-btn class="ma-3" color="#718F94" style="color: white" @click="tuttoNostro">Tutto nostro</v-btn-->
                    </div>
                </v-layout>
                <v-layout style="justify-content: center; height: 24vh">
                    <v-flex xs2 style="overflow-y: auto">
                        <section id="log">
                            <b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Tocca a: </b><b style="font-size: 2vh; color: #6e0021">{{getNomeGiocatore(toccaA)}}</b>
                            <ul>
                                <li v-for="logMessage in logMessages" :key="logMessage">
                                    <b v-if="logMessage.chi" style="color: #4454e3">{{logMessage.chi}}</b>
                                    <b :style="logMessage.colorMessage">{{logMessage.azione}}</b>
                                    <b style="color: #4454e3" v-if="logMessage.valore">{{logMessage.valore}}<b style="color: #4454e3" v-if="logMessage.seme"> di {{logMessage.seme}}</b></b>
                                    <b style="font-style: italic; color: #158a42" v-if="logMessage.soprannomeCarta !== undefined && logMessage.soprannomeCarta !== ''">({{logMessage.soprannomeCarta}})</b>
                                </li>
                            </ul>
                        </section>
                    </v-flex>
                    <v-flex xs8 :style="getStyleGiocatore(0)">
                        <v-btn v-if="carte.length > 0" color="#718F94" style="color: white; float: right; margin: 5px" @click="toggleCarte">{{getNomePulsanteToggleCarte()}}</v-btn>
                        <v-layout style="justify-content: center; margin-left: 4vw;" ref="box0">
                            <div v-for="carta in carte" :key="carta">
                                <v-img :class="{transform: selectedCarta === carta}"
                                        :style="getStyleCarta(carta)"
                                        contain
                                        :src="getSpriteCartaGiocatore(carta)"
                                        @click="toggleSelection(carta)"
                                ></v-img>
                            </div>
                        </v-layout>
                    </v-flex>
                    <v-flex xs2>                
                        <v-layout v-if="showCartaSocio" style="justify-content: center"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Carta del socio:</b></v-layout>
                        <v-layout v-else style="justify-content: center"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Carta chiamata:</b></v-layout>
                        <v-layout style="justify-content: center;">
                            <b style="font-size: 2vh; color: #6e0021">{{cartaChiamata}}</b>
                        </v-layout>
                        <v-layout style="justify-content: center" class="mt-4"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Chiamante:</b><b style="font-size: 2vh; color: #6e0021">{{getNomeGiocatore(chiamante)}}</b></v-layout>
                        <v-layout style="justify-content: center" class="mt-4"><b style="font-size: 2vh; font-style: italic; padding-right: 0.5vh;">Punti per vincere:</b><b style="font-size: 2vh; color: #6e0021">{{puntiVittoria}}</b></v-layout>
                    </v-flex>
                </v-layout>
            </v-col>
            <v-col md="2">
                <the-chat :enabled="abilitaChat" :myName="getNomeGiocatore(me)" :logChat="logChat" @mandaMessaggioChat="mandaMessaggioChat"></the-chat>
            </v-col>
        </v-row>

        <v-dialog
                v-model="chiamaDialog"
                width="400"
                persistent
        >
            <v-card>
                <v-card-title class="headline lighten-2 pa-3" style="background-color: #DBCFB0">
                    <b>Fase di Chiamata</b>
                    <v-spacer/>
                    <div v-if="getNomeCarta(valChiamato) !== 'niente'" class="ma-3"> 
                        <b style="color: #4454e3">{{getNomeGiocatore(chiamanteProvvisorio)}}</b> ha chiamato: <b style="color: #4454e3">{{getNomeCarta(valChiamato)}}</b>
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

    export default {
        name: "Partita",
        components:{
            BoxGiocatore,
            Tavolo,
            TheChat
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
            datiVittoria: {},
            giocatori: [],
            puntiDiOggi: {},
            mano: undefined,
            coordGiocatore: undefined,
            haIniziatoIlRound: undefined,
            logMessages: [],
            carteChiamate:[],
            logChat: [],
            abilitaChat: false,
            carteVisibili: true
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
                let val = this.getNomeCarta(this.valChiamato);
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
                this.logChat = []
            },
            iniziaPartita(resp) {
                this.clearVariabili()
                this.carte = resp.Carte
                this.$store.dispatch('setGiocatori', resp.CurrentGiocatori)
                this.setTurno(resp.ToccaA)
                this.chiamaDialog = true
                this.abilitaChat = true
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

                let nomeCarta = this.getNomeCarta(this.valChiamato);
                const carteChiamateStr = this.chiamanteProvvisorio + "-"+ this.valChiamato + "-" + this.puntiChiamati;
                if(nomeCarta != undefined && nomeCarta !== 'niente' && this.carteChiamate.indexOf(carteChiamateStr) === -1) {
                    this.carteChiamate.push(carteChiamateStr);  //per evitare che lo stesso log venga spammato più volte in chat
                    let nomeGiocatore = this.getNomeGiocatore(this.chiamanteProvvisorio);
                    if(this.puntiChiamati > 61){
                        const punti = this.puntiChiamati - 1;
                        this.addLogMessage(" ha chiamato ", nomeGiocatore, this.getNomeCarta(this.valChiamato) + " a " + punti);
                    }
                    else{
                        this.addLogMessage(" ha chiamato ", nomeGiocatore, this.getNomeCarta(this.valChiamato));
                    }
                    
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
                            let nomeCarta = this.getNomeCarta(carta.Valore);
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
            getStyleCarta(carta) {
                let border = "border: ;"
                if (carta.isSelected) {
                    border += "border: 5px solid #24ccf2;"
                }
                return "width: 8vh; margin-top: 30px; margin-left: 2px; margin-right: 2px; margin-bottom: 5px;" + border
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
                                el = this.$refs.box0;    

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
            toggleCarte(){
                this.carteVisibili = !this.carteVisibili
            },
            getNomePulsanteToggleCarte(){
                if(this.carteVisibili){
                    return 'Nascondi';
                }
                else{
                    return 'Mostra';
                }
            },
            getSpriteCartaGiocatore(carta){
                if(this.carteVisibili){
                    return require(`../assets/${this.$store.state.mazzo}/${carta.Valore + carta.SemeStr}.png`);
                }
                else{
                    return require(`../assets/${this.$store.state.mazzo}/retro.png`);
                }
            },
            aMonte(){
                
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

#log ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

#log li {
  font-size: 1.6vh;
  margin: 0.3rem 1rem;
  text-align: left;
}

.transform {
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
