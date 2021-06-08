
package main

import (
   "github.com/wukongcloud/xxdns/routers"
   _ "github.com/wukongcloud/xxdns/routers"
)

func main() {
   r := routers.SetupRouter()
   r.Run()
}

