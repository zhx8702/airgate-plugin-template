package extension

import (
	"encoding/json"
	"log/slog"
	"net/http"

	sdk "github.com/DouDOU-start/airgate-sdk"
)

// RegisterRoutes 注册扩展插件的 API 路由
// Core 会将 /api/v1/ext/ext-template/* 的请求代理到这里
func RegisterRoutes(r sdk.RouteRegistrar, logger *slog.Logger) {
	r.Handle("GET", "/hello", func(w http.ResponseWriter, req *http.Request) {
		logger.Info("收到 hello 请求")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello from Template Extension Plugin!",
			"plugin":  PluginID,
			"version": PluginVersion,
		})
	})

	r.Handle("GET", "/status", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "running",
			"plugin":  PluginID,
			"version": PluginVersion,
		})
	})
}
