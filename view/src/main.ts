import 'ant-design-vue/dist/reset.css';

import { createApp } from 'vue'
import Antd from 'ant-design-vue'
import {store} from '@/pinia'
import App from '@/App.vue'
import router from '@/router'
import '@/permission'


const app = createApp(App)

app.use(router)
app.use(Antd)
app.use(store)

app.mount('#app')
