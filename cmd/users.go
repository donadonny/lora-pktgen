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

// CreateChannel - creates new channel and generates UUID
func CreateUser(user string, pwd string) string {
	var err error
	msg := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user, pwd)

	sr := strings.NewReader(msg)

	url := UrlAuth + "/users"
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var s = fmt.Sprintf("POST/users: Status %d\n", resp.StatusCode)
	if resp.StatusCode == 201 {
		body := GetHttpRespBody(resp, err)
		s = fmt.Sprintf("%s %s", s, GetPrettyJson(body))
	} else {
		s = fmt.Sprintf("%s %s", s, http.StatusText(resp.StatusCode))
	}

	return s
}

// CreateChannel - creates new channel and generates UUID
func LogginInUser(user string, pwd string) string {
	var err error
	msg := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user, pwd)

	sr := strings.NewReader(msg)

	url := UrlAuth + "/sessions"
	resp, err := netClient.Post(url, "application/json", sr)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var s = fmt.Sprintf("POST/sessions: Status %d\n", resp.StatusCode)
	if resp.StatusCode == 201 {
		body := GetHttpRespBody(resp, err)
		s = fmt.Sprintf("%s %s", s, GetPrettyJson(body))
	} else {
		s = fmt.Sprintf("%s %s", s, http.StatusText(resp.StatusCode))
	}

	return s
}

// GetApiKeys - Retrieved a list of created keys
func GetApiKeys(auth string) string {
	url := UrlAuth + "/api-keys"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// CreateApiKeys - An API key can be given to the user, device or channel
func CreateApiKeys(auth string, owner string) string {
	url := UrlAuth + "/api-keys"

	req, err := http.NewRequest("POST", url, strings.NewReader(owner))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// DeleteApiKeys - Completely removes the key from the API key list
func DeleteApiKeys(auth string, key string) string {
	url := UrlAuth + "/api-keys/" + key

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// GetOwnerApiKeys - Get key owner
func GetOwnerApiKeys(auth string, key string) string {
	url := UrlAuth + "/api-keys/" + key
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
	req.Header.Set("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}

// UpdateOwnerApiKeys - Updates the key scope
func UpdateOwnerApiKeys(auth string, key string, owner string) string {
	url := UrlAuth + "/api-keys/" + key
	req, err := http.NewRequest("PUT", url, strings.NewReader(owner))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth))
	req.Header.Add("Content-Type", "application/json")

	resp, err := netClient.Do(req)
	body := GetHttpRespBody(resp, err)

	return GetPrettyJson(body)
}
