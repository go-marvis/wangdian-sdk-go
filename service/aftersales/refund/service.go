package refund

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	Refund *refund
}

func NewService(config *core.Config) *Service {
	return &Service{
		Refund: &refund{config},
	}
}
