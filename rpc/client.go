package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"sync"
)

var (
	once       sync.Once
	AuthClient authservice.Client
	UserClient userservice.Client
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
		initUserClient()
	})
}
