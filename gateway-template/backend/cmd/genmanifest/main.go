package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DouDOU-start/airgate-plugin-template/gateway-template/internal/gateway"
)

func main() {
	info := gateway.BuildPluginInfo()
	models := gateway.BuildModels()
	routes := gateway.BuildRoutes()

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("id: %s\n", info.ID))
	sb.WriteString(fmt.Sprintf("name: %s\n", info.Name))
	sb.WriteString(fmt.Sprintf("version: %s\n", info.Version))
	sb.WriteString(fmt.Sprintf("author: %s\n", info.Author))
	sb.WriteString(fmt.Sprintf("description: %s\n", info.Description))
	sb.WriteString(fmt.Sprintf("type: %s\n", info.Type))
	sb.WriteString(fmt.Sprintf("platform: %s\n", gateway.PluginPlatform))

	if len(models) > 0 {
		sb.WriteString("models:\n")
		for _, m := range models {
			sb.WriteString(fmt.Sprintf("  - id: %s\n", m.ID))
			sb.WriteString(fmt.Sprintf("    name: %s\n", m.Name))
		}
	}

	if len(routes) > 0 {
		sb.WriteString("routes:\n")
		for _, r := range routes {
			sb.WriteString(fmt.Sprintf("  - method: %s\n", r.Method))
			sb.WriteString(fmt.Sprintf("    path: %s\n", r.Path))
		}
	}

	if err := os.WriteFile("plugin.yaml", []byte(sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "写入 plugin.yaml 失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("plugin.yaml 已生成")
}
