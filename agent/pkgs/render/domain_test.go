package render

import (
	"embed"
	"github.com/wukongcloud/xxdns/server/models"
	"testing"
)


var (
	domain = models.Domain{
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
	re1 = models.Record{
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

	domainInfo = DomainStruct{}
)







//go:embed static
var static embed.FS
func TestDomain(t *testing.T) {
	records := make([]models.Record, 0)
	records = append(records, re1)
	domainInfo.Records=records
	domainInfo.Domain=domain

	var R=Render{FS: static}
	R.Domain(domainInfo,"./etc/bind/views")

}
