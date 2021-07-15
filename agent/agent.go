package agent

import (
	"embed"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler"
	"github.com/wukongcloud/xxdns/agent/pkgs/handler/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:embed pkgs/render/static
var tmpl embed.FS

func RunAgent(ServerIP string) (err error) {
	//run gRPC
	lis, err := net.Listen("tcp", ServerIP)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	services.RegisterFlushFileServer(s, &handler.FlushFile{})
	err = s.Serve(lis)
	if err != nil {
		log.Println("grpc server run err:", err)
	}
	log.Printf("grpc server in: %s", ServerIP)
	return nil
}
