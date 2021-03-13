import Vue from 'vue'
import VueRouter from 'vue-router'
// ng版本为离线版本 抛弃所有cdn和懒加载服务
import home from "../views/home";
import about from "../views/about";
import post from "../views/post";
import archive from "../views/archive";
import tag from "../views/tag";
import message from "../views/message";
import search from "../views/search";
import overview from "../views/overview";
import weindex from "../views/weui/index";
import wepost from "../views/weui/post"
import article from "../views/weui/article";
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
  },
    //weui
  {
    path: "/newui",
    name: "newui",
    component: weindex
  },
  {
    path: "/newui/p/:url",
    name: "newui_post",
    component: wepost
  },
  {
    path: "/newui/article",
    name: "newui_article",
    component: article
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router
