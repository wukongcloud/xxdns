package render

import (
	"fmt"
	"github.com/wukongcloud/xxdns/server/models"
	"os"
	"path/filepath"
	"text/template"
)

func (m *Render) Acl(country, province, isp, dirPath string, acls []models.IPDB) (err error) {
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
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	aclFileName := fmt.Sprintf("%s.acl", aclName)
	f, err := os.OpenFile(filepath.Join(dirPath, aclFileName), os.O_WRONLY|os.O_CREATE, 0644)
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

func (m *Render) AclHeader(dirPath string, aclList []Acl) (err error) {
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(dirPath, "acls.conf"), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles("static/templates/acls.tpl")
	if err != nil {
		return err
	}
	if err = t.Execute(f, aclList); err != nil {
		return err
	}
	return err
}
