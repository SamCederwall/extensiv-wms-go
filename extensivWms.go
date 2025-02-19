package extensivWms

import (
	"net/http"
)

const (
	baseUrl            string = "https://secure-wms.com"
	host               string = "secure-wms.com"
	connection         string = "keep-alive"
	contentType        string = "application/json"
	accept             string = "application/json"
	acceptEncoding     string = "gzip,deflate,sdch"
	acceptLanguage     string = "en-US,en;q=0.8"
	defaultHttpTimeout int    = 10
)

type Client struct {
	// The HTTP Client used to communicate with the Extensiv API
	Client *http.Client

	// Access token generated via the extensivWms/auth package or externally
	AccessToken string

	// The various endpoints / services
	Order OrderService
}
