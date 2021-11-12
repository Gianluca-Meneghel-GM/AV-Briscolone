<template>
    <v-card
        flat
        class="d-flex flex-column fill-height"
        style="margin-right: 1vw;"
        >
        <v-card-title>
            <v-chip
                    color="primary"
                    dark
                    style="white-space: normal;"
                    class="ma-2"
                    large
                    >
                    <v-icon left>
                    mdi-account-circle-outline
                    </v-icon>
            {{myName}}
        </v-chip>
        </v-card-title>
        <v-card-text style="max-height: 75vh;" class="flex-grow-1 overflow-y-auto">
            <ul>
                <li v-for="msg in logChat"
                    :key="msg.id"
                    class="chatfield p-2"
                    :class="{ self: msg.me }">

                    <v-row style="width: 12vw; word-wrap: break-word; display: inline-block;">
                        <div style="width: 12vw; word-wrap: break-word; display: inline-block;">
                            <span v-if="!msg.me" style="font-weight: bold; padding-right: 0.5vh; color: #4454e3; font-size: 1rem">{{ msg.chi }}:</span>
                            <span :class="{ messaggio: msg.me }" style="padding-right: 0.5vh; font-size: 1rem; word-wrap: break-word;" v-html="convertToLink(msg.messaggio)"></span>
                        </div>

                        <div style="font-size: 0.7rem; font-style: italic; margin-left: 5px; color: rgb(79, 79, 85)">{{ (msg.timestamp) }}</div>
                    </v-row>
                </li>
            </ul>
        </v-card-text>
        <v-footer absolute>
            <v-container class="ma-0 pa-0">
                <v-row no-gutters>
                <v-col>
                    <div class="d-flex flex-row align-center">
                        <v-text-field style="margin: 2px;"
                            :disabled="!enabled"
                            v-model="messaggioChat"
                            append-outer-icon="mdi-send"
                            solo
                            clearable
                            label="Messaggio"
                            type="text"
                            @click:append-outer="mandaMessaggioChat()"
                            @keyup.enter="mandaMessaggioChat()"
                        ></v-text-field>
                    </div>
                </v-col>
                </v-row>
            </v-container>
        </v-footer>
    </v-card>

</template>

<script>
export default {
    name: "TheChat",
    props:['logChat', 'myName', 'enabled'],
    emits:['mandaMessaggioChat'],
    data: () => ({
        messaggioChat: ""
    }),
    methods:{
        mandaMessaggioChat(){
            if(this.messaggioChat != null && this.messaggioChat !== "") {
                this.$emit('mandaMessaggioChat', this.messaggioChat);
            }
            this.pulisciMessaggiochat();
            
        },
        pulisciMessaggiochat(){
            this.messaggioChat = "";
        },
        convertToLink: function (text) {
            const URLMatcher =
                /(?:(?:https?|ftp|file):\/\/|www\.|ftp\.)(?:\([-A-Z0-9+&@#/%=~_|$?!:,.]*\)|[-A-Z0-9+&@#/%=~_|$?!:,.])*(?:\([-A-Z0-9+&@#/%=~_|$?!:,.]*\)|[A-Z0-9+&@#/%=~_|$])/gim;
            return text.replace(
                URLMatcher,
                (text) =>
                `<a class="text-primary" rel="noreferrer" target="_blank" href="${text}">${text}</a>`,
            );
        },
    }
}
</script>


<style scoped>

ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

li {
  margin: 0.3rem 1rem;
  text-align: left;
}

.chatfield {
  color: black;
  width: fit-content;
  word-wrap: break-word;
}

.messaggio{
    color: #541166;
    font-weight: bold;
}

.self {
  text-align: right;
  float: right;
  flex-flow: row-reverse;
}
</style>
