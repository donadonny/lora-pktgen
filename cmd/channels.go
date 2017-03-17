/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/mainflux/mainflux-core/models"
)

// CreateChannel - creates new channel and generates UUID
func CreateChannel(msg string) string {
	url := UrlHTTP + "/channels"
	resp, err := netClient.Post(url, "application/json", nil)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	s := ""
	if resp.StatusCode == 201 {
		s = fmt.Sprintf("Created resource %s", resp.Header.Get("Location"))
	} else {
		s = http.StatusText(resp.StatusCode)
	}

	return s
}

// GetChannels - gets all channels
func GetChannels(limit int) string {
	url := UrlHTTP + "/channels?climit=" + strconv.Itoa(limit)
	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// GetChannel - gets channel by ID
func GetChannel(id string) string {
	url := UrlHTTP + "/channels/" + id
	body := GetHttpRespBody(netClient.Get(url))

	return GetPrettyJson(body)
}

// UpdateChannel - publishes SenML message on the channel
func UpdateChannel(id string, msg string) string {
	url := UrlHTTP + "/channels/" + id
	sr := strings.NewReader(msg)
	req, err := http.NewRequest("PUT", url, sr)
	if err != nil {
		return err.Error()
	}

	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))

	body := GetHttpRespBody(netClient.Do(req))

	return GetPrettyJson(body)
}

// DeleteChannel - removes channel
func DeleteChannel(id string) string {
	url := UrlHTTP + "/channels/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}
	body := GetHttpRespBody(netClient.Do(req))

	return GetPrettyJson(body)
}

// DeleteAllChannels - removes all channels
func DeleteAllChannels() string {
	url := UrlHTTP + "/channels"
	body := GetHttpRespBody(netClient.Get(url))

	var channels []models.Channel
	json.Unmarshal([]byte(body), &channels)
	s := ""
	for i := 0; i < len(channels); i++ {
		s = s + DeleteChannel(channels[i].ID)
	}

	return s
}

// PlugChannel - plugs list of devices into the channel
func PlugChannel(id string, devices string) string {
	url := UrlHTTP + "/channels/" + id + "/plug"
	sr := strings.NewReader(devices)


	body := GetHttpRespBody(netClient.Post(url, "application/json", sr))

	return GetPrettyJson(body)
}
