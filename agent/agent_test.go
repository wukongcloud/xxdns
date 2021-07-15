package agent

import (
	"context"
	"fmt"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"google.golang.org/grpc"
	"testing"
)

var (
	mockDomain	= services.Domain{
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
	mockRecord = services.Record{
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
)


func TestRunAgent(t *testing.T) {
	gRPCClient()
}

func gRPCClient() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn err:",err)
	}
	defer conn.Close()
	client := services.NewFlushFileClient(conn)
	var req services.DomainData
	records := make([]*services.Record, 0)
	records = append(records, &mockRecord)

	req.Records = records
	req.Domain = &mockDomain
	//todo:  FlushAclFile
	//client.FlushAclFile(context.TODO(),)
	_, err = client.FlushDomainFile(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
	}
}
