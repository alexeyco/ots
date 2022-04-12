package ots

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client of service.
type Client struct {
	userName   string
	apiKey     string
	httpClient HTTPClient
}

// New returns new client instance.
func New(userName, apiKey string, opts ...ClientOption) *Client {
	c := Client{
		userName:   userName,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}

	for _, o := range opts {
		o(&c)
	}

	return &c
}

// Status returns API status.
func (c *Client) Status(ctx context.Context) (Status, error) {
	var (
		res    statusResponse
		status Status
	)

	if err := c.requestJSON(ctx, http.MethodGet, "status", &res); err != nil {
		return status, err
	}

	if res.Status == "nominal" {
		status = Nominal
	}

	return status, nil
}

// Share a secret.
func (c *Client) Share(ctx context.Context, secret string, opts ...ShareOption) (Metadata, error) {
	form := url.Values{
		"secret": []string{secret},
	}

	for _, o := range opts {
		o(form)
	}

	var res Metadata
	if err := c.requestJSON(ctx, http.MethodPost, "share", &res, form); err != nil {
		return Metadata{}, err
	}

	return res, nil
}

// Generate a secret.
func (c *Client) Generate(ctx context.Context, opts ...GenerateOption) (Metadata, error) {
	form := url.Values{}
	for _, o := range opts {
		o(form)
	}

	var res Metadata
	if err := c.requestJSON(ctx, http.MethodPost, "generate", &res, form); err != nil {
		return Metadata{}, err
	}

	return res, nil
}

// Secret returns secret by secret key.
func (c *Client) Secret(ctx context.Context, secretKey string, opts ...SecretOption) (Secret, error) {
	form := url.Values{}
	for _, o := range opts {
		o(form)
	}

	var res Secret
	if err := c.requestJSON(ctx, http.MethodPost, fmt.Sprintf("secret/%s", secretKey), &res, form); err != nil {
		return Secret{}, err
	}

	return res, nil
}

// Metadata returns metadata by metadata key.
func (c *Client) Metadata(ctx context.Context, metadataKey string) (Metadata, error) {
	var res Metadata
	if err := c.requestJSON(ctx, http.MethodPost, fmt.Sprintf("private/%s", metadataKey), &res, url.Values{}); err != nil {
		return Metadata{}, err
	}

	return res, nil
}

// Burn a secret by metadata key.
func (c *Client) Burn(ctx context.Context, metadataKey string) error {
	return c.requestJSON(ctx, http.MethodPost, fmt.Sprintf("private/%s/burn", metadataKey), nil, url.Values{})
}

func (c *Client) requestJSON(ctx context.Context, method string, resource string, v interface{}, form ...url.Values) error {
	var body io.Reader
	if len(form) > 0 {
		body = strings.NewReader(form[0].Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", endpoint, resource), body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.userName, c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return &Error{
			StatusCode: res.StatusCode,
		}
	}

	err = c.parseJSON(res, v)

	return err
}

func (c *Client) parseJSON(res *http.Response, v interface{}) error {
	if v == nil {
		return nil
	}

	d := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		var response errorResponse
		if err := d.Decode(&response); err != nil {
			return err
		}

		return &Error{
			Message:    response.Message,
			StatusCode: res.StatusCode,
		}
	}

	err := d.Decode(v)

	return err
}
