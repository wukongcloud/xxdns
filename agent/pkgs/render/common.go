package render

import (
	"embed"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"io/fs"
)

//type DomainStruct struct {
//	Domain  models.Domain
//	Records []models.Record
//}
type AclStruct struct {
	Name string
	Acls []*services.IP
}

type Render struct {
	FS fs.FS
}

type Acl struct {
	CreateTime string
	UpdateTime string
	ISP        string
}

//go:embed static
var Static embed.FS
func NewRender() *Render {
	return &Render{
		Static,
	}
}
