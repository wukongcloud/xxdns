package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(25 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*20)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	return client
}

func HttpRequest(url string, method string) (data []byte, err error) {
	client := HttpClient()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if body, err := ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}
