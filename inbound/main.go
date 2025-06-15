package inbound

import "encoding/json"

type InboundList struct {
	Success bool      `json:"success"`
	Msg     string    `json:"msg"`
	Obj     []Inbound `json:"obj"`
}

type InboundAdd struct {
	Success bool    `json:"success"`
	Msg     string  `json:"msg"`
	Obj     Inbound `json:"obj"`
}

type Inbound struct {
	ID             int64                 `json:"id"`
	Up             int64                 `json:"up"`
	Down           int64                 `json:"down"`
	Total          int64                 `json:"total"`
	Remark         string                `json:"remark"`
	Enable         bool                  `json:"enable"`
	ExpiryTime     int64                 `json:"expiryTime"`
	Listen         string                `json:"listen"`
	Port           int64                 `json:"port"`
	Protocol       string                `json:"protocol"`
	Tag            string                `json:"tag"`
	ClientStats    []ClientStat          `json:"clientStats"`
	Settings       InboundSettings       `json:"settings"`
	StreamSettings InboundStreamSettings `json:"streamSettings"`
	Sniffing       InboundSniffing       `json:"sniffing"`
	Allocate       InboundAllocate       `json:"allocate"`
}

type ClientStat struct {
	ID         int64  `json:"id"`
	InboundID  int64  `json:"inboundId"`
	Enable     bool   `json:"enable"`
	Email      string `json:"email"`
	Up         int64  `json:"up"`
	Down       int64  `json:"down"`
	ExpiryTime int64  `json:"expiryTime"`
	Total      int64  `json:"total"`
	Reset      int64  `json:"reset"`
}

type InboundSettings struct {
	Clients    []InboundSettingsClient `json:"clients"`
	Decryption string                  `json:"decryption"`
	Fallbacks  []any                   `json:"fallbacks"` // TODO: implement
}

type InboundSettingsClient struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Email      string `json:"email"`
	LimitIP    int64  `json:"limitIp"`
	TotalGB    int64  `json:"totalGB"`
	ExpiryTime int64  `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgID       string `json:"tgId"`
	SubID      string `json:"subId"`
	Comment    string `json:"comment"`
	Reset      int64  `json:"reset"`
}

type InboundStreamSettings struct {
	Network         string                               `json:"network"`
	Security        string                               `json:"security"`
	ExternalProxy   []any                                `json:"externalProxy"` // TODO: implement
	RealitySettings InboundStreamSettingsRealitySettings `json:"realitySettings"`
	TcpSettings     InboundStreamSettingsTcpSettings     `json:"tcpSettings"`
}

type InboundStreamSettingsRealitySettings struct {
	Show        bool                            `json:"show"`
	Xver        int64                           `json:"xver"`
	Dest        string                          `json:"dest"`
	ServerNames []string                        `json:"serverNames"`
	PrivateKey  string                          `json:"privateKey"`
	MinClient   string                          `json:"minClient"`
	MaxClient   string                          `json:"maxClient"`
	MaxTimediff int64                           `json:"maxTimediff"`
	ShortIds    []string                        `json:"shortIds"`
	Settings    InboundStreamSettingsRealitySub `json:"settings"`
}

type InboundStreamSettingsRealitySub struct {
	PublicKey   string `json:"publicKey"`
	Fingerprint string `json:"fingerprint"`
	ServerName  string `json:"serverName"`
	SpiderX     string `json:"spiderX"`
}

type InboundStreamSettingsTcpSettings struct {
	AcceptProxyProtocol bool                                   `json:"acceptProxyProtocol"`
	Header              InboundStreamSettingsTcpSettingsHeader `json:"header"`
}

type InboundStreamSettingsTcpSettingsHeader struct {
	Type string `json:"type"`
}

type InboundSniffing struct {
	Enabled      bool     `json:"enabled"`
	DestOverride []string `json:"destOverride"`
	MetadataOnly bool     `json:"metadataOnly"`
	RouteOnly    bool     `json:"routeOnly"`
}

type InboundAllocate struct {
	Strategy    string `json:"strategy"`
	Refresh     int64  `json:"refresh"`
	Concurrency int64  `json:"concurrency"`
}

func decodeStringToStruct(s string, d any) error {
	err := json.Unmarshal([]byte(s), d)
	return err
}

// ------------------------------------------------------------

type inboundList struct {
	Success bool      `json:"success"`
	Msg     string    `json:"msg"`
	Obj     []inbound `json:"obj"`
}

type inboundAdd struct {
	Success bool    `json:"success"`
	Msg     string  `json:"msg"`
	Obj     inbound `json:"obj"`
}

type inbound struct {
	ID             int64        `json:"id"`
	Up             int64        `json:"up"`
	Down           int64        `json:"down"`
	Total          int64        `json:"total"`
	Remark         string       `json:"remark"`
	Enable         bool         `json:"enable"`
	ExpiryTime     int64        `json:"expiryTime"`
	Listen         string       `json:"listen"`
	Port           int64        `json:"port"`
	Protocol       string       `json:"protocol"`
	Tag            string       `json:"tag"`
	ClientStats    []ClientStat `json:"clientStats"`
	Settings       string       `json:"settings"`
	StreamSettings string       `json:"streamSettings"`
	Sniffing       string       `json:"sniffing"`
	Allocate       string       `json:"allocate"`
}
