module github.com/cvetkovski98/zvax-slots

replace github.com/cvetkovski98/zvax-common => ../common

go 1.19

require (
	github.com/cvetkovski98/zvax-common v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v9 v9.0.0-beta.2
	github.com/pkg/errors v0.8.1
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
