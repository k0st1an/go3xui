# go3xui
## Supported API

- https://github.com/MHSanaei/3x-ui/wiki/Configuration#api
- https://www.postman.com/hsanaei/3x-ui/collection/q1l5l0u/3x-ui

- [x] `POST /login`
- [x] `POST /server/status`
- [x] `GET /panel/api/inbounds/list`
- [x] `POST /server/getNewX25519Cert`
- [x] `POST /panel/api/inbounds/add`
- [x] `POST /panel/api/inbounds/addClient`


## Usage

```go
package main

import (
	"log"

	"github.com/k0st1an/go3xui"
	"github.com/k0st1an/go3xui/server"
)

func main() {
	apiClient, err := go3xui.NewClient("https://...")
	if err != nil {
		log.Fatal(err)
	}

	if err := apiClient.Login("username", "password"); err != nil {
		log.Fatal(err)
	}

	status, err := server.GetServerStatus(apiClient)
	if err != nil {
		log.Fatal(err)
	}

  log.Printf("%+v", status)
}
```

Output:

```
2025/06/15 00:03:45 &{Success:true Msg: Obj:{CPU:1.0050251498513232 CPUCores:2 LogicalPro:2 CPUSpeedMhz:2194.842 Mem:{Current:509382656 Total:2062512128} Swap:{Current:0 Total:0} Disk:{Current:3725287424 Total:50318852096} Xray:{State:running ErrorMsg: Version:25.3.6} Uptime:5970403 Loads:[0.05 0.04 0.05] TCPCount:49 UDPCount:11 NetIO:{Up:2924 Down:3392} NetTraffic:{Sent:8201743938 Recv:7242729289} PublicIP:{IPv4:111.222.333.444 IPv6:N/A} AppStats:{Threads:10 Mem:45176072 Uptime:2240383}}}
```


## Version

[version](version.go)

## License

MIT
