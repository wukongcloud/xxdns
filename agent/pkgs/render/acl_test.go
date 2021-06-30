package render

import (
	"github.com/wukongcloud/xxdns/server/models"
	"testing"
	"time"
)

var R = Render{FS: static}

func TestAcl(t *testing.T) {
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
	acls = append(acls, acl1, acl2)

	R.Acl("cn", "hangzhou", "laintong", "./etc/bind/named/acl", acls)
}

func TestRender_AclHeader(t *testing.T) {
	var AclList []Acl

	var Acl1 = Acl{
		CreateTime: time.Now().Format("20060102-15:04:05"),
		UpdateTime: time.Now().Format("20060102-15:04:05"),
		ISP:        "dianxin",
	}
	var Acl2 = Acl{
		CreateTime: time.Now().Format("20060102-15:04:05"),
		UpdateTime: time.Now().Format("20060102-15:04:05"),
		ISP:        "yidong",
	}

	AclList = append(AclList, Acl1, Acl2)

	t.Log("R.AclHeader")
	err := R.AclHeader("./etc/bind/named", AclList)
	if err != nil {
		t.Log(err)
	}
}
