package render

import (
	"github.com/wukongcloud/xxdns/server/models"
	"testing"
)

func TestAcl(t *testing.T) {
	var R = Render{FS: static}
	R.Domain(domainInfo)

	acl1 := models.IPDB{
		CIDR:     "111",
		Country:  "111",
		Province: "111",
		ISP:      "111",
		Comment:  "111",
	}
	acl2 := models.IPDB{
		CIDR:     "222",
		Country:  "222",
		Province: "222",
		ISP:      "222",
		Comment:  "222",
	}

	acls := make([]models.IPDB, 0)
	acls = append(acls, acl1,acl2)

	R.Acl("cn", "hangzhou", "laintong",acls)
}
