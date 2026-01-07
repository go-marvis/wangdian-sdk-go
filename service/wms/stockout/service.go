package stockout

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	Other *other
	Sales *sales
}

func NewService(config *core.Config) *Service {
	return &Service{
		Other: &other{config},
		Sales: &sales{config},
	}
}
