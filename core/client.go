package core

import (
	"net/http"
	"net/url"
	"time"
)

func CreateCustomClient(proxy string, timeout int, useTransport bool) (*http.Client, error) {
	// create and return http client
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	if useTransport {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		// configure the http.Transport
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	return client, nil
}