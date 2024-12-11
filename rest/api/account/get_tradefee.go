package account

import "github.com/wangbaojin/go-okx/rest/api"

const GetTradeFeeLimitNumPerSec = 5
const GetTradeFeeLimitRule = "UserID"

func NewGetTradeFee(param *GetTradeFeeParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/trade-fee",
		Method: api.MethodGet,
		Param:  param,
	}, &GetTradeFeeResponse{}
}

type GetTradeFeeParam struct {
	InstId   string `url:"instId,omitempty"`
	InstType string `url:"instType,omitempty"`
}

// 查看当前用户的交易费率
type GetTradeFeeResponse struct {
	api.Response
	Data []struct {
		Level     string `json:"level"`     // 手续费等级
		Taker     string `json:"taker"`     // 对于币币，为 USDT 交易区的吃单手续费率
		Maker     string `json:"maker"`     // 对于币币，为 USDT 交易区的挂单手续费率
		TakerU    string `json:"takerU"`    // USDT 合约吃单手续费率，仅适用于交割/永续
		MakerU    string `json:"makerU"`    // USDT 合约挂单手续费率，仅适用于交割/永续
		Delivery  string `json:"delivery"`  // 交割手续费率
		Exercise  string `json:"exercise"`  // 行权手续费率
		InstType  string `json:"instType"`  // 产品类型
		TakerUSDC string `json:"takerUsdc"` // 对于币币，为 USDⓈ&Crypto 交易区的吃单手续费率
		MakerUSDC string `json:"makerUsdc"` // 对于币币，为 USDⓈ&Crypto 交易区的挂单手续费率
		Ts        string `json:"ts"`        // 数据返回时间，Unix时间戳的毫秒数格式
		Fiat      []struct {
			Ccy   string `json:"ccy"`   // 法币币种
			Taker string `json:"taker"` // 吃单手续费率
			Maker string `json:"maker"` // 挂单手续费率
		} `json:"fiat"` // 法币费率
	} `json:"data"`
}
