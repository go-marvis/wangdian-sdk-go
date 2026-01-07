package outer

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	OuterIn  *outerIn
	OuterOut *outerOut
}

func NewService(config *core.Config) *Service {
	return &Service{
		OuterIn:  &outerIn{config: config},
		OuterOut: &outerOut{config: config},
	}
}
