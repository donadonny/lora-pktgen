/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

// Status - server health check
func Status() string {
	url := UrlHTTP + "/status"
	body := GetHttpRespBody(netClient.Get(url))

	return GetPrettyJson(body)
}
