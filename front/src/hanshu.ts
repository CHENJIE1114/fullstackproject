export async function islog(): Promise<boolean> {
    
    try {
        const response = await fetch("/api/islog", {
            credentials: "include",
        });
        if (!response.ok) {
            return false;
        } else {
            return true;
        }
    } catch (err) {
        return false;
    }
}

export async function quit() :Promise<boolean>{
   
    try {
        const response = await fetch("/api/logout",{method:"POST"});
        if (!response.ok){
           
            return false

            
        }
     return true
        
    } catch(err) {
        return false
     }
}

export interface notes{
    content :string
    title: string
    createdtime: string
    updatetime: string
    userid :number
    cos:number
    id:number
   
}
export async function load(ques :string):Promise<notes[]> {

    const response=await fetch("/api/load",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify({
            question:ques,
        })
    })
    if(!response.ok){
        throw new Error("err in load notes")
    }
    const notes=await response.json()

    return notes.data
}

export async function addnote(text :string,content:string) {
    const response=await fetch("/api/add",{
        method:"POST",
        headers:{"Content-Type":"application/json"},
        body:JSON.stringify({
            text:text,
            content:content,
        })
    })
    if(!response.ok){
        throw new Error("err in add note")
    }

}

export function formatetime(time :string):string{
    if (time==""){
        return ""

    }
   return new Date(time).toLocaleString("zh-CN",{hour12:false})

}