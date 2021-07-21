// import Vue from 'vue'
// import VueRouter from 'vue-router'

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    meta: {
      title: "主页 . Blog"
    },
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
    meta: {
      title: "标签 . Blog"
    },
    component: ()=>import("@/views/tag.vue")
  },
  {
    path: '/about',
    name: 'about',
    meta: {
      title: "关于 . Blog"
    },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '@/views/about.vue')
  },
  {
    path: "/archive",
    name: "archive",
    meta: {
      title: "归档 . Blog"
    },
    component: () => import("../views/archive.vue")
  },
  {
    path: "/message",
    name: "message",
    meta: {
      title: "留言 . Blog"
    },
    component: () => import("../views/message.vue")
  },
  {
    path: "/search",
    name: "search",
    meta: {
      title: "搜索 . Blog"
    },
    component: () => import("../views/search.vue")
  },
  {
    path: "/overview",
    name: "overview",
    meta: {
      title: "概览 . Blog"
    },
    component: () => import("../views/overview.vue")
  },
  // 专栏
  {
    path: "/zhuanlan",
    name: "zhuanlan",
    meta: {
      title: "专栏 . Blog"
    },
    component: () => import("../views/zhuanlan.vue")
  },
  {
    path: "/zhuanlan/:link",
    name: "zhuanlan_page",
    meta: {
      title: "专栏 . Blog"
    },
    component: () => import("../views/zhuanlan_page.vue")
  },
  // 控制台
  {
    path: "/dashboard",
    name: "dashboard",
    meta: {
      title: "管理 . Blog"
    },
    component: () => import("../views/dashboard.vue")
  },
  // 404
  {
    path: "*",
    name: "404",
    meta: {
      title: "404 . Blog"
    },
    component: () => import("../views/404.vue")
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  if (document.title === process.env.VUE_APP_TITLE || to.meta.title && to.meta.title !== document.title) {
    document.title = to.meta.title ? to.meta.title : process.env.VUE_APP_TITLE
  }
  next();
})

export default router
