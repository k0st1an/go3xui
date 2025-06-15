```go
package main

import (
	"log"

	"github.com/k0st1an/go3xui"
	"github.com/k0st1an/go3xui/inbound"
	"github.com/k0st1an/go3xui/server"
)

func main() {
	apiClient, err := go3xui.NewClient("http://localhost:2053")
	if err != nil {
		log.Fatal(err)
	}

	if err := apiClient.Login("admin", "admin"); err != nil {
		log.Fatal(err)
	}

	newX25519Cert, err := server.GetNewX25519Cert(apiClient)
	if err != nil {
		log.Fatal(err)
	}

	newInbound := inbound.Inbound{
		Up:         0,
		Down:       0,
		Total:      0,
		Remark:     "test",
		Enable:     true,
		ExpiryTime: 0,
		Listen:     "",
		Port:       4321,
		Protocol:   "vless",

		Settings: inbound.InboundSettings{
			Clients: []inbound.InboundSettingsClient{
				{
					ID:         go3xui.GenerateUUID(),
					Flow:       "xtls-rprx-vision",
					Email:      go3xui.RandomHexString(9),
					LimitIP:    0,
					TotalGB:    0,
					ExpiryTime: 0,
					Enable:     true,
					TgID:       "",
					SubID:      go3xui.RandomHexString(9),
					Comment:    "",
					Reset:      0,
				},
			},
			Decryption: "none",
			Fallbacks:  []any{},
		},

		StreamSettings: inbound.InboundStreamSettings{
			Network:       "tcp",
			Security:      "reality",
			ExternalProxy: []any{},
			RealitySettings: inbound.InboundStreamSettingsRealitySettings{
				Show:        false,
				Xver:        0,
				Dest:        "yahoo.com:443",
				ServerNames: []string{"yahoo.com", "www.yahoo.com"},
				PrivateKey:  newX25519Cert.Obj.PrivateKey,
				MinClient:   "",
				MaxClient:   "",
				MaxTimediff: 0,
				ShortIds:    go3xui.GenerateShortIds(),
				Settings: inbound.InboundStreamSettingsRealitySub{
					PublicKey:   newX25519Cert.Obj.PublicKey,
					Fingerprint: "firefox",
					ServerName:  "",
					SpiderX:     "/",
				},
			},
			TcpSettings: inbound.InboundStreamSettingsTcpSettings{
				AcceptProxyProtocol: false,
				Header:              inbound.InboundStreamSettingsTcpSettingsHeader{Type: "none"},
			},
		},

		Sniffing: inbound.InboundSniffing{
			Enabled:      true,
			DestOverride: []string{"http", "tls", "quic", "fakedns"},
			MetadataOnly: false,
			RouteOnly:    false,
		},

		Allocate: inbound.InboundAllocate{
			Strategy:    "always",
			Refresh:     5,
			Concurrency: 3,
		},
	}

	inboundResponse, err := inbound.Add(apiClient, newInbound)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inbound added response:")
	log.Printf("%+v", inboundResponse)
}
```

Output:

```
2025/06/16 00:29:24 Inbound added response:
2025/06/16 00:29:24 &{Success:true Msg:Create Successfully Obj:{ID:1 Up:0 Down:0 Total:0 Remark:test Enable:true ExpiryTime:0 Listen: Port:4321 Protocol:vless Tag:inbound-4321 ClientStats:[] Settings:{Clients:[{ID:1e2a3c4a-c5b1-4ab2-b0bb-c17a05c0f8fd Flow:xtls-rprx-vision Email:dd8efca47 LimitIP:0 TotalGB:0 ExpiryTime:0 Enable:true TgID: SubID:f015dc070 Comment: Reset:0}] Decryption:none Fallbacks:[]} StreamSettings:{Network:tcp Security:reality ExternalProxy:[] RealitySettings:{Show:false Xver:0 Dest:yahoo.com:443 ServerNames:[yahoo.com www.yahoo.com] PrivateKey:2LlBkQwrtrKEi5z-GepphseP5AUWVafMVazOCAZJk14 MinClient: MaxClient: MaxTimediff:0 ShortIds:[1be5 5e10b31a 8a 022bbe c92d88b012 a8f060c6 b7ea360314d3f4 90da00d1d32380b6] Settings:{PublicKey:KdlrLmLNVHDAC9tDvRuf_VQ9W_ASZ3C6J9mohBJ70Vc Fingerprint:firefox ServerName: SpiderX:/}} TcpSettings:{AcceptProxyProtocol:false Header:{Type:none}}} Sniffing:{Enabled:true DestOverride:[http tls quic fakedns] MetadataOnly:false RouteOnly:false} Allocate:{Strategy:always Refresh:5 Concurrency:3}}}
```
