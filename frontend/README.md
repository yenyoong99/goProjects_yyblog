# yyblog

### new tb npm source
```
$ yrm add tn https://registry.npmmirror.com

$ yrm use tn
```

### Install Arco Design Library
```
# npm
npm install --save-dev @arco-design/web-vue

# yarn
yarn add --dev @arco-design/web-vue
```

main.js file import library
```js
import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue';
import App from './App.vue';
import '@arco-design/web-vue/dist/arco.css';

const app = createApp(App);
app.use(ArcoVue);
app.mount('#app');
```
