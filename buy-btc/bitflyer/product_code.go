package bitflyer

type ProductCode int

// iotaはBtcjpyを0としてインクリメントした数字をそれぞれに割り振る
const (
	Btcjpy ProductCode = iota
	Ethjpy
	Fxbtcjpy
	Ethbtc
	bchbtc
)

func (code ProductCode) String() string {
	switch code {
	case Btcjpy:
		return "BTC_JPY"
	case Ethjpy:
		return "ETH_JPY"
	case Fxbtcjpy:
		return "EX_BTC_JPY"
	case Ethbtc:
		return "ETH_BTC"
	case bchbtc:
		return "BCH_BTC"
	default:
		return "BTC_JPY"
	}
}
