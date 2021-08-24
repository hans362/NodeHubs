import { createRouter, createWebHashHistory } from "vue-router";

import DashboardLayout from "@/layout/DashboardLayout";

import Dashboard from "../views/Dashboard.vue";
import Nodes from "../views/Nodes.vue";
import Services from "../views/Services.vue";

const routes = [
  {
    path: "/",
    redirect: "/dashboard",
    component: DashboardLayout,
    children: [
      {
        path: "/dashboard",
        name: "dashboard",
        components: { default: Dashboard },
      },
      {
        path: "/nodes",
        name: "nodes",
        components: { default: Nodes },
      },
      {
        path: "/services",
        name: "services",
        components: { default: Services },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  linkActiveClass: "active",
  routes,
});

export default router;
