# AirGate 插件模板

AirGate 网关插件的标准开发模板，包含两种插件类型：

| 模板 | 类型 | 适用场景 |
|---|---|---|
| `gateway-template/` | Gateway | API 请求转发（OpenAI、Claude 等） |
| `extension-template/` | Extension | 独立页面、自定义 API、后台任务 |

## 快速开始

### 1. 选择模板并复制

```bash
# Gateway 插件（API 转发）
cp -r gateway-template ../airgate-myplugin

# Extension 插件（独立页面/自定义 API）
cp -r extension-template ../airgate-myplugin
```

### 2. 全局替换标识

**Gateway 模板**：

| 占位符 | 替换为 | 示例 |
|---|---|---|
| `template` (小写) | 你的插件 ID | `azure`, `claude` |
| `Template` (首字母大写) | 你的插件名 | `Azure`, `Claude` |
| `gateway-template` | 你的项目名 | `airgate-azure` |

**Extension 模板**：

| 占位符 | 替换为 | 示例 |
|---|---|---|
| `ext-template` | 你的插件 ID | `ext-payment`, `ext-monitor` |
| `Template` (首字母大写) | 你的插件名 | `Payment`, `Monitor` |
| `extension-template` | 你的项目名 | `airgate-payment` |

### 3. 安装依赖

```bash
cd web && npm install
cd ../backend && go mod tidy
```

### 4. 配置 Core

在 `airgate-core/backend/config.yaml` 中添加开发插件：

```yaml
plugins:
  dev:
    - name: my-plugin
      path: /path/to/your-plugin/backend
```

### 5. 启动开发

```bash
# 终端 1: Core 后端
cd airgate-core/backend && go run ./cmd/server/

# 终端 2: Core 前端
cd airgate-core/web && npm run dev

# 终端 3: 插件前端 watch 构建
cd your-plugin/web && npm run dev
```

> `npm run dev` 是 `vite build --watch`，会持续运行并监听文件变化。

## 项目结构

```
├── gateway-template/           # Gateway 网关插件模板
│   ├── Makefile
│   ├── backend/
│   │   ├── main.go             # sdkgrpc.Serve(&gateway.TemplateGateway{})
│   │   └── internal/gateway/
│   │       ├── metadata.go     # 模型、账号类型、路由定义
│   │       ├── gateway.go      # GatewayPlugin 接口实现
│   │       ├── forward.go      # 请求转发逻辑
│   │       └── assets.go       # 前端资源嵌入
│   └── web/
│       └── src/
│           ├── index.ts        # 导出 { accountForm }
│           └── components/     # 账号表单组件
│
├── extension-template/         # Extension 扩展插件模板
│   ├── Makefile
│   ├── backend/
│   │   ├── main.go             # sdkgrpc.Serve(&extension.TemplateExtension{})
│   │   └── internal/extension/
│   │       ├── metadata.go     # FrontendPages 声明
│   │       ├── extension.go    # ExtensionPlugin 接口实现
│   │       ├── routes.go       # 自定义 API 路由
│   │       └── assets.go       # 前端资源嵌入
│   └── web/
│       └── src/
│           ├── index.ts        # 导出 { routes }
│           └── pages/          # 页面组件
│
├── CLAUDE.md                   # AI 辅助开发上下文
└── README.md
```

## 构建

在各模板目录下执行：

```bash
make build        # 完整构建（前端 + 后端）
make build-web    # 仅构建前端
make build-backend # 仅构建后端
make manifest     # 生成 plugin.yaml
make clean        # 清理构建产物
```

## 参考

- [CLAUDE.md](./CLAUDE.md) — AI 辅助开发的完整上下文
- [插件前端样式规范](https://github.com/DouDOU-start/airgate-sdk/blob/master/PLUGIN_STYLE_GUIDE.md)
- [airgate-openai](https://github.com/DouDOU-start/airgate-openai) — 完整的生产插件参考
