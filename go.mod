module github.com/go-to/egp-server

go 1.23.3

//replace egp-protobuf/pb => ../egp-protobuf/pb

require (
	github.com/go-to/egp-protobuf/pb v0.0.0-20250101165031-941231392b1b
	google.golang.org/grpc v1.69.2
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)
