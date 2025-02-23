package extensivWms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	baseUrl            string = "https://secure-wms.com"
	host               string = "secure-wms.com"
	connection         string = "keep-alive"
	contentType        string = "application/json"
	accept             string = "application/json"
	acceptEncoding     string = "gzip,deflate,sdch"
	acceptLanguage     string = "en-US,en;q=0.8"
	defaultHttpTimeout        = 10
)

type HttpError struct {
	Status  int
	Message interface{}
}

func (e HttpError) GetStatus() int {
	return e.Status
}

func (e HttpError) GetMessage() string {
	messageBytes, err := json.Marshal(e.Message)

	if err != nil {
		return ""
	}

	return string(messageBytes)
}

func (e HttpError) Error() string {
	if e.Message != nil {
		return e.GetMessage()
	}

	return "Unknown Error"
}

type Client struct {
	// The HTTP Client used to communicate with the Extensiv API
	Client *http.Client

	// Temporary access token
	accessToken string

	// Base URL for making requests to the extensiv API.
	baseUrl *url.URL

	// Constants requried by extensiv to be set as headers
	host           string
	connection     string
	contentType    string
	accept         string
	acceptEncoding string
	acceptLanguage string

	// The various endpoints / services
	Order OrderService
}

func NewClient(accessToken string) (*Client, error) {

	baseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	cli := &Client{
		Client: &http.Client{
			Timeout: time.Second * defaultHttpTimeout,
		},
		baseUrl:        baseUrl,
		accessToken:    accessToken,
		host:           host,
		connection:     connection,
		contentType:    contentType,
		accept:         accept,
		acceptEncoding: acceptEncoding,
		acceptLanguage: acceptLanguage,
	}

	cli.Order = &OrderServiceOp{client: cli}

	return cli, nil

}

func (c *Client) BuildHeaderMap() map[string]string {

	headers := make(map[string]string)

	headers["host"] = c.host
	headers["content-type"] = c.contentType
	headers["accept"] = c.accept
	headers["accept-language"] = c.acceptLanguage
	headers["authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)

	return headers
}

func (c *Client) DoGet(req http.Request, v interface{}) error {

	res, err := c.Client.Do(&req)

	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("request to extensiv failed with status code %d", res.StatusCode)
	}

	if v != nil {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Get(ctx context.Context, relPath string, resource interface{}, opts interface{}) error {
	return c.BuildAndMakeRequest(ctx, "GET", relPath, nil, opts, resource)
}

func (c *Client) BuildAndMakeRequest(
	ctx context.Context,
	method string,
	relPath string,
	body io.Reader,
	opts interface{},
	resource interface{},
) error {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}

	req, err := c.BuildRequest(ctx, method, relPath, body, opts)

	if err != nil {
		return err
	}

	c.MakeRequest(req, resource)

	return nil
}

func (c *Client) MakeRequest(req *http.Request, resource interface{}) (http.Header, error) {

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkHttpResponse(res)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	if resource != nil {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(resource)

		if err != nil {
			return nil, err
		}
	}

	return res.Header, nil

}

func checkHttpResponse(r *http.Response) error {
	var parsedBody interface{}

	if http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	resBody, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &parsedBody)

	if err != nil {
		return err
	}

	return HttpError{
		Status:  r.StatusCode,
		Message: &parsedBody,
	}
}

func (c *Client) BuildRequest(
	ctx context.Context,
	method string,
	relPath string,
	body interface{},
	options interface{},
) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	endpoint := c.baseUrl.ResolveReference(rel)

	if options != nil {
		optionsQuery, err := query.Values(options)

		if err != nil {
			return nil, err
		}

		for k, values := range endpoint.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}

		}
	}

	var payload []byte = nil
	if body != nil {
		payload, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	headers := c.BuildHeaderMap()

	req, err := http.NewRequest(method, endpoint.String(), bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, nil
}
