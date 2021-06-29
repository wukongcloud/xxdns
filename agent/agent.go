package agent

import (
	"embed"
)

//go:embed pkgs/render/static
var tmpl embed.FS

func RunAgent(ServerIP string) (err error) {
	//render view
	//err = domain.RenderDomain(ServerIP)
	//if err != nil {
	//	return err
	//}
	//err=acl.RenderAcl("", "hebei", "unicom")
	//if err != nil {
	//	return err
	//}
	return nil
}
