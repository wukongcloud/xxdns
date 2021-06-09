
package main

import (
   "github.com/wukongcloud/xxdns/models"
   "github.com/wukongcloud/xxdns/routers"
   _ "github.com/wukongcloud/xxdns/routers"
)

func main() {
   models.InitDb()
   r := routers.SetupRouter()
   r.Run()
}

