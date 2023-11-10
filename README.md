# go-vue 前后端分离参考示例

## 跨域

由于浏览器的同源策略，向不同源（协议，域名，端口）发送 ajax 请求会失败。  
所以，页面请求"localhost:80"页面，ajax 请求"localhost:9090"接口会失败。  
所以，需要同源策略请求接口，需要配置前端代理，将请求发送给"localhost:80 -> localhost:9090"实现跨域请求。

使用 Header 允许跨越，允许来自任何源的请求访问资源，这样确实可以解决跨域访问的问题。然而，这也意味着所有的网站都可以访问你的资源，存在安全风险，可以指定允许源，但是也有风险。

```
func main() {
	http.HandleFunc("/list", sayhelloName)   //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有源跨域请求，生产环境中应谨慎使用
	w.Write("Hello world")
}
```

配置前端代理(vite.config.js，工程化开发时使用 vite)：

```
需要接口地址以"/api"开头
export default {
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:9090', // 接口地址
        changeOrigin: true,  // 修改源
        rewrite: (path) => path.replace(/^\/api/, '')  // api替换为''
      }
    }
  }
}
```

## vue 工程化环境

```
npm install -g @vue/cli
vue create my-vue-app
cd my-vue-app
npm run serve
```

## vue 路由和子路由

```
const routes = [
  { path: '/', component: Home },
  {
    path: '/products',
    component: Products,
    children: [
      { path: 'shoes', component: Shoes },
      { path: 'clothing', component: Clothing }
    ]
  }
]
```

## axios 响应拦截器，对结果做处理并返回

```
import axios from 'axios';

// 创建 Axios 实例
const instance = axios.create({
  baseURL: 'https://api.example.com'
});

// 添加响应拦截器
instance.interceptors.response.use(
  function (response) {
    // 对响应数据做点什么
    return response;
  },
  function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

// 发起请求
instance.get('/user/123')
  .then(function (response) {
    // 处理成功的响应
    console.log(response.data);
  })
  .catch(function (error) {
    // 处理错误的响应
    console.log(error);
  });
```

## pinia 状态管理器，管理 token 信息

```
首先，安装 Pinia：
npm install pinia

然后，在你的 main.js 文件中创建一个 Pinia 实例并将其挂载到应用程序上：
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.mount('#app')

接下来，在你的组件中定义一个 Pinia store：
// store.js
import { defineStore } from 'pinia'

export const useCounterStore = defineStore({
  id: 'counter',
  state: () => ({
    count: 0
  }),
  actions: {
    increment() {
      this.count++
    }
  }
})

然后，在你的组件中使用这个 store：
<script setup>
import { useCounterStore } from './store'

const counterStore = useCounterStore()
</script>

<template>
  <div>
    <p>Count: {{ counterStore.count }}</p>
    <button @click="counterStore.increment">Increment</button>
  </div>
</template>
```

### 当使用 Axios 发送请求时，可以通过请求拦截器在发送请求前对请求进行处理，例如在请求头中添加 token 信息。同时，我们可以结合 Pinia 状态管理器来获取 token 信息。

```
// store.js
import { defineStore } from 'pinia'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    token: null
  }),
  actions: {
    setToken(token) {
      this.token = token
    },
    getToken() {
      return this.token
    }
  }
})
```

```
// axios.js
import axios from 'axios'
import { useAuthStore } from './store'

const instance = axios.create({
  baseURL: 'https://api.example.com'
});

// 添加请求拦截器
instance.interceptors.request.use(
  function (config) {
    // 在发送请求之前做些什么
    const token = useAuthStore().getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

export default instance;
```

## pinia 持久化插件 persist

pinia 默认是内存存储，刷新浏览器数据丢失，persist 可以将 pinia 数据持久化存储

```
npm install pinia vuex-persistedstate
```

```
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createPersistedState } from 'vuex-persistedstate'
import App from './App.vue'

const app = createApp(App)
const pinia = createPinia()

pinia.use(createPersistedState())

app.use(pinia)
app.mount('#app')
```

```
import { defineStore } from 'pinia'

export const useCounterStore = defineStore({
  id: 'counter',
  state: () => ({
    count: 0
  }),
  persist: true // 将状态持久化
})
```

## 将 Vue 项目打包部署上线并运行需要经过以下步骤

打包 Vue 项目： 在 Vue 项目目录中，你可以使用以下命令来打包项目：

```
npm run build
```

这将会生成一个名为 dist 的目录，其中包含了打包后的静态文件，包括 HTML、CSS、JavaScript 等。

部署静态文件： 将 dist 目录中的静态文件部署到你的服务器或者静态文件托管服务中。你可以使用 FTP、SCP、rsync 等工具将文件上传到服务器，或者使用云存储服务（如 AWS S3、Google Cloud Storage）来托管静态文件。

配置服务器： 确保你的服务器已经安装了适当的 Web 服务器软件（如 Nginx、Apache）并进行了正确的配置，以便能够正确地提供静态文件。你需要配置服务器以将用户的请求定向到你部署的静态文件目录 (一般都是 index.html 为入口文件)。

启动服务器： 一旦静态文件部署到服务器上，你可以启动服务器，并确保你的 Vue 项目能够通过公开的 URL 进行访问。你可以使用 Nginx、Apache 或其他 Web 服务器软件来启动服务器。
