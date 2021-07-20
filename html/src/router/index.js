// import Vue from 'vue'
// import VueRouter from 'vue-router'

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/views/home.vue')
  },
  {
    path: "/p/:url",
    name: "post",
    component: ()=>import("@/views/post.vue")
  },
  {
    path: "/t/:tag",
    name: "tag",
    component: ()=>import("@/views/tag.vue")
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '@/views/about.vue')
  },
  {
    path: "/archive",
    name: "archive",
    component: () => import("../views/archive.vue")
  },
  {
    path: "/message",
    name: "message",
    component: () => import("../views/message.vue")
  },
  {
    path: "/search",
    name: "search",
    component: () => import("../views/search.vue")
  },
  {
    path: "/overview",
    name: "overview",
    component: () => import("../views/overview.vue")
  },
  // 专栏
  {
    path: "/zhuanlan",
    name: "zhuanlan",
    component: () => import("../views/zhuanlan.vue")
  },
  {
    path: "/zhuanlan/:link",
    name: "zhuanlan_page",
    component: () => import("../views/zhuanlan_page.vue")
  },
  // 控制台
  {
    path: "/dashboard",
    name: "dashboard",
    component: () => import("../views/dashboard.vue")
  },
  // 404
  {
    path: "*",
    name: "404",
    component: () => import("../views/404.vue")
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router
