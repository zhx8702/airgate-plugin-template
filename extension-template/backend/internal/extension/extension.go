package extension

import (
	"context"
	"log/slog"

	sdk "github.com/DouDOU-start/airgate-sdk"
)

// TemplateExtension 实现 sdk.ExtensionPlugin 接口
type TemplateExtension struct {
	logger *slog.Logger
}

func (e *TemplateExtension) Info() sdk.PluginInfo {
	return BuildPluginInfo()
}

func (e *TemplateExtension) Init(ctx sdk.PluginContext) error {
	e.logger = ctx.Logger()
	e.logger.Info("Template 扩展插件初始化")
	return nil
}

func (e *TemplateExtension) Start(ctx context.Context) error {
	e.logger.Info("Template 扩展插件启动")
	return nil
}

func (e *TemplateExtension) Stop(ctx context.Context) error {
	e.logger.Info("Template 扩展插件停止")
	return nil
}

func (e *TemplateExtension) RegisterRoutes(r sdk.RouteRegistrar) {
	RegisterRoutes(r, e.logger)
}

func (e *TemplateExtension) Migrate() error {
	// 无需数据库迁移
	return nil
}

func (e *TemplateExtension) BackgroundTasks() []sdk.BackgroundTask {
	// 无后台任务
	return nil
}
