import { useCallback } from 'react';
import {
  Field,
  SecretInput,
  TextInput,
  Section,
  useScopedPluginTheme,
} from '@airgate/theme/plugin';

/** 账号表单 Props（由核心 AccountsPage 注入） */
export interface AccountFormProps {
  credentials: Record<string, string>;
  onChange: (credentials: Record<string, string>) => void;
  mode: 'create' | 'edit';
  accountType?: string;
  onAccountTypeChange?: (type: string) => void;
  onSuggestedName?: (name: string) => void;
}

// ⚠️ 替换 "template" 为你的插件标识
const THEME_ATTRIBUTE = 'data-theme';
const STORAGE_KEY = 'ag-template-theme';

function usePluginScopedTheme<T extends HTMLElement>() {
  return useScopedPluginTheme<T>({
    themeAttribute: THEME_ATTRIBUTE,
    storageKey: STORAGE_KEY,
  });
}

export function AccountForm({ credentials, onChange, mode }: AccountFormProps) {
  const rootRef = usePluginScopedTheme<HTMLDivElement>();

  const updateField = useCallback(
    (key: string, value: string) => {
      onChange({ ...credentials, [key]: value });
    },
    [credentials, onChange],
  );

  return (
    <div ref={rootRef} data-ag-template-root className="agw-form-shell">
      <Section
        title="API 凭证"
        description="输入你的 API Key 和服务地址以接入上游 API。"
        eyebrow="认证信息"
        panel
      >
        <Field label="API Key" required hint="在服务商控制台获取的 API 密钥">
          <SecretInput
            value={credentials.api_key || ''}
            onChange={(e) => updateField('api_key', e.target.value)}
            placeholder="sk-..."
            disabled={mode === 'edit'}
          />
        </Field>

        <Field label="API 地址" hint="留空使用默认地址">
          <TextInput
            value={credentials.base_url || ''}
            onChange={(e) => updateField('base_url', e.target.value)}
            placeholder="https://api.example.com"
          />
        </Field>
      </Section>
    </div>
  );
}
