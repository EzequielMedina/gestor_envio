package client

type ClientHttp struct {
	BaseUrl string
}

func NewClientHttp(baseUrl string) *ClientHttp {
	return &ClientHttp{BaseUrl: baseUrl}
}
