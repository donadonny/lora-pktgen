/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/mainflux/mainflux-core/models"
)

// CreateDevice - creates new device and generates device UUID
func CreateDevice(msg string) string {
	url := UrlHTTP + "/devices"
	resp, err := netClient.Post(url, "application/json", nil)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// GetDevices - gets all devices
func GetDevices() string {
	url := UrlHTTP + "/devices"
	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// GetDevice - gets device by ID
func GetDevice(id string) string {
	url := UrlHTTP + "/devices/" + id
	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// UpdateDevice - updates device by ID
func UpdateDevice(id string, msg string) string {
	url := UrlHTTP + "/devices/" + id
	sr := strings.NewReader(msg)
	req, err := http.NewRequest("PUT", url, sr)
	if err != nil {
		return err.Error()
	}

	req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(msg)))

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// DeleteDevice - removes device
func DeleteDevice(id string) string {
	url := UrlHTTP + "/devices/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// DeleteAllDevices - removes all devices
func DeleteAllDevices() string {
	url := UrlHTTP + "/devices"
	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	var devices []models.Device
	json.Unmarshal([]byte(body), &devices)
	s := ""
	for i := 0; i < len(devices); i++ {
		s = s + DeleteDevice(devices[i].ID)
	}

	return s
}

// CreateDevice - creates new device and generates device UUID
func PlugDevice(id string, channels string) string {
	url := UrlHTTP + "/devices/" + id + "/plug"
	sr := strings.NewReader(channels)
	rsp, err := netClient.Post(url, "application/json", sr)
	body := GetHttpRespBody(rsp, err)

	return GetPrettyJson(body)
}
