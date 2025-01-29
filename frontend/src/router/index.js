import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/HomeView.vue'
// import Settings from '../views/Settings.vue'

const routes = [
  { path: '/', name: 'Home', component: Home },
  // { path: '/settings', name: 'Settings', component: Settings },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
