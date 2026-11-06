import "./style.css" 
import app from "./app.vue"
import router from "./router"
import {createApp} from "vue"

createApp(app).use(router).mount("#app")