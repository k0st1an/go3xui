package server

type ServerStatusResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     struct {
		CPU         float64 `json:"cpu"`
		CPUCores    int     `json:"cpuCores"`
		LogicalPro  int     `json:"logicalPro"`
		CPUSpeedMhz float64 `json:"cpuSpeedMhz"`
		Mem         struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"mem"`
		Swap struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"swap"`
		Disk struct {
			Current uint64 `json:"current"`
			Total   uint64 `json:"total"`
		} `json:"disk"`
		Xray struct {
			State    string `json:"state"`
			ErrorMsg string `json:"errorMsg"`
			Version  string `json:"version"`
		} `json:"xray"`
		Uptime   int64     `json:"uptime"`
		Loads    []float64 `json:"loads"`
		TCPCount int       `json:"tcpCount"`
		UDPCount int       `json:"udpCount"`
		NetIO    struct {
			Up   uint64 `json:"up"`
			Down uint64 `json:"down"`
		} `json:"netIO"`
		NetTraffic struct {
			Sent uint64 `json:"sent"`
			Recv uint64 `json:"recv"`
		} `json:"netTraffic"`
		PublicIP struct {
			IPv4 string `json:"ipv4"`
			IPv6 string `json:"ipv6"`
		} `json:"publicIP"`
		AppStats struct {
			Threads int    `json:"threads"`
			Mem     uint64 `json:"mem"`
			Uptime  int64  `json:"uptime"`
		} `json:"appStats"`
	} `json:"obj"`
}

type X25519Cert struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	} `json:"obj"`
}
