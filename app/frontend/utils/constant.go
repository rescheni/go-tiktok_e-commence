package utils

import (
	"os"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

var ServerName_Frontend = "frontend"
var RegistryAddr_Frontend = os.Getenv(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
