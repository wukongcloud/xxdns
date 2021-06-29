package render

import (
	"github.com/wukongcloud/xxdns/server/models"
	"io/fs"
)


type DomainStruct struct {
	Domain  models.Domain
	Records []models.Record
}
type AclStruct struct {
	Name string
	Acls []models.IPDB
}

type Render struct {
	FS fs.FS
}
