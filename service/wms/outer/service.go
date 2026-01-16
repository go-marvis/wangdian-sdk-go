package outer

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer/outer_in"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer/outer_out"
)

type Service struct {
	OuterIn  *outer_in.Service
	OuterOut *outer_out.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		OuterIn:  outer_in.NewService(config),
		OuterOut: outer_out.NewService(config),
	}
}
