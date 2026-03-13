package gateway

import (
	"context"
	"fmt"
	"log/slog"

	sdk "github.com/DouDOU-start/airgate-sdk"
)

// TemplateGateway 实现 sdk.GatewayPlugin 接口
type TemplateGateway struct {
	logger *slog.Logger
}

func (g *TemplateGateway) Info() sdk.PluginInfo {
	return BuildPluginInfo()
}

func (g *TemplateGateway) Init(ctx sdk.PluginContext) error {
	g.logger = ctx.Logger()
	g.logger.Info("Template 网关插件初始化")
	return nil
}

func (g *TemplateGateway) Start(ctx context.Context) error {
	g.logger.Info("Template 网关插件启动")
	return nil
}

func (g *TemplateGateway) Stop(ctx context.Context) error {
	g.logger.Info("Template 网关插件停止")
	return nil
}

func (g *TemplateGateway) Platform() string {
	return PluginPlatform
}

func (g *TemplateGateway) Models() []sdk.ModelInfo {
	return BuildModels()
}

func (g *TemplateGateway) Routes() []sdk.RouteDefinition {
	return BuildRoutes()
}

func (g *TemplateGateway) Forward(ctx context.Context, req *sdk.ForwardRequest) (*sdk.ForwardResult, error) {
	return ForwardRequest(ctx, req, g.logger)
}

func (g *TemplateGateway) ValidateAccount(ctx context.Context, credentials map[string]string) error {
	apiKey := credentials["api_key"]
	if apiKey == "" {
		return fmt.Errorf("API Key 不能为空")
	}
	return nil
}

func (g *TemplateGateway) QueryQuota(ctx context.Context, credentials map[string]string) (*sdk.QuotaInfo, error) {
	return nil, sdk.ErrNotSupported
}

func (g *TemplateGateway) HandleWebSocket(ctx context.Context, conn sdk.WebSocketConn) (*sdk.ForwardResult, error) {
	return nil, sdk.ErrNotSupported
}
