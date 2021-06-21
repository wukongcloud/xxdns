package acl

import (
	"encoding/json"
	"fmt"
	"github.com/wukongcloud/xxdns/server/models"
	utils2 "github.com/wukongcloud/xxdns/utils"
	"os"
	"text/template"
)

const (
	APIServer = "http://127.0.0.1:8080"
)

type aclStruct struct {
	Name string
	Acls []models.IPDB
}

func RenderAcl(country string, province string, isp string)(err error) {
	var (
		aclName     string
		queryString string
	)
	if province != "" {
		aclName = province
		queryString = fmt.Sprintf("province=%s", province)
	}

	if isp != "" {
		aclName = isp
		queryString = fmt.Sprintf("isp=%s", isp)
	}
	if province != "" && isp != "" {
		aclName = province + "_" + isp
		queryString = fmt.Sprintf("province=%s&isp=%s", province, isp)
	}

	if country != "" && province != "" && isp != "" {
		aclName = country + "_" + province + "_" + isp
		queryString = fmt.Sprintf("country=%s&province=%s&isp=%s", country, province, isp)
	}



	URL := fmt.Sprintf("%s/api/v1/ipdb?%s&pagesize=1000000", APIServer, queryString)

	var acls = []models.IPDB{}

	aclData, err := utils2.HttpRequest(URL, "GET")
	if err = json.Unmarshal(aclData, &acls); err != nil {
		return err
	}

	aclInfo := aclStruct{
		aclName,
		acls,
	}

	//fmt.Println(aclInfo.Acls)

	aclFileName := fmt.Sprintf("./tmp/%s.cidr", aclName)
	f, err := os.OpenFile(aclFileName, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles("angent/templates/acl.tpl")
	if err != nil {
		return err
	}
	if err = t.Execute(f, aclInfo); err != nil {
		return err
	}
	return err
}
