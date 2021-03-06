package entities

// JResult 响应内容
type JResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// IPAddrs IP 地址信息
type IPAddrs struct {
	IPv4 string
	IPv6 string
}
