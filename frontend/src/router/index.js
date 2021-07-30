import { createRouter, createMemoryHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Overview from '../views/Overview.vue'
import Settings from '../views/Settings.vue'


const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/overview',
    name: 'Overview',
    // You can only use pre-loading to add routes, not the on-demand loading method.
    component: Overview
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
