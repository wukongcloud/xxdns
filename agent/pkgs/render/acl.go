package render

import (
	"fmt"
	"github.com/wukongcloud/xxdns/server/models"
	"os"
	"text/template"
)

func (m *Render) Acl(country, province, isp string, acls []models.IPDB) (err error) {
	var (
		aclName string
	)

	switch {
	case province != "":
		aclName = province
	case isp != "":
		aclName = isp
	case province != "" && isp != "":
		aclName = province + "_" + isp
	case country != "" && province != "" && isp != "":
		aclName = country + "_" + province + "_" + isp
	}

	aclInfo := AclStruct{
		aclName,
		acls,
	}

	aclFileName := fmt.Sprintf("%s.cidr", aclName)
	f, err := os.OpenFile(aclFileName, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles("static/templates/acl.tpl")
	if err != nil {
		return err
	}
	if err = t.Execute(f, aclInfo); err != nil {
		return err
	}
	return err
}
