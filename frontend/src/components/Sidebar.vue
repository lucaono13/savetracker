<template>
    <div class="col-fixed h-full" style="width:205px">
      <Menu :model="$router.getRoutes()" class="fixed align-content-evenly">
        <template #item="{ item }">
          <!-- <br> -->
          <span :class="{ 'hidden': item.meta.secondary}">
          <font-awesome-icon :icon="item.meta.icon" />
          <router-link :to="item.path" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">            
            <a :href="href" @click="navigate" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }" style="color:darkturquoise">
              {{ route.name }}
            </a>
          </router-link>
          </span>
        </template>
      </Menu>
    </div>
</template>
