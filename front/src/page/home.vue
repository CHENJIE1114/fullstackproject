<script setup lang="ts">
import { RouterLink, useRouter } from "vue-router";
import { quit } from "../hanshu";
const router = useRouter();

async function logout() {
  const result = await quit();
  if (result) {
    router.push("/log")
  }
}
</script>

<template>
  <div class="div1">
    <div class="star1">
      <div class="tail1"></div>
    </div>

     <div class="star2">
      <div class="tail2"></div>
    </div>
    <div class="star3">
      <div class="tail3"></div>
    </div>

    


  <div class="choice">
    <RouterLink to="/crud" class="rou">编辑笔记</RouterLink>
    <RouterLink to="/aic" class="rou">ai对话</RouterLink>
    <button @click="logout">退出登录</button>
  </div>
  
  </div>
</template>
<style scoped>


.rou{
  background:none;
}
.choice{
  width:fit-content;
  margin:400px auto;
  z-index:2;
  display:flex;
 
  gap:30px;
}
.star1{
  position:absolute;
  right:20vw;
  top:14vh;
  width:12px;
  height:12px;
  border-radius:50%;
  background:white;
  box-shadow:0 0 10px 5px white;
  animation:move 2s linear infinite;
  z-index:0;
}
.tail1{
  position:absolute;
  left:4px;
  top:4px;
  width:304px;
  height:8px;
  border:none;
  background:linear-gradient(
    to right,
    white  0%,
    transparent 100%
  );
  transform-origin:left center;
  transform:rotate(-45deg);
}
@keyframes move{
  from{
    transform:translate(0);
opacity:1;
    }
    to{
      transform:translate(-90vw,90vh);
      opacity:0.1;
    }
}

.star2{
  position:absolute;
  right:26vw;
  top:45vh;
  width:8px;
  height:8px;
  border-radius:50%;
  background:white;
  box-shadow:0 0 10px 5px white;  
  opacity:0;
  animation:move2 2s linear infinite;
  animation-delay:1s;
 z-index:0;
}
.tail2{
  position:absolute;
  left:4px;
  top:4px;
  width:179px;
  height:2px;
  border:none;
  background:linear-gradient(
    to right,
    white  0%,
    transparent 100%
  );
  transform-origin:left center;
  transform:rotate(-45deg);
}
@keyframes move2{
  0%{
    transform:translate(0);
    opacity:1
    }
    85%{
      transform:translate(-90vw,90vh);
         opacity:1
    }
    100%{
      transform:translate(-90vw,90vh);
         opacity:0.3
    }
}

.star3{
  position:absolute;
  left:84vw;
  top:80vh;
  width:8px;
  height:8px;
  border-radius:50%;
  background:white;
  box-shadow:0 0 10px 5px white;
  animation:move3 1.8s linear infinite;
   z-index:0;
}
.tail3{
  position:absolute;
  left:4px;
  top:4px;
  width:179px;
  height:2px;
  border:none;
  background:linear-gradient(
    to right,
    white  0%,
    transparent 100%
  );
  transform-origin:left center;
  transform:rotate(-45deg);
}
@keyframes move3{
  from{
    transform:translate(0);

    }
    to{
      transform:translate(-90vw,67vh);
    }
}
.div1 {
  position: relative;
  min-height: 100vh;
  box-sizing: border-box;
  overflow: hidden;

  background:rgb(117, 188, 218);
}


button {
  width: fit-content;
  margin: auto ;
  display: block;
  padding: 10px 22px;

  color: #a34d58;
  background: none;

  border: 1px solid #e7bec4;
  border-radius: 9px;
}
button:hover {
  color: white;
  background: #b04242;

  transform: translateY(-5px);
}
.rou {
  padding: 12px 24px;

  color: #3f596a;
  text-decoration: none;

  background:none;

  border: 1px solid #bfd4e0;
  border-radius: 10px;

  transition: 0.5s;
}

.choice a:hover {
  color: white;
  background: #bda8a8;

  transform: translateY(-5px);
}
.div1{
z-index:0;
background-image:url("../image/b.png");
background-size:cover;
background-position: center;
}
</style>
