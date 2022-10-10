<template>
  <div>
    <router-link class="link" to="/">Start</router-link> |
    <router-link class="link" to="/features">Feature Docs</router-link> |
    <router-link class="link" to="/home">Home</router-link>
  </div>
  <br />
  <!-- <div class="absolute top-0 left-0">
    <Menu  :model="items" />
  </div>
  <div>
    <router-view />
  </div> -->
  <div class="grid" style="width: 100%">
    <div class="col-2" >
      <Menu :model="$router.getRoutes()" class="fixed">
        <template #item="{ item }">
          <span :class="{ 'hidden': item.meta.secondary}">
          <i :class="item.meta.icon"></i>
          <router-link :to="item.path" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">            
            <a :href="href" @click="navigate" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }">
              {{ route.name }}
            </a>
          </router-link>
          </span>
        </template>
      </Menu>
    </div>
    <div class="col-10">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
</script>

<style lang="scss">
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 400;
  src: local(''), url('assets/fonts/nunito-v16-latin-regular.woff2') format('woff2');
}

html {
  background-color: rgba(33, 37, 43, 1);
  color: white;
  font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell',
    'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
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

.p-menu {
  height: 100%!important;
  width: 100%!important;
  // width:calc(max-content+25px)!important;
  // background: aliceblue!important;
}

.p-menu-list {
  justify-content:space-around!important;
}

</style>
