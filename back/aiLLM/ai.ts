import express,{type Request,type Response} from "express"
import {HumanMessage,AIMessage} from "@langchain/core/messages"
import {ChatOpenAI} from "@langchain/openai"
import "dotenv/config"

const app=express()
app.use(express.json())
function createmodel():ChatOpenAI{
const apiKey=process.env.DEEPSEEK_API_KEY
if(!apiKey){
    throw new Error("DEEPSEEK_API_KEY is not configured")
}
return new ChatOpenAI({
    apiKey,
    model:process.env.DEEPSEEK_MODEL??"deepseek-v4-pro",
    temperature:0.7,
    configuration:{
        baseURL:process.env.DEEPSEEK_BASE_URL??"https://api.deepseek.com",
    },
})
}
type message={
    role:string,
    content:string,
}
app.post("/ai/llm",async function(request:Request,response:Response){
const allay=request.body.data as message[]
const cfmessage =allay.map((item)=>{
    if (item.role=="user"){
        return new HumanMessage({
            content:item.content,
        })
    }else{
        return new AIMessage({
            content:item.content,
        })
    }
})
   try{ const model=createmodel()
    const stream=await model.stream(cfmessage)
    response.setHeader("Content-Type","text/plain;charset=utf-8")
    response.setHeader("Cache-Control","no-cache")
    for await(const chunk of stream){
        response.write(chunk.content)
    }
    response.end()
}catch(err){
    console.log("err in backendai")
     response.status(500)
     response.end()
}



})
const port=Number(process.env.AI_PORT??"3000")
const host=process.env.AI_HOST??"127.0.0.1"
app.listen(port,host)
