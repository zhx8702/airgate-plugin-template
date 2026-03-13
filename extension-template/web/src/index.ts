import { ensurePluginStyles } from './theme/runtime';
import { HomePage } from './pages/HomePage';

// 初始化插件样式
ensurePluginStyles();

// 导出 PluginFrontendModule
export default {
  routes: [
    { path: '/', component: HomePage },
  ],
};
