package extension

import sdk "github.com/DouDOU-start/airgate-sdk"

const (
	PluginID          = "ext-template"
	PluginDisplayName = "Template 扩展插件"
	PluginVersion     = "0.1.0"
	PluginAuthor      = "AirGate"
	PluginDescription = "扩展插件模板，用于快速开发独立页面和自定义 API"
)

func BuildPluginInfo() sdk.PluginInfo {
	return sdk.PluginInfo{
		ID:          PluginID,
		Name:        PluginDisplayName,
		Version:     PluginVersion,
		Author:      PluginAuthor,
		Description: PluginDescription,
		Type:        sdk.PluginTypeExtension,
		FrontendPages: []sdk.FrontendPage{
			{
				Path:        "/",
				Title:       "模板页面",
				Icon:        "puzzle",
				Description: "扩展插件示例页面",
			},
		},
	}
}
