// import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
// import App from './App.vue';
import SignUpOne from './components/SignUpOne.vue';
import SignUpTow from './components/SignUpTow.vue';
import LogInPsw from './components/LogInPsw.vue';
import LogInEmail1 from './components/LoginEmail1.vue';
import LogInEmail2 from './components/LoginEmail2.vue';
import HomeV from './components/HomeV.vue';
import RePasswordPsw from './components/RePasswordPsw.vue';
import RePasswordEmail1 from './components/RePasswordEmail1.vue';
import RePasswordEmail2 from './components/RePasswordEmail2.vue';


const routes = [
  { path: '/', redirect: '/signup1' },
  { path: '/signup1', component: SignUpOne },
  { path: '/signup2', component: SignUpTow },
  { path: '/LogInPsw', component: LogInPsw },
  { path: '/LogInEmail', component: LogInEmail1 },
  { path: '/LogInEmail2', component: LogInEmail2 },
  { path: '/HomeV', component: HomeV },
  { path: '/RePasswordPsw', component: RePasswordPsw },
  { path: '/RePasswordEmail1', component: RePasswordEmail1 },
  { path: '/RePasswordEmail2', component: RePasswordEmail2 }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});
export default router;