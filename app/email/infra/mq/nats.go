package mq

import (
	"os"

	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect("nats://" + os.Getenv("GOMALL_NATS_URL") + ":" + os.Getenv("GOMALL_NATS_SERVER_PORT"))
	if err != nil {
		panic("Nats connect error")
	}
}
