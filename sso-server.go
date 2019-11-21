package main

import (
	"fmt"

	"github.com/gethinyan/sso-server/pkg/setting"
	"github.com/gethinyan/sso-server/routes"
)

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()
	router.Run(listenAddr)
}
