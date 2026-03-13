package main

import (
	"github.com/DouDOU-start/airgate-plugin-template/extension-template/internal/extension"

	sdkgrpc "github.com/DouDOU-start/airgate-sdk/grpc"
)

func main() {
	sdkgrpc.Serve(&extension.TemplateExtension{})
}
