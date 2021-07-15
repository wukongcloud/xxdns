package render

import (
	"fmt"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"os"
	"path/filepath"
	"text/template"
)

func (m *Render) Domain(domainInfo *services.DomainData, dirPath string) (err error) {
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("db.%s", domainInfo.Domain.Name)
	f, err := os.OpenFile(filepath.Join(dirPath, filename), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	t, err := template.ParseFS(m.FS, "static/templates/*.tpl")
	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(f, "domain.tpl", domainInfo); err != nil {
		return err
	}
	return nil
}
