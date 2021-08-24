<template>
  <div class="wrapper" :class="{ 'nav-open': $sidebar.showSidebar }">
    <side-bar
      :background-color="sidebarBackground"
      short-title="NodeHubs"
      title="NodeHubs"
    >
      <template v-slot:links>
        <sidebar-item
          :link="{
            name: '概览',
            icon: 'ni ni-tv-2 text-primary',
            path: '/dashboard',
          }"
        />

        <sidebar-item
          :link="{
            name: '节点监控',
            icon: 'ni ni-sound-wave text-green',
            path: '/nodes',
          }"
        />
        <sidebar-item
          :link="{
            name: '服务监控',
            icon: 'ni ni-app text-orange',
            path: '/services',
          }"
        />
      </template>
    </side-bar>
    <div class="main-content" :data="sidebarBackground">
      <div @click="toggleSidebar">
        <router-view></router-view>
        <content-footer v-if="!$route.meta.hideFooter"></content-footer>
      </div>
    </div>
  </div>
</template>
<script>
import ContentFooter from "./ContentFooter.vue";

export default {
  components: {
    ContentFooter,
  },
  data() {
    return {
      sidebarBackground: "vue",
    };
  },
  methods: {
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    },
  },
};
</script>
<style lang="scss"></style>
