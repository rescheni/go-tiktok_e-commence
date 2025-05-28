package email

import (
	"context"
	"e-commence/app/email/infra/mq"
	"e-commence/app/email/infra/notify"
	"e-commence/rpc_gen/kitex_gen/email"

	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

func ConsumerInit() {

	tracer := otel.Tracer("shop-nats-consumer")

	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		// 解析数据上传 OpenTelemetry
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))

		_, span := tracer.Start(ctx, "shop-email-consumer")

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})

	if err != nil {
		panic("Consumer Init Err")
	}

	// 服务关闭清理操作
	server.RegisterShutdownHook(func() {
		// 取消订阅
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
