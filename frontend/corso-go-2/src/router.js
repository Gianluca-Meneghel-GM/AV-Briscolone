import VueRouter from "vue-router"
import Login from "./views/Login";
import Partita from "./views/Partita";
import store from "./store/store.js"

const routes = [
    { path: '/', redirect: '/login' },
    { path: '/login', name: 'login', component: Login },
    { path: '/partita', name: 'partita', component: Partita, meta: { requiresAuth: true } },
]

const router = new VueRouter({
    routes // short for `routes: routes`
})

router.beforeEach((to, from, next) => {
    if(to.meta.requiresAuth && store.state.giocatore.id === undefined){
        next('/login')
    }
    else{
        next();
    }
});


export default router

