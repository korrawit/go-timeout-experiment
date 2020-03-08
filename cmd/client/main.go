package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	mode              = envDefaultString(os.Getenv("MODE"), "DEFAULT_CLIENT")
	respHeaderTimeout = envDefaultDuration(os.Getenv("RESP_HEADER_TIMEOUT"), time.Duration(3*time.Second))
)

func main() {
	c, err := createHttpClient(mode)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running on:", mode)
	if mode == "CUSTOM_CLIENT" {
		log.Println("Timeout:", respHeaderTimeout)
	}

	resp, err := c.Get("http://localhost:8080/api/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("resp:", string(b))
}

func createHttpClient(mode string) (*http.Client, error) {
	var c *http.Client
	switch mode {
	case "DEFAULT_CLIENT":
		c = http.DefaultClient
	case "CUSTOM_CLIENT":
		c = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				ResponseHeaderTimeout: respHeaderTimeout,
			},
		}
	default:
		return nil, errors.New("Unsupported http client mode")
	}

	return c, nil
}

func envDefaultString(s string, d string) string {
	if s == "" {
		return d
	}
	return s
}

func envDefaultDuration(s string, d time.Duration) time.Duration {
	v, err := time.ParseDuration(s)
	if err != nil {
		return d
	}
	return v
}
