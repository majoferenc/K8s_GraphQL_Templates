import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from './components/HelloWorld'
import ApolloExample from './components/ApolloExample'

Vue.use(Router)

const router = new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    {
      path: '/',
      name: 'helloworld',
      component: HelloWorld,
      props: true
    },
    {
      path: '/apopplo',
      name: 'apollo',
      component: ApolloExample,
      props: false
    }
  ]
})

export default router
