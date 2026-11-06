<script setup lang="ts">
import {ref,onMounted} from "vue"
import {load,type notes,addnote,formatetime} from "../hanshu"
const asking=ref(false)
const working=ref(false)
const search=ref("")
const title=ref("")
const content=ref("")
const id=ref(0)
const notes=ref<notes[]>([])
async function searchfunc(){
asking.value=true
    try {
        notes.value=await load(search.value)
    }catch(err){
        console.log(err instanceof Error ?err.message:"err in log")
    }finally{
asking.value=false
    }

}

async function add(){
working.value=true
    try {
       await addnote(title.value,content.value)
       searchfunc()
         title.value=""
        content.value=""
    }catch(err){
        console.log(err instanceof Error ?err.message:"err in log")
    }finally{
working.value=false
    }

}

function edit(title1 :string,content1:string,id1:number){
    title.value=title1
        content.value=content1
        id.value=id1
}
function cancel(){
    title.value=""
        content.value=""
        id.value=0
}
async function sendedit(title1 :string,content1:string,id1:number){
    try{
        
        const response=await fetch("/api/edit",{
            method:"PUT",
            headers:{"Content-Type":"application/json"},
            body:JSON.stringify({
                id:id1,
                content:content1,
                title:title1,
            })
        })
        if(!response.ok){
            throw new Error("err in edit")
        }
        searchfunc()
    }catch(err){

    }finally{
        title.value=""
        content.value=""
        id.value=0;
    }
}

async function deletenote(id1:number){
    const confirm=window.confirm("确定删除吗")
    if (!confirm){
        return
    }
     try{
        
        const response=await fetch(`/api/delete/${id1}`,{
            method:"PUT",
          
        })
        if(!response.ok){
            throw new Error("err in delete")
        }
        searchfunc()
    }catch(err){
        console.log("err in delete")
    }
   

}

onMounted(searchfunc)
</script>
<template>
     <div class="div1">
        <section class="secta">
            <div class="adiv1">
        <p>输入你想搜索的笔记</p>
           <form @submit.prevent="searchfunc">
            <div class="divform">
            <textarea placeholder="例如:解释vue框架" rows="2"
            class="input1" v-model="search"></textarea>
            <button :disabled="asking" >{{ asking?"正在回答":"提问" }}</button>
            </div>

           </form> 
           
           </div>

           <div class="adiv2">
            <p>{{ id?"修改笔记":"新加笔记" }}</p>
            <form class="a2form" @submit.prevent="add">
                <label class="a2l">标题<input required v-model="title"></label>
                
                <label for="diva2">内容</label>
                <textarea rows="7" required  id="diva2" v-model="content"></textarea>
                
               <div class="adiv22" v-if="!id"><button :disabled="working">{{id?"保存":"新加"}}</button></div>
               <div class="buttonedit" v-else>
                    <button type="button" @click="sendedit(title,content,id)">保存</button>
                    <button type="button" @click="cancel">取消</button>
               </div>
            </form>
           </div>
           
        </section>
        <section class="sectb">
            <div class="title">
            <h2>全栈笔记</h2><h3>随心编写</h3>
            </div>
            <div class="bcontent">

                
                <div v-for="value,index in notes" class="bfor">
                    <div class="message">
                     <span>标题 : {{ value.title}} </span>  
                     <span>用户id : {{value.userid}}</span>
                     <span>问题的余弦相似度 : {{ value.cos}}</span>
                    </div>
                    <p class="forp">{{ value.content }}</p>
                    <div class="contime">
                     <span>创建时间 : {{ formatetime(value.createdtime)}} </span>  
                     <span>更新时间 : {{formatetime(value.updatetime)}}</span>
                     
                    </div>

                    <div class="conbutton">
                        <button @click="edit(value.title,value.content,value.id)">编辑</button>
                        <button @click="deletenote(value.id)">删除</button>
                    </div>

                </div>

            </div>
        </section >
     </div>




</template>
<style scoped>
.contime{
display:flex;
gap:30px;

}
.buttonedit{
    display:flex;
    width:20vw;
    gap:14px;
    justify-content:flex-end;
    margin-top:10px;
}
.conbutton{
    display:flex;
    gap:14px;
    justify-content:flex-end;
}
.bfor{
    margin-bottom:20px;
    border-bottom:1px solid rgb(185, 230, 185);
}
.forp{
    white-space:pre-wrap;
    overflow-wrap:anywhere;

}
.message{
    display:flex;
    padding-bottom:10px;
    gap:29px;
}
.bcontent{
    margin:20px;
   
    overflow:auto;
    overscroll-behavior:contain;
    height:47vh;
}
.bcontent::-webkit-scrollbar {
  display: none;
}

.div1{
    display:flex;
    min-height:87vh;
    justify-content:center;
    gap:50px;
    align-items:center;
    background:linear-gradient(
        45deg,
        rgb(134, 228, 134), 
        rgb(175, 190, 235),
        rgb(240, 180, 180)  
    )
}
.secta p{
    text-align:center;
}
.secta{
    width:25vw;
    
    height:60vh;
}
.adiv1{
 border:2px solid rgb(237, 237, 235);
 max-height:20vh;
 margin-bottom:10vh;
 overflow:hidden;
}
.adiv2{
    height:35vh;
 border:2px solid rgb(237, 237, 235);
}
.adiv2 p{
    margin:20px 0;;
}
 .a2form
 {
        display:flex;
            flex-direction:column;
            justify-content:center;
            align-items:center;
           }
           .a2l{
            display:block;
           }
.a2form label{
    display:block;
}
.a2l{
    margin-bottom:20px;
}
.a2form input ,.a2form textarea{
    background:none;
    border:2px solid white;
    font-size:18px;
}
.a2form textarea{
    width:20vw;
    resize:none;
} 
.adiv22{
     width:20vw;
    }
 .adiv22 button{
  margin-left:auto;
  margin-top:20px;
 }

p{
    margin:0;
   
}
.input1{
    background:none;
    border:none;
    font-size:22px;
    resize:vertical;
    display:block;
    font-size:18px;
    width:18vw;
    max-height:13vh;
    margin-bottom:20px;
}
.divform{
    margin:auto;
    display:flex;
    justify-content:center;
    align-items:center;
    min-height:8vh;
    gap:5px;
}
button{
    display:block;
    font-size:18px;
    background:none;
    border:1px solid white;
}
button:hover{
    background:rgb(173, 235, 173);
}
  .title {
    display:flex;
    justify-content:center;
    gap:14px;
    margin:20px 20px;
    padding:0 0 14px;
    border-bottom:2px dashed rgb(160, 99, 212);
    align-items:baseline;
  }  
  .title h2,.title h3 {
    margin: 0;
}
.sectb{
    width:45vw;
    border:2px solid rgb(239, 240, 236);
    height:60vh;
}
</style>
