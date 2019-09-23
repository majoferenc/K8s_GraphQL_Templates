import Vue from "vue";
import Router from "vue-router";
import HelloWorld from "./components/HelloWorld";

Vue.use(Router);

const router = new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    {
      path: "/",
      name: "helloworld",
      component: HelloWorld,
      props: true
    }
  ]
});

export default router;
