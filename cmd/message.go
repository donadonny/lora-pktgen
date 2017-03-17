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
	"strings"
)

// GetMsg - gets messages from the channel
func GetMsg(id string, startTime string, endTime string) string {
	url := UrlHTTP + "/channels/" + id +
	                 "/msg?start_time=" + startTime + "&end_time=" + endTime

	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// SendMsg - publishes SenML message on the channel
func SendMsg(id string, msg string) string {
	var err error

	url := UrlHTTP + "/channels/" + id + "/msg"
	sr := strings.NewReader(msg)
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	s := ""
	if resp.StatusCode == 202 {
		s = fmt.Sprintf("Message sent")
	} else {
		s = http.StatusText(resp.StatusCode)
	}

	return s
}
