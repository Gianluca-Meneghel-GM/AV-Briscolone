<template>
    <ul>
        <li v-for="(logMessage, i) in logMessages" :key="i">
            <b v-if="logMessage.chi" style="color: #4454e3">{{logMessage.chi}}</b>
            <b :style="logMessage.colorMessage">{{logMessage.azione}}</b>
            <b style="color: #4454e3" v-if="logMessage.valore">{{logMessage.valore}}<b style="color: #4454e3" v-if="logMessage.seme"> di {{logMessage.seme}}</b></b>
            <b style="font-style: italic; color: #158a42" v-if="logMessage.soprannomeCarta !== undefined && logMessage.soprannomeCarta !== ''">({{logMessage.soprannomeCarta}})</b>
        </li>
    </ul>
</template>

<script>
export default {
    props:['logMessages']
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

</style>