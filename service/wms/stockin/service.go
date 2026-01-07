package stockin

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	Other    *other
	Purchase *purchase
	Refund   *refund
}

func NewService(config *core.Config) *Service {
	return &Service{
		Other:    &other{config},
		Purchase: &purchase{config},
		Refund:   &refund{config},
	}
}
