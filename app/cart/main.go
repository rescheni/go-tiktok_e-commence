package main

import (
	"context"
	"e-commence/app/cart/biz/dal"
	"e-commence/common/mtl"
	"e-commence/common/serversuite"

	"net"
	"os"
	"time"

	"e-commence/app/cart/conf"
	"e-commence/app/cart/rpc"
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"

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

	err := godotenv.Load()

	if err != nil {
		panic("err env load err")
	}

	//
	mtl.IniMetric(serverName, conf.GetConf().Kitex.MetricsPort, registryAddr)
	d := mtl.InitTracing(serverName)
	defer d.Shutdown(context.Background())
	// ⬆️ 下面的两个有可能依赖于mtl

	dal.Init()
	rpc.Init()

	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err = svr.Run()
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
