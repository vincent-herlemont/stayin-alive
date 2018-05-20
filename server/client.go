package server

import (
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
)


type Client struct {
	cfg *CfgWebServer
	client http.Client
}

func (c *Client) getUri(path string) string {
	url := url.URL{
		Scheme: "http",
		Host:   c.cfg.Addr,
		Path:   path,
	}
	return url.String()
}

func (c *Client) decodeTo(resp *http.Response,to interface{}) error {
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(to)
}

func (c *Client) Get(path string) (resp *http.Response, err error) {
	return c.client.Get(c.getUri(path));
}

func (c *Client) GetTO(path string,to interface{}) (error) {
	resp,err := c.Get(path)
	if err != nil {
		return err
	}
	return c.decodeTo(resp,to)
}

func (c *Client) Post(path string, jsonBody string) (resp *http.Response, err error) {
	body := strings.NewReader(jsonBody)
	return c.client.Post(c.getUri(path),"application/json", body);
}

func (c *Client) PostTO(path string,inputTo interface{}, outputTo interface{}) (error) {
	jsonBody,err := json.Marshal(inputTo)
	if err != nil {
		return err
	}
	resp,err := c.Post(path,string(jsonBody))
	if err != nil {
		return err
	}
	return c.decodeTo(resp,outputTo)
}
