<template>
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
                    <v-btn @click="addBots">Gioca coi bot</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
</template>

<script>
export default {
    props:['missingPlayers', 'me', 'notYetDialog'],
    emits:['addBots'],
    methods:{
        addBots() {
            this.$emit('addBots');
        }
    }
}
</script>