package inbound

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/k0st1an/go3xui"
)

func Add(c *go3xui.Client, inbound Inbound) (*InboundAdd, error) {
	form := url.Values{}
	form.Add("up", strconv.FormatInt(inbound.Up, 10))
	form.Add("down", strconv.FormatInt(inbound.Down, 10))
	form.Add("total", strconv.FormatInt(inbound.Total, 10))
	form.Add("remark", inbound.Remark)
	form.Add("enable", strconv.FormatBool(inbound.Enable))
	form.Add("expiryTime", strconv.FormatInt(inbound.ExpiryTime, 10))
	form.Add("listen", inbound.Listen)
	form.Add("port", strconv.FormatInt(inbound.Port, 10))
	form.Add("protocol", inbound.Protocol)

	if inbound.Settings.Decryption == "" {
		inbound.Settings.Decryption = "none"
	}

	settings, err := json.Marshal(inbound.Settings)
	if err != nil {
		return nil, err
	}
	form.Add("settings", string(settings))

	streamSettings, err := json.Marshal(inbound.StreamSettings)
	if err != nil {
		return nil, err
	}
	form.Add("streamSettings", string(streamSettings))

	sniffing, err := json.Marshal(inbound.Sniffing)
	if err != nil {
		return nil, err
	}
	form.Add("sniffing", string(sniffing))

	allocate, err := json.Marshal(inbound.Allocate)
	if err != nil {
		return nil, err
	}
	form.Add("allocate", string(allocate))

	resp, err := c.HttpClient.PostForm(go3xui.EndpointInboundsAdd(c.URL), form)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to add inbound: %w", err)
	}

	var inboundAddTmp inboundAdd
	if err := json.NewDecoder(resp.Body).Decode(&inboundAddTmp); err != nil {
		return nil, fmt.Errorf("failed to decode inbounds: %w", err)
	}

	var inboundAdd InboundAdd
	inboundAdd.Success = inboundAddTmp.Success
	inboundAdd.Msg = inboundAddTmp.Msg
	inboundAdd.Obj = Inbound{
		ID:          inboundAddTmp.Obj.ID,
		Up:          inboundAddTmp.Obj.Up,
		Down:        inboundAddTmp.Obj.Down,
		Total:       inboundAddTmp.Obj.Total,
		Remark:      inboundAddTmp.Obj.Remark,
		Enable:      inboundAddTmp.Obj.Enable,
		ExpiryTime:  inboundAddTmp.Obj.ExpiryTime,
		Listen:      inboundAddTmp.Obj.Listen,
		Port:        inboundAddTmp.Obj.Port,
		Protocol:    inboundAddTmp.Obj.Protocol,
		Tag:         inboundAddTmp.Obj.Tag,
		ClientStats: inboundAddTmp.Obj.ClientStats,
	}

	err = decodeStringToStruct(inboundAddTmp.Obj.Settings, &inboundAdd.Obj.Settings)
	if err != nil {
		return nil, fmt.Errorf("failed to decode inbound settings: %w", err)
	}

	err = decodeStringToStruct(inboundAddTmp.Obj.StreamSettings, &inboundAdd.Obj.StreamSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to decode inbound stream settings: %w", err)
	}

	err = decodeStringToStruct(inboundAddTmp.Obj.Sniffing, &inboundAdd.Obj.Sniffing)
	if err != nil {
		return nil, fmt.Errorf("failed to decode inbound sniffing: %w", err)
	}

	err = decodeStringToStruct(inboundAddTmp.Obj.Allocate, &inboundAdd.Obj.Allocate)
	if err != nil {
		return nil, fmt.Errorf("failed to decode inbound allocate: %w", err)
	}

	return &inboundAdd, nil
}
