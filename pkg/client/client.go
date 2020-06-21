package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client   *http.Client
	baseAddr string
}

func NewClient(client *http.Client, baseAddr string) *Client {
	return &Client{client: client, baseAddr: baseAddr}
}

func (c *Client) Invoke(ctx context.Context, path string, in, out interface{}) error {

	bs, e := json.Marshal(in)

	if e != nil {
		panic(e)
	}

	res, err := c.client.Post(c.baseAddr+path, "application/json", bytes.NewReader(bs))

	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(r, out)

	if err != nil {
		panic(err)
	}

	return nil
}
