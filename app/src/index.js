/** GLOBAL IMPORTS **/
import './assets/scss/index.scss';
import Vue from 'vue'
import App from './assets/templates/app.vue'
import Tone from 'tone'

// TODO: remove when in production
Vue.config.devtools = true;

/** INIT WEBSOCKET **/
let wsuri = "ws://127.0.0.1:8080";
let sock = new WebSocket(wsuri);
sock.onopen = () => { console.log(`connected to ${wsuri}`); }
sock.onclose = (e) => { console.log(`connection closed (${e.code})`); }

// ADD GLOBAL IMPORTS TO VUE
Object.defineProperty(Vue.prototype, '$tone', { value: Tone });
Object.defineProperty(Vue.prototype, '$sock', {value: sock });

/** INIT VUE **/
new Vue({
  el: '#app',
  render: h => h(App),
});




