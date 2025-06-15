package main

import (
	"log"

	"github.com/k0st1an/go3xui"
	"github.com/k0st1an/go3xui/inbound"
)

func main() {
	apiClient, err := go3xui.NewClient("https://...")
	if err != nil {
		log.Fatal(err)
	}

	if err := apiClient.Login("username", "password"); err != nil {
		log.Fatal(err)
	}

	addClient := inbound.AddClient{
		ID: 1,
		Settings: inbound.AddClientSettings{
			Client: []inbound.AddClientClient{
				{
					ID:         go3xui.GenerateUUID(),
					Flow:       "xtls-rprx-vision",
					Email:      go3xui.RandomHexString(9),
					LimitIP:    0,
					TotalGB:    0,
					ExpiryTime: 0,
					Enable:     true,
					TgId:       "",
					SubId:      go3xui.RandomHexString(9),
					Comment:    "",
					Reset:      0,
				},
			},
		},
	}

	addClientResponse, err := addClient.AddClient(apiClient)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Add client response:")
	log.Printf("%+v", addClientResponse)
}

// Output:
// 2025/06/15 23:52:28 main.go:56: Add client response:
// 2025/06/15 23:52:28 main.go:57: &{Success:true Msg:Client(s) added Successfully Obj:<nil>}
