<script setup lang="ts">
import { RouterLink, RouterView, useRoute } from "vue-router"
import { computed } from "vue"

const route = useRoute()

const navigation = computed(() => {
  return route.meta.navigation === true
})
</script>

<template>
  <div>
    <header v-if="navigation" class="top-nav">
      <div class="nav-title">
        <p>REI 学习中心</p>
        <span>个人知识与 AI 工作台</span>
      </div>

      <nav class="nav-links">
        <RouterLink class="nav-link" to="/home">
          主页
        </RouterLink>

        <RouterLink class="nav-link" to="/crud">
          编辑笔记
        </RouterLink>

        <RouterLink class="nav-link" to="/aic">
          AI 对话
        </RouterLink>
      </nav>
    </header>

    <section class="content" :class="{ 'has-nav': navigation }">
      <RouterView />
    </section>
  </div>
</template>

<style scoped>
/* 整个顶部导航栏 */
.top-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;

  height: 129px;
  box-sizing: border-box;

  background: #f8fbfd;
  border-bottom: 1px solid #7abfee;


  z-index: 100;
}

/* 上面的标题部分 */
.nav-title {
  height: 56px;

  display: flex;
  align-items: center;
  justify-content: center;
  gap: 14px;

  background: linear-gradient(
    90deg,
    #e8f7dc,
    #eef5f9,
    #f3e2e7
  );
}

/* 清除 p 标签自带的上下空白 */
.nav-title p {
  margin: 0;

  color: #294d65;
  font-size: 23px;
  font-weight: 600;
}

/* 标题旁边的小字 */
.nav-title span {
  color: #718895;
  font-size: 14px;
}

/* 四个按钮所在区域 */
.nav-links {
  height: 72px;
  box-sizing: border-box;

  display: flex;
  align-items: center;
  justify-content: space-evenly;

  padding: 0 30px;
}

/* 普通按钮 */
.nav-link {
  width: 135px;
  box-sizing: border-box;

  padding: 9px 16px;

  color: #315c75;
  background: #eaf3f8;
  border: 1px solid #b8cedc;
  border-radius: 10px;

  text-align: center;
  text-decoration: none;

  transition:
    background 0.2s,
    color 0.2s,
    transform 0.2s,
    box-shadow 0.2s;
}

/* 鼠标移动到按钮上 */
.nav-link:hover {
  color: #714f5a;
  background: #f2dfe4;
  border-color: #d7aeb8;

  transform: translateY(-2px);
  box-shadow: 0 4px 10px rgba(65, 88, 105, 0.15);
}

/* 当前所在页面对应的按钮 */
.nav-link.router-link-active {
  color: white;
  background: #5d8eac;
  border-color: #5d8eac;
}

/* 给正文留出导航栏的高度 */
.has-nav {
  padding-top: 128px;
}
</style>
