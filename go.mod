module github.com/networkservicemesh/sdk-vppagent

go 1.13

require (
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.3
	github.com/kr/pretty v0.2.0 // indirect
	github.com/ligato/vpp-agent v2.5.1+incompatible
	github.com/networkservicemesh/api v0.0.0-20200223155536-6728cf448703
	github.com/networkservicemesh/sdk v0.0.0-20200226130054-b405aac645ab
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	go.ligato.io/vpp-agent/v3 v3.0.1
	google.golang.org/grpc v1.27.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
