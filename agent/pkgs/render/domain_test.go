package render

import (
	"embed"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"testing"
)


var (
	domain = &services.Domain{
		ID:       1,
		Name:     "test",
		Type:     "Type1",
		TTL:      2,
		Provider: "ttt",
		Contact:  "dwq",
		Serial:   3,
		Refresh:  4,
		Retry:    5,
		Expire:   6,
		Mininum:  7,
		Comment:  "",
		Disabled: false,
	}
	re1 = &services.Record{
		ID:       8,
		Domain:   "Domain",
		View:     "View",
		Name:     "Name",
		Type:     "Type",
		Content:  "Content",
		TTL:      9,
		Priority: 10,
		Disabled: false,
		Comment:  "Comment",
	}

	domainInfo = &services.DomainData{}
)







//go:embed static
var static embed.FS
func TestDomain(t *testing.T) {
	records := make([]*services.Record, 0)
	records = append(records, re1)
	domainInfo.Records=records
	domainInfo.Domain=domain

	var R=NewRender()
	R.Domain(domainInfo,"./etc/bind/views")

}
