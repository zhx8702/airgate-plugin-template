package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DouDOU-start/airgate-plugin-template/extension-template/internal/extension"
)

func main() {
	info := extension.BuildPluginInfo()

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("id: %s\n", info.ID))
	sb.WriteString(fmt.Sprintf("name: %s\n", info.Name))
	sb.WriteString(fmt.Sprintf("version: %s\n", info.Version))
	sb.WriteString(fmt.Sprintf("author: %s\n", info.Author))
	sb.WriteString(fmt.Sprintf("description: %s\n", info.Description))
	sb.WriteString(fmt.Sprintf("type: %s\n", info.Type))

	if len(info.FrontendPages) > 0 {
		sb.WriteString("frontend_pages:\n")
		for _, p := range info.FrontendPages {
			sb.WriteString(fmt.Sprintf("  - path: %s\n", p.Path))
			sb.WriteString(fmt.Sprintf("    title: %s\n", p.Title))
			sb.WriteString(fmt.Sprintf("    icon: %s\n", p.Icon))
		}
	}

	if err := os.WriteFile("plugin.yaml", []byte(sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "写入 plugin.yaml 失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("plugin.yaml 已生成")
}
