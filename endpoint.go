package go3xui

import "fmt"

func EndpointLogin(url string) string {
	return fmt.Sprintf("%s/login", url)
}

func EndpointServerStatus(url string) string {
	return fmt.Sprintf("%s/server/status", url)
}

func EndpointGetNewX25519Cert(url string) string {
	return fmt.Sprintf("%s/server/getNewX25519Cert", url)
}

func EndpointInboundsList(url string) string {
	return fmt.Sprintf("%s/panel/api/inbounds/list", url)
}

func EndpointInboundsAdd(url string) string {
	return fmt.Sprintf("%s/panel/api/inbounds/add", url)
}
