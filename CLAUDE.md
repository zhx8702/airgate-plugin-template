# AirGate 插件开发上下文

本文件为 AI 辅助开发提供完整上下文。开发新插件时，AI 应读取此文件了解项目架构和规范。

## 项目概述

本仓库包含 **两种 AirGate 插件模板**：

| 目录 | 类型 | 用途 |
|---|---|---|
| `gateway-template/` | Gateway 网关插件 | 转发 API 请求到上游服务（如 OpenAI、Claude） |
| `extension-template/` | Extension 扩展插件 | 独立页面、自定义 API、后台任务（如支付、监控） |

架构均为 Go gRPC 后端 + React 前端：

- **后端**：Go 进程，通过 gRPC 与 AirGate Core 通信（基于 hashicorp/go-plugin）
- **前端**：React 组件，编译为 ES Module，由 Core 前端动态加载
- **SDK**：`@airgate/theme`（前端）和 `github.com/DouDOU-start/airgate-sdk`（后端）

## 开发环境

- Go 1.23+
- Node 20+
- 本地依赖：`airgate-sdk`（`../../airgate-sdk`）、`airgate-core`（运行环境）

### 启动顺序

1. `cd airgate-core/backend && go run ./cmd/server/` — Core 后端
2. `cd airgate-core/web && npm run dev` — Core 前端（端口 3000）
3. `cd 插件/web && npm run dev` — 插件前端 watch 构建（`vite build --watch`，无独立端口）
4. Core 会自动通过 gRPC 启动插件后端进程

> `npm run dev` 是 `vite build --watch`，会一直挂起，这是正常的。

---

## Gateway 插件开发规范

### 文件职责

| 文件 | 职责 |
|---|---|
| `backend/main.go` | 入口，调用 `sdkgrpc.Serve()` |
| `backend/internal/gateway/metadata.go` | 插件元数据：ID、名称、版本、模型列表、账号类型 |
| `backend/internal/gateway/gateway.go` | `GatewayPlugin` 接口实现 |
| `backend/internal/gateway/forward.go` | HTTP 请求转发逻辑 |
| `backend/internal/gateway/assets.go` | 前端资源嵌入（WebAssetsProvider） |

### 必须实现的接口

```go
type GatewayPlugin interface {
    Info() PluginInfo
    Init(ctx PluginContext) error
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
    Platform() string
    Models() []ModelInfo
    Routes() []RouteDefinition
    Forward(ctx context.Context, req *ForwardRequest) (*ForwardResult, error)
    ValidateAccount(ctx context.Context, credentials map[string]string) error
    QueryQuota(ctx context.Context, credentials map[string]string) (*QuotaInfo, error)
    HandleWebSocket(ctx context.Context, conn WebSocketConn) (*ForwardResult, error)
}
```

### Forward() 要点

- `req.Account.Credentials` — 凭证 map
- `req.Body` — 原始请求体
- `req.Headers` — 原始请求头
- `req.Model` — 请求的模型
- `req.Stream` — 是否流式
- `req.Writer` — 流式写入（SSE）

返回 `ForwardResult`：StatusCode, InputTokens, OutputTokens, Model, Duration, AccountStatus

### 模型定义

```go
sdk.ModelInfo{
    ID: "model-id", Name: "Model Name", MaxTokens: 4096,
    InputPrice: 0.001, OutputPrice: 0.002, CachePrice: 0.0005,
}
```

### 账号类型定义

```go
sdk.AccountType{
    Key: "apikey", Label: "API Key", Description: "说明文字",
    Fields: []sdk.CredentialField{
        {Key: "api_key", Label: "API Key", Type: "password", Required: true},
    },
}
```

---

## Extension 插件开发规范

### 文件职责

| 文件 | 职责 |
|---|---|
| `backend/main.go` | 入口，调用 `sdkgrpc.Serve()` |
| `backend/internal/extension/metadata.go` | 插件元数据：ID、名称、FrontendPages |
| `backend/internal/extension/extension.go` | `ExtensionPlugin` 接口实现 |
| `backend/internal/extension/routes.go` | 自定义 API 路由注册 |
| `backend/internal/extension/assets.go` | 前端资源嵌入（WebAssetsProvider） |

