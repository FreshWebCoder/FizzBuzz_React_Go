package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
	host := os.Getenv("API_HOST")
	//nolint:gosec // reason: insecure healthcheck is okay.
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	healthCheckURL := net.JoinHostPort(fmt.Sprintf("https://%s/health", host), os.Getenv("PORT"))
	// https://zenn.dev/spiegel/articles/20210125-http-get
	resp, err := url.Parse(healthCheckURL)
	if err != nil {
		return
	}
	//nolint:noctx // reason: healthcheck is okay.
	response, err := client.Get(resp.String())
	if err != nil {
		os.Exit(1)
	}
	response.Body.Close()
}
