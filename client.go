package extensivWms

type Client struct {
	AccessToken string
}

func New(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}
}
