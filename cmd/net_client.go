package cmd

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	dialTout      = 5
	httpTout      = 10
	handshakeTout = 10
)

var (
	netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: dialTout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: handshakeTout * time.Second,
	}

	netClient = &http.Client{
		Timeout:   time.Second * httpTout,
		Transport: netTransport,
	}

	UrlHTTP = ""
	UrlAuth = ""
)

func SetServerAddr(HTTPHost string, HTTPPort int) {
	UrlHTTP = "http://" + HTTPHost + ":" + strconv.Itoa(HTTPPort)
}

func SetAuthServerAddr(AuthHost string, AuthPort int) {
	UrlAuth = "http://" + AuthHost + ":" + strconv.Itoa(AuthPort)
}
