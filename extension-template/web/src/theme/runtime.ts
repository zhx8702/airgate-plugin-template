import { ensurePluginStyleFoundation } from '@airgate/theme/plugin';
import tailwindCssText from '../styles/tailwind.css?inline';

// ⚠️ 替换 "ext-template" 为你的插件标识
export const THEME_SCOPE_SELECTOR = '[data-ag-ext-template-root]';
export const THEME_ATTRIBUTE = 'data-theme';
export const STYLE_ID = 'ag-ext-template-theme-vars';
export const FOUNDATION_STYLE_ID = 'ag-ext-template-plugin-foundation';
export const TAILWIND_STYLE_ID = 'ag-ext-template-tailwind';
export const STORAGE_KEY = 'ag-ext-template-theme';

export function ensurePluginStyles(): void {
  ensurePluginStyleFoundation({
    scopeSelector: THEME_SCOPE_SELECTOR,
    themeAttribute: THEME_ATTRIBUTE,
    storageKey: STORAGE_KEY,
    themeStyleId: STYLE_ID,
    foundationStyleId: FOUNDATION_STYLE_ID,
    extraCssText: tailwindCssText,
    extraStyleId: TAILWIND_STYLE_ID,
  });
}
