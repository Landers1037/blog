import Vue from 'vue'
import VueRouter from 'vue-router'
import home from "../views/home";
import post from "../views/post";
import tag from "../views/tag";
import about from "../views/about";
import archive from "../views/archive";
import message from "../views/message";
import search from "../views/search";
import overview from "../views/overview";

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: home
  },
  {
    path: "/p/:url",
    name: "post",
    component: post
  },
  {
    path: "/t/:tag",
    name: "tag",
    component: tag
  },
  {
    path: '/about',
    name: 'about',
    component: about
  },
  {
    path: "/archive",
    name: "archive",
    component: archive
  },
  {
    path: "/message",
    name: "message",
    component: message
  },
  {
    path: "/search",
    name: "search",
    component: search
  },
  {
    path: "/overview",
    name: "overview",
    component: overview
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router
