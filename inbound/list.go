package inbound

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0st1an/go3xui"
)

func List(c *go3xui.Client) (*InboundList, error) {
	req, err := http.NewRequest("GET", go3xui.EndpointInboundsList(c.URL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get inbounds: %w", err)
	}

	var inboundListR inboundList
	if err := json.NewDecoder(resp.Body).Decode(&inboundListR); err != nil {
		return nil, fmt.Errorf("failed to decode inbounds: %w", err)
	}

	inboundList := InboundList{
		Success: inboundListR.Success,
		Msg:     inboundListR.Msg,
		Obj:     make([]Inbound, 0),
	}

	for _, inboundItem := range inboundListR.Obj {
		var err error
		inbound := Inbound{
			ID:          inboundItem.ID,
			Up:          inboundItem.Up,
			Down:        inboundItem.Down,
			Total:       inboundItem.Total,
			Remark:      inboundItem.Remark,
			Enable:      inboundItem.Enable,
			ExpiryTime:  inboundItem.ExpiryTime,
			ClientStats: inboundItem.ClientStats,
			Listen:      inboundItem.Listen,
			Port:        inboundItem.Port,
			Protocol:    inboundItem.Protocol,
			Tag:         inboundItem.Tag,
		}

		err = decodeStringToStruct(inboundItem.Settings, &inbound.Settings)
		if err != nil {
			return nil, fmt.Errorf("failed to decode inbound settings: %w", err)
		}

		err = decodeStringToStruct(inboundItem.StreamSettings, &inbound.StreamSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to decode inbound stream settings: %w", err)
		}

		err = decodeStringToStruct(inboundItem.Sniffing, &inbound.Sniffing)
		if err != nil {
			return nil, fmt.Errorf("failed to decode inbound sniffing: %w", err)
		}

		err = decodeStringToStruct(inboundItem.Allocate, &inbound.Allocate)
		if err != nil {
			return nil, fmt.Errorf("failed to decode inbound allocate: %w", err)
		}

		inboundList.Obj = append(inboundList.Obj, inbound)
	}

	return &inboundList, nil
}
