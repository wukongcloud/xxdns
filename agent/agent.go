package agent

import (
	"github.com/wukongcloud/xxdns/agent/pkgs/acl"
	"github.com/wukongcloud/xxdns/agent/pkgs/domain"
)

func RunAgent(ServerIP string) (err error) {
	//render view
	err = domain.RenderDomain(ServerIP)
	if err != nil {
		return err
	}
	err=acl.RenderAcl("", "hebei", "unicom")
	if err != nil {
		return err
	}
	return nil
}
