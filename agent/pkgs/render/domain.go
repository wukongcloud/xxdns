package render

import (
	"fmt"
	"os"
	"text/template"
)




func (m *Render) Domain(domainInfo DomainStruct)(err error)  {
	f, err := os.OpenFile(fmt.Sprintf("db.%s",domainInfo.Domain.Name), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	t, err := template.ParseFS(m.FS, "static/templates/*.tpl")
	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(f, "domain.tpl",domainInfo); err != nil {
		return err
	}
	return nil
}
