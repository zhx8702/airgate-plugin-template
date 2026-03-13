import { useEffect, useState } from 'react';
import { useScopedPluginTheme, Section, Card, Badge, Button } from '@airgate/theme/plugin';
import { THEME_SCOPE_SELECTOR, STORAGE_KEY } from '../theme/runtime';

interface StatusResponse {
  status: string;
  plugin: string;
  version: string;
}

export function HomePage() {
  const rootRef = useScopedPluginTheme({ storageKey: STORAGE_KEY });
  const [status, setStatus] = useState<StatusResponse | null>(null);
  const [loading, setLoading] = useState(false);

  const fetchStatus = async () => {
    setLoading(true);
    try {
      const token = localStorage.getItem('token');
      const resp = await fetch('/api/v1/ext/ext-template/status', {
        headers: token ? { Authorization: `Bearer ${token}` } : {},
      });
      if (resp.ok) {
        setStatus(await resp.json());
      }
    } catch (err) {
      console.error('Failed to fetch status:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStatus();
  }, []);

  return (
    <div ref={rootRef} data-ag-ext-template-root="">
      <Section
        title="扩展插件示例"
        description="这是一个 Extension 类型插件的示例页面，展示了如何创建独立页面和自定义 API。"
      >
        <div className="agw-space-y-4">
          <Card>
            <div className="agw-flex agw-items-center agw-justify-between agw-mb-4">
              <h3 className="agw-text-sm agw-font-semibold agw-text-text">插件状态</h3>
              <Badge variant={status?.status === 'running' ? 'success' : 'neutral'}>
                {status?.status || '未知'}
              </Badge>
            </div>

            {status && (
              <div className="agw-space-y-2 agw-text-sm agw-text-text-secondary">
                <div>插件 ID: <span className="agw-text-text">{status.plugin}</span></div>
                <div>版本: <span className="agw-text-text">{status.version}</span></div>
              </div>
            )}

            <div className="agw-mt-4">
              <Button variant="secondary" onClick={fetchStatus} disabled={loading}>
                {loading ? '刷新中...' : '刷新状态'}
              </Button>
            </div>
          </Card>

          <Card>
            <h3 className="agw-text-sm agw-font-semibold agw-text-text agw-mb-2">开发指南</h3>
            <ul className="agw-space-y-1 agw-text-sm agw-text-text-secondary agw-list-disc agw-pl-4">
              <li>后端路由注册在 <code className="agw-text-text-tertiary">routes.go</code></li>
              <li>前端页面组件在 <code className="agw-text-text-tertiary">pages/</code> 目录</li>
              <li>API 路径格式：<code className="agw-text-text-tertiary">/api/v1/ext/ext-template/*</code></li>
              <li>使用 SDK 组件保持与 Core 统一的视觉风格</li>
            </ul>
          </Card>
        </div>
      </Section>
    </div>
  );
}
