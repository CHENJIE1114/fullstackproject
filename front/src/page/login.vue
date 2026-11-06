<script setup lang="ts">
import {ref} from "vue" 
import {useRouter} from "vue-router"
const name=ref("")
const pw=ref("")
const ing=ref("")
const r =useRouter()

async function or(){
if(pw.value.trim()==""||name.value.trim()==""){
    return 
}
ing.value="true"
try{
    const response=await fetch("/api/or",{
   method:"POST",
   headers:{"Content-Type":"application/json"},
   body:JSON.stringify({
        name:name.value,
        pw: pw.value,
    }),
    })
 if(!response.ok){throw new Error()}
r.push("/home")
}catch(err){
return 
}finally{
    ing.value=""
}
}

</script>

<template>
    <div class="back">
    <div class="wai">
    <form @submit.prevent="or" class="f">
    <input v-model="name" placeholder="name" required class="input1">
    <input v-model="pw" placeholder="password" required type="password" class="input1">
    <button>{{ing?"正在登录":"登录/注册"}}</button>
    </form>
</div>
</div>
</template>

<style scoped>

.wai{
    width:20vw;
    margin:40vh auto 0;
    font-size:50px;
    
}
.back{
   
  
    background-image:url("../image/c.png")
}
.f{
  display:flex;
    flex-direction:column;
    gap:30px;  
}
input{
    font-size:22px;
}
button{
    width:7vw;
    margin-left:auto;
    border-radius:5px;
    padding:7px;
    background:none;
}
.input1{
background-color:transparent;
border:none;
}

</style>