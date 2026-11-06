import { createRouter, createWebHistory } from "vue-router"
import aic from "./page/aicommunication.vue"
import home from "./page/home.vue"
import crud from "./page/crud.vue"
import login from "./page/login.vue"
import { islog } from "./hanshu"

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", redirect: "/home" },
        { path: "/home", component: home, meta: { reqlog: true, navigation: false } },
        { path: "/aic", component: aic, meta: { reqlog: true, navigation: true } },
        { path: "/crud", component: crud, meta: { reqlog: true, navigation: true } },
        { path: "/log", component: login },
                                            
    ],
})
router.beforeEach(async (to) => {
    const logstatus = await islog()
     if (to.path=="/log"&&logstatus==true){
        return "/home"
    }
    if (to.meta.reqlog==false){
        return true
    }
    if (to.meta.reqlog==true&&logstatus==false){
        return "/log"
    }
   



})
export default router
