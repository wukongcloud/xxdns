package render

import (
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"testing"
	"time"
)


func TestAcl(t *testing.T) {
	acl1 := &services.IP{
		CIDR:     "111",
		Country:  "111",
		Province: "111",
		ISP:      "111",
		Comment:  "111",
	}
	acl2 := &services.IP{
		CIDR:     "222",
		Country:  "222",
		Province: "222",
		ISP:      "222",
		Comment:  "222",
	}

	acls := make([]*services.IP, 0)
	acls = append(acls, acl1, acl2)

	var R=NewRender()
	R.Acl("cn", "hangzhou", "laintong", "./etc/bind/named/acl", acls)
}

func TestRender_AclHeader(t *testing.T) {
	var R=NewRender()

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