### 必须实现的接口

```go
type ExtensionPlugin interface {
    Info() PluginInfo
    Init(ctx PluginContext) error
    Start(ctx context.Context) error
    Stop(ctx context.Context) error
    RegisterRoutes(r RouteRegistrar)  // 注册自定义 API
    Migrate() error                    // 数据库迁移（可返回 nil）
    BackgroundTasks() []BackgroundTask // 后台任务（可返回 nil）
}
```

### 路由注册

```go
func RegisterRoutes(r sdk.RouteRegistrar, logger *slog.Logger) {
    r.Handle("GET", "/hello", func(w http.ResponseWriter, req *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
    })

    api := r.Group("/api")
    api.Handle("POST", "/action", actionHandler)
}
```

Core 会将 `/api/v1/ext/{pluginID}/*` 的请求代理到插件的 `HandleRequest` gRPC 方法。

### FrontendPages 声明

```go
FrontendPages: []sdk.FrontendPage{
    {Path: "/", Title: "页面标题", Icon: "puzzle", Description: "描述"},
}
```

Core 前端侧边栏会自动显示这些页面入口。

### 后台任务

```go
func (e *MyExtension) BackgroundTasks() []sdk.BackgroundTask {
    return []sdk.BackgroundTask{
        {Name: "cleanup", Interval: 5 * time.Minute, Handler: e.cleanup},
    }
}
```

---

## 前端开发规范（两种类型通用）

### 核心规则

1. **所有 Tailwind 类名自动带 `agw-` 前缀**（Tailwind config 已配置）
2. **颜色只用 token**，禁止硬编码 `#xxx` / `rgb()` / `text-gray-400`
3. **根节点必须有 `data-ag-{plugin}-root` 属性**（CSS 作用域锚点）
4. **根节点必须用 `useScopedPluginTheme`**（跟随 Core 明暗切换）
5. **React 由 Core 提供**，`vite.config.ts` 中已 external，不打包
6. **优先使用 SDK 组件**

### SDK 组件

```tsx
import {
  cn, Field, TextInput, SecretInput, TextArea,
  Section, Card, SelectableCard, Button, FormActions,
  Badge, StatusText,
} from '@airgate/theme/plugin';
```

### 导出格式

**Gateway 插件**（导出 accountForm）：
```ts
export default {
  accountForm: AccountFormComponent,
};
```

**Extension 插件**（导出 routes）：
```ts
export default {
  routes: [
    { path: '/', component: HomePage },
    { path: '/settings', component: SettingsPage },
  ],
};
```

### 设计 Token

- 主色：`agw-bg-primary`, `agw-text-primary`, `agw-bg-primary-hover`
- 状态：`agw-text-success`, `agw-text-warning`, `agw-text-danger`
- 背景层级：`agw-bg-bg-deep` → `agw-bg-bg` → `agw-bg-bg-elevated` → `agw-bg-surface`
- 文字：`agw-text-text`（主）, `agw-text-text-secondary`, `agw-text-text-tertiary`
- 边框：`agw-border-border`, `agw-border-glass-border`, `agw-border-border-focus`
- 圆角：`agw-rounded-sm` (6px), `agw-rounded-md` (10px), `agw-rounded-lg` (14px)

---

## 构建与部署

### 构建命令（在各模板目录下执行）

```bash
make build          # 完整构建：SDK → 前端 → 复制到后端 → Go 编译
make build-web      # 仅前端
make build-backend  # 仅后端（会自动复制前端产物）
make manifest       # 生成 plugin.yaml
```

### 开发模式

在 Core 的 `config.yaml` 中配置：

```yaml
plugins:
  dev:
    - name: my-gateway
      path: /path/to/gateway-template/backend
    - name: my-extension
      path: /path/to/extension-template/backend
```

## 参考项目

- `airgate-openai/` — 完整的 OpenAI 网关插件（含 OAuth、WebSocket、多模型）
- `airgate-sdk/` — SDK 源码
- `airgate-sdk/PLUGIN_STYLE_GUIDE.md` — 前端样式规范完整版
