import VueRouter from "vue-router"
import Login from "./views/Login";
import Partita from "./views/Partita";

const routes = [
    { path: '/', name: 'LOGIN', component: Login },
    { path: '/partita', name: 'PARTITA', component: Partita },
]

const router = new VueRouter({
    routes // short for `routes: routes`
})


export default router

