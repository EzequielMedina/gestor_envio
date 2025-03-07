package client

import (
	"bytes"
	"net/http"
)

func (c *ClientHttp) Get(url string) (string, error) {

	//hacemos la peticion a otro microservicio
	urlPeticion := c.BaseUrl + url
	client := &http.Client{}
	resp, err := client.Get(urlPeticion)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(resp.Body)

	return buf.String(), nil
}
