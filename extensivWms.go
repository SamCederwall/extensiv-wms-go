package extensivWms

import "net/http"

const (
	Endpoint           string = "https://secure-wms.com/AuthServer/api/Token"
	Host               string = "secure-wms.com"
	Connection         string = "keep-alive"
	ContentType        string = "application/json"
	Accept             string = "application/json"
	AcceptEncoding     string = "gzip,deflate,sdch"
	AcceptLanguage     string = "en-US,en;q=0.8"
	DefaultHttpTimeout int    = 10
)

type Client struct {
	// The HTTP Client used to communicate with the Extensiv API
	Client      *http.Client
	AccessToken string
	Order       OrderService
}

type ResponseError struct {
	Status  int
	Message string
	Errors  []string
}
