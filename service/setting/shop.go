package setting

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type shop struct {
	config *core.Config
}

type QueryShopReq struct {
	ShopNo     string `json:"shop_no,omitempty"`     // 店铺编号
	PlatformId int    `json:"platform_id,omitempty"` // 平台id
}

type QueryShopResp struct {
	Details    []*ShopDetail `json:"details"`
	TotalCount int64         `json:"total_count"`
}

type ShopDetail struct {
	ShopId        int    `json:"shop_id"`         // 店铺id
	ShopName      string `json:"shop_name"`       // 店铺名称
	ShopNo        string `json:"shop_no"`         // 店铺编号
	PlatformId    int    `json:"platform_id"`     // 平台id
	SubPlatformId int    `json:"sub_platform_id"` // 子平台id
	Contact       string `json:"contact"`         // 联系人
	Province      string `json:"province"`        // 省份
	City          string `json:"city"`            // 城市
	District      string `json:"district"`        // 区县
	Address       string `json:"address"`         // 地址
	Telno         string `json:"telno"`           // 电话
	Mobile        string `json:"mobile"`          // 手机
	Zip           string `json:"zip"`             // 邮编
	Email         string `json:"email"`           // Email
	Remark        string `json:"remark"`          // 备注
	Website       string `json:"website"`         // 网址
	GroupId       string `json:"group_id"`        // 分组
	AccountId     string `json:"account_id"`      // 平台的主键
	IsDisabled    bool   `json:"is_disabled"`     // 停用
	AuthState     int    `json:"auth_state"`      // 授权状态
	AuthTime      string `json:"auth_time"`       // 授权时间
	ReExpireTime  string `json:"re_expire_time"`  // refresh_token过期时间
	Modified      string `json:"modified"`        // 修改时间
	ExpireTime    string `json:"expire_time"`     // session过期时间
	Created       string `json:"created"`         // 创建时间
}

func (s *shop) QueryShop(ctx context.Context, req *QueryShopReq, options ...core.ReqOptionFunc) (*QueryShopResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "setting.Shop.queryShop",
		Body:       []any{req},
	}
	return core.Retry[QueryShopResp](ctx, s.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, s.config, options...)
	})
}
