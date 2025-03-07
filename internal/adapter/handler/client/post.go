package client

import (
	"bytes"
	"net/http"
)

func (c *ClientHttp) Post(url string, body []byte) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", c.BaseUrl+url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return buf.Bytes(), nil

}
