# AirGate 插件模板

两种插件模板：`gateway-template/`（API 网关转发）和 `extension-template/`（独立页面/自定义 API）。

## 构建与开发

```bash
# 在各模板目录下执行
make build          # 完整构建：SDK前端 → 插件前端 → Go后端
make build-web      # 仅前端
make build-backend  # 仅后端
make manifest       # 生成 plugin.yaml
make clean          # 清理产物

# 开发模式（前端 watch）
cd web && npm run dev   # vite build --watch，会一直挂起，正常行为
```

Go 后端编译产物：`bin/plugin-template`（gateway）/ `bin/ext-template`（extension）。

## 技术栈

- Go 1.25+, Node 20+, React 19（由 Core 提供，插件不打包）
- 后端 SDK：`github.com/DouDOU-start/airgate-sdk`（本地 replace `../../../airgate-sdk`）
- 前端 SDK：`@airgate/theme`（本地 `file:../../../airgate-sdk/frontend`）
- 前端构建：Vite 库模式，ESM only，入口 `src/index.ts`

## 代码规范

### Go 后端

- Gateway 核心文件在 `backend/internal/gateway/`，Extension 在 `backend/internal/extension/`
- `metadata.go` 是插件元信息单源（ID、名称、版本、模型、账号类型）
- `main.go` 只做 `sdkgrpc.Serve()`，不放业务逻辑

### 前端

- **IMPORTANT**: 所有 Tailwind 类必须带 `agw-` 前缀（如 `agw-flex agw-gap-3`）
- 根节点必须有 `data-ag-{plugin}-root` 属性 + `useScopedPluginTheme` hook
- 颜色只用 token（`agw-bg-primary`, `agw-text-text`），禁止硬编码色值
- 优先使用 `@airgate/theme/plugin` 导出的 SDK 组件（Field, TextInput, Section, Card 等）
- Gateway 导出 `{ accountForm }`，Extension 导出 `{ routes }`

### 前端样式详细规范

参见 [PLUGIN_STYLE_GUIDE.md](https://github.com/DouDOU-start/airgate-sdk/blob/master/PLUGIN_STYLE_GUIDE.md)

## 开发环境配置

在 `airgate-core/backend/config.yaml` 添加插件：

```yaml
plugins:
  dev:
    - name: my-plugin
      path: /path/to/模板目录/backend
```

启动顺序：Core 后端 → Core 前端 → 插件前端 watch。Core 自动通过 gRPC 启动插件后端进程。

## 参考

- [airgate-openai](https://github.com/DouDOU-start/airgate-openai) — 完整的生产 Gateway 插件
- [airgate-sdk](https://github.com/DouDOU-start/airgate-sdk) — SDK 源码
