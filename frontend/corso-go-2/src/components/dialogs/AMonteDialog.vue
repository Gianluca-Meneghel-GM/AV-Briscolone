<template>
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
</template>

<script>
export default {
    props:['aMonteDialog','nomeGiocatoreAMonte']
}
</script>