package server

import (
	"fmt"
	"github.com/wukongcloud/xxdns/server/models"
	"github.com/wukongcloud/xxdns/server/routers"
	_ "github.com/wukongcloud/xxdns/server/routers"
)

func RunServer(port string) {
	models.InitDb()

	fmt.Printf("\nweb server listen on %s\n", port)
	router := routers.SetupRouter()
	if port != "" {
		router.Run(":" + port)
	}
	router.Run()
}
