cd services/protos
protoc --go_out=plugins=grpc:../ xxddns.proto
protoc-go-inject-tag  -input=../xxddns.pb.go
cd ..&cd ..

