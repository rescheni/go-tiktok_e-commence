package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"e-commence/app/checkout/biz/dal"
	"e-commence/app/checkout/conf"
	"e-commence/app/checkout/infra/rpc"
	"e-commence/app/email/infra/mq"
	"e-commence/common/mtl"
	"e-commence/common/serversuite"
	"e-commence/rpc_gen/kitex_gen/checkout/checkoutservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serverName   = conf.GetConf().Kitex.Service
	registryAddr = os.Getenv(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
)

func main() {

	_ = godotenv.Load()

	fmt.Println("Starting Checkout Service...")
	fmt.Println(os.Getenv("GOMALL_NATS_URL") + ":" + os.Getenv("GOMALL_NATS_SERVER_PORT"))

	mtl.IniMetric(serverName, conf.GetConf().Kitex.MetricsPort, registryAddr)
	d := mtl.InitTracing(serverName)
	defer d.Shutdown(context.Background())

	mq.Init()
	rpc.InitClient()
	dal.Init()

	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(
		serversuite.CommonServerSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		},
	))
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
