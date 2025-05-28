package clientsuite

import (
	"os"

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"k8s.io/klog"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {

	opts := []client.Option{

		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
	}

	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
