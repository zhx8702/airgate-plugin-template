import type { Config } from 'tailwindcss';
import { createPluginTailwindConfig } from '@airgate/theme/plugin';

const config: Config = {
  content: ['./src/**/*.{ts,tsx}'],
  ...createPluginTailwindConfig({
    scopeSelector: '[data-ag-template-root]',
  }),
};

export default config;
