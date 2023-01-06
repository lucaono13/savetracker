<template>
  <!-- <div>
    <router-link class="link" to="/">Start</router-link> |
    <router-link class="link" to="/features">Feature Docs</router-link> |
    <router-link class="link" to="/home">Home</router-link>
  </div> -->
  <!-- <br /> -->
  <Menubar class="w-full" :model="items">
    <template #start>
      <OverlayPanel ref="op" :showCloseIcon="true" >
        
      </OverlayPanel>
    </template>
    <template #end>
      <p>Save Tracker</p>
    </template>
  </Menubar>
  <div class="grid w-full h-full">
    <div class="col-fixed h-full" v-if="sbVisible" style="width:205px">
      <Menu :model="$router.getRoutes()" class="fixed align-content-evenly">
        <template #item="{ item }">
          <!-- <br> -->
          <span :class="{ 'hidden': item.meta.secondary}">
          <font-awesome-icon :icon="item.meta.icon" />
          <!-- <i :class="item.meta.icon"></i> -->
          <router-link :to="item.path" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">            
            <a :href="href" @click="navigate" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }" style="color:darkturquoise">
              {{ route.name }}
            </a>
          </router-link>
          </span>
          <!-- <br> -->
          <!-- <hr> -->
        </template>
      </Menu>
    </div>
    <div class="col">
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
  // Below is checking if there are any saves in database
  // TODO: when page is loaded, check
  var saves: number
  let sidebarV = false
  if (localStorage.length == 0) {
    localStorage.setItem('saves', "0")
    saves = 0
  } else {
    saves = JSON.parse(localStorage.getItem('saves') || "0")
  }
  if (saves > 0) {
    sidebarV = true
  }
  export default {
    data() {
      return {
        sbVisible : sidebarV,
        items: [
          {
            label: "Test",
            icon: "pi pi-fw pi-file"
          }
        ]
      }
    }
  }

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
  padding: .75rem!important;
  // width: 100%!important;
  // width:calc(max-content+25px)!important;
  background: rgba(240, 248, 255, 0)!important;
  border: rgba(240, 248, 255, 0)!important;
}

.p-menu-list {
  justify-content:space-around!important;
}

.col-fixed {
  padding: 0!important;
}

.active-link-exact {
  color: fuchsia!important;
}

.p-menubar{
  background: rgba(240, 248, 255, 0)!important;
  border: rgba(240, 248, 255, 0)!important;
  // padding-right: 5px!important;
}

</style>
