package domain

import (
	"encoding/json"
	"fmt"
	"github.com/wukongcloud/xxdns/server/models"
	"github.com/wukongcloud/xxdns/utils"
	"os"
	"text/template"
)

type domainStruct struct {
	Domain  models.Domain
	Records []models.Record
}

func RenderDomain(Server string) (err error) {
	s := fmt.Sprintf("http://%s", Server)
	var domain = models.Domain{}
	var records = []models.Record{}

	recordsData, err := utils.HttpRequest(s+"/api/v1/records", "GET")
	if err = json.Unmarshal(recordsData, &records); err != nil {
		return err
	}

	domainData, err := utils.HttpRequest(s+"/api/v1/domains/1", "GET")
	if err = json.Unmarshal(domainData, &domain); err != nil {
		return err
	}

	domainInfo := domainStruct{
		domain,
		records,
	}

	fmt.Println(domainInfo.Domain)

	f, err := os.OpenFile("agent/tmp/db.xxdns.org", os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles("agent/templates/domain.tpl")
	if err != nil {
		return err
	}
	if err = t.Execute(f, domainInfo); err != nil {
		return err
	}
	return nil
}
