package gateway

import sdk "github.com/DouDOU-start/airgate-sdk"

const (
	PluginID          = "gateway-template"
	PluginDisplayName = "Template 网关插件"
	PluginVersion     = "0.1.0"
	PluginAuthor      = "AirGate"
	PluginDescription = "网关插件模板，用于快速开发新的 API 网关插件"
	PluginPlatform    = "template"
)

func BuildPluginInfo() sdk.PluginInfo {
	return sdk.PluginInfo{
		ID:          PluginID,
		Name:        PluginDisplayName,
		Version:     PluginVersion,
		Author:      PluginAuthor,
		Description: PluginDescription,
		Type:        sdk.PluginTypeGateway,
		AccountTypes: []sdk.AccountType{
			{
				Key:         "apikey",
				Label:       "API Key",
				Description: "使用 API Key 访问",
				Fields: []sdk.CredentialField{
					{Key: "api_key", Label: "API Key", Type: "password", Required: true, Placeholder: "sk-..."},
					{Key: "base_url", Label: "API 地址", Type: "text", Required: false, Placeholder: "https://api.example.com"},
				},
			},
		},
	}
}

func BuildModels() []sdk.ModelInfo {
	return []sdk.ModelInfo{
		{
			ID:          "template-model-1",
			Name:        "Template Model 1",
			MaxTokens:   4096,
			InputPrice:  0.001,
			OutputPrice: 0.002,
		},
	}
}

func BuildRoutes() []sdk.RouteDefinition {
	return []sdk.RouteDefinition{
		{Method: "POST", Path: "/v1/chat/completions"},
	}
}
