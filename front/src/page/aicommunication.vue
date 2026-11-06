/*<script setup lang="ts">
import {marked} from "marked"
import DOMPurify from "dompurify"
 import {ref,onMounted} from "vue"
type message ={
    role:"user"|"assistant";
    content:string;
}
const input=ref("")
const messages=ref<message[]>([])
const asking=ref(false)
function mark(con :string):string{
    const html=marked.parse(con) as string
    return DOMPurify.sanitize(html) 
}
function loadcon(){
    const jsoncon=localStorage.getItem("aicon")
     messages.value=JSON.parse(jsoncon??"[]")
}
function saveaicon(){
    const jsoncon=JSON.stringify(messages.value)
    localStorage.setItem("aicon",jsoncon)
}
function delete1(){
    localStorage.setItem("aicon","[]")
    loadcon()
}
async function ask(){
    if(input.value.trim()==""){
        return
    }
    messages.value.push({
        role:"user",
        content:input.value,
    })
    
    const sendmessage=JSON.stringify({
        data:messages.value,
    })
   try{     asking.value=true
            const response=await fetch("/ai/llm",{
              method:"POST",
              headers:{"Content-Type":"application/json"},
              body:sendmessage,
         })
         if(!response.ok){
            throw new Error("err in ai communication")
         }
         if(!response.body){
            throw new Error("err in ai communication")
         }
         const reader=response.body.getReader()
         const decoder=new TextDecoder()
          messages.value.push({
        role:"assistant",
        content:"",
            })
         while(true){
            const {value,done}=await reader.read()
            if (done){
                break
            }
            const index=messages.value.length
            messages.value[index-1].content+=decoder.decode(value,{stream:true})
         }
         saveaicon()
    loadcon()
        }catch(err){
            console.log("err in alllm")
    }finally{
        asking.value=false
    }
}


onMounted(loadcon)
</script>
<template>
   <div class="div1">
    <div class="chat">
        <div v-for="value in messages">
            <div class="role" v-if="value.role=='user'">
                <p class="p1">{{value.content}}</p>
            </div>
            <div class="content" v-else>
                <div class="p2" v-html="mark(value.content)">
                </div>
            </div>
        </div>
        <h1 v-if="messages.length===0">我们想探索什么</h1>

    </div>
    <form class="form1" @submit.prevent="ask">
        <textarea rows="2" placeholder="本模型由 deepseek-v4-pro 提供" v-model="input"></textarea>
        <div class="button1">
        <button class="button2" :disabled="asking">{{asking?"正在提问":"提交"}}</button>
        <button class="button3" type="button" @click="delete1">删除历史记录</button>
        </div>

    </form>


    
   </div>
</template>


<style scoped>
.chat h1{
    margin-top:37vh;
    font-style:italic;
    text-align:center;
}
.chat p{
    overflow-wrap:anywhere;
    white-space:pre-wrap;
}
*{
    box-sizing:border-box;
  
}
p{
    margin:0;
}
.p1{
    font-size:23px;
  margin-bottom:10px;
    line-height:1.5;
    margin-right:30px;
    margin-left:auto;
    background:rgb(191, 230, 191);
    width:fit-content;
    padding:5px 7px;
    border-radius:10px;
    max-width:90%;
}

   

.p2{
    font-size:23px;
   margin-right:auto;
    line-height:1.6;
    margin-left:30px;
    margin-bottom:10px;
    background:rgb(229, 205, 209);
     width:fit-content;
      padding:5px 5px;
       border-radius:10px;
       max-width:90%;
}
.p2 :deep(p) {
  margin: 0;
}

.content{
    z-index:0;
}
.chat{
    overflow:auto;
  height:79vh;
  
}
.div1{
   border: 2px solid rgb(238, 203, 209);
   
   height:87vh;
   position:relative;
   
  margin:0 100px;
  display:flex;
  flex-direction:column;
}
.form1{
    position:absolute;
    height:7vh;
    z-index:4;
    right:10px;
    left:10px;
    bottom:5px;
    margin:0 10px;
    display:flex;
    align-items:center;
    background:none;

}
textarea{
    font-size:20px;
    resize:none;
    width:69vw;
    border-radius:10px;
    line-height:1.5;
  
     background:none;
}
.button1 button{
    font-size:20px;
    border-radius:10%;
    margin-left:12px;
    border:none;
    color:rgb(241, 237, 237);
    font-weight:900;
}
.button2{
    background:rgb(168, 225, 168);
  padding:6px 10px;

}
.button3{
    background:rgb(239, 130, 130);
   padding:6px 10px;

}
.button2:hover{
    background:rgb(67, 228, 67);
   
}
.button3:hover{
    background:rgb(234, 48, 48);
}

</style>
