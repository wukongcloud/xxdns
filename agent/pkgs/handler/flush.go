package handler

import (
	"context"
	"fmt"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"github.com/wukongcloud/xxdns/agent/pkgs/render"
)



var (
	renderHandler =render.NewRender()
)

type FlushFile struct{}

func (m *FlushFile) FlushDomainFile(ctx context.Context, req *services.DomainData) (resp *services.IsFlushed, err error) {
	fmt.Println(req)
	err=renderHandler.Domain(req,"./")
	return
}

func (m *FlushFile) FlushAclFile(ctx context.Context, req *services.ACLS) (resp *services.IsFlushed, err error) {
	err=renderHandler.Acl(req.Country,req.Province,req.ISP,"./",req.IPS)
	return
}
