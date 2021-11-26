<template>
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
                <v-btn v-if="me === 0" @click="bastaCosi()">GG, basta così</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    props:['vittoriaDialog','showPunteggiTotali','datiVittoria','giocatori','me','puntiDiOggi'],
    emits:['iniziaAltroRound','bastaCosi'],
    methods:{
        getLarghezzaDialogVittoria() {
            if (this.showPunteggiTotali) {
                return 900
            }
            return 400
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
        getNomeGiocatore(id) {
            return this.$store.state.giocatori[id] ? this.$store.state.giocatori[id].Nome : ''
        },
        iniziaAltroRound(){
            this.$emit('iniziaAltroRound');
        },
        bastaCosi(){
            this.$emit('bastaCosi');
        }
    }
}
</script>