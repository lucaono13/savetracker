<template>
  <div>
    <router-link class="link" to="/">Start</router-link> |
    <router-link class="link" to="/features">Feature Docs</router-link> |
    <router-link class="link" to="/home">Home</router-link>
  </div>
  <br>
  <div class="grid" style="width:100%">
    <div class="col-1" v-if="enough" >
      <ul>
        <li>Item 1</li>
        <li>Item 2</li>
        <li>Item 3</li>
      </ul>
    </div>
    <div class="col-11">

      <router-view />
    </div>
    <div class="col">
      <Button type="button" label="visibility" v-on:click="enough = !enough" />
    </div>"
  </div>
  <!-- <div class="dock-window">
    <Dock :model="items" position="left">
      <template #icon="{ item }">
        <img :alt="item.label" :src="item.icon"/>
      </template>
    </Dock>
  </div> -->
  <!-- <router-view /> -->
</template>

<script setup lang="ts">
  function showSideBar() {
    enough = !enough
  }

  const items = [
    {
      label: "First",
      icon: "pi-hashtag"
    },
    {
      label: "Second",
      icon: "pi-server",
    },
  ]
  let enough = false
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
</script>

<style lang="scss">
@font-face {
  font-family: "Nunito";
  font-style: normal;
  font-weight: 400;
  src: local(""),
  url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
}

html {
  background-color: rgba(33, 37, 43, 1);
  color: white;
  font-family: "Nunito", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
  "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
  sans-serif;
}

body {
  margin: 0;
}

h3 {
  margin: 0;
}

#app {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.link {
  cursor: pointer;
  color: #df0000;

  &:hover {
    text-decoration: underline;
  }
}
</style>
