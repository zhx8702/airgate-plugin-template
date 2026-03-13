package main

import (
	"github.com/DouDOU-start/airgate-plugin-template/gateway-template/internal/gateway"

	sdkgrpc "github.com/DouDOU-start/airgate-sdk/grpc"
)

func main() {
	sdkgrpc.Serve(&gateway.TemplateGateway{})
}
