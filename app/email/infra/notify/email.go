package notify

import (
	"e-commence/rpc_gen/kitex_gen/email"

	"github.com/kr/pretty"
)

type NoopEmail struct {
}

func (e *NoopEmail) Send(req *email.EmailReq) (err error) {
	pretty.Printf("%v", req)
	return
}

func NewNoopEmail() NoopEmail {

	return NoopEmail{}
}
