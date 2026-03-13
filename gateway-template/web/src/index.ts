import { AccountForm } from './components/AccountForm';
import type { AccountFormProps } from './components/AccountForm';
import { ensurePluginStyles } from './theme/runtime';

// 模块加载时注入主题和样式
ensurePluginStyles();

/** 插件前端模块导出 */
export interface PluginFrontendModule {
  /** 账号表单组件（嵌入核心的"添加账号"弹窗） */
  accountForm?: React.ComponentType<AccountFormProps>;
  // 如需注册独立页面，取消注释：
  // routes?: Array<{ path: string; component: React.ComponentType }>;
  // 如需添加侧边栏菜单，取消注释：
  // menuItems?: Array<{ key: string; label: string; icon?: React.ReactNode }>;
}

const plugin: PluginFrontendModule = {
  accountForm: AccountForm,
};

export default plugin;
