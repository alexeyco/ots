package ots

import (
	"fmt"
	"net/url"
	"time"
)

// ClientOption client option.
type ClientOption func(*Client)

// WithHTTPClient to set custom HTTP client.
func WithHTTPClient(httpClient HTTPClient) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// ShareOption share option.
type ShareOption func(url.Values)

// SharePassphrase share a secret with a passphrase.
func SharePassphrase(passPhrase string) ShareOption {
	return func(v url.Values) {
		v.Set("passphrase", passPhrase)
	}
}

// ShareTTL share a secret with TTL.
func ShareTTL(ttl time.Duration) ShareOption {
	return func(v url.Values) {
		v.Set("ttl", fmt.Sprintf("%d", int64(ttl.Seconds())))
	}
}

// ShareRecipient share a secret and send it to recipient.
func ShareRecipient(recipient string) ShareOption {
	return func(v url.Values) {
		v.Set("recipient", recipient)
	}
}

// GenerateOption generate option.
type GenerateOption func(url.Values)

// GeneratePassphrase generate a secret with a passphrase.
func GeneratePassphrase(passPhrase string) GenerateOption {
	return func(v url.Values) {
		v.Set("passphrase", passPhrase)
	}
}

// GenerateTTL generate a secret with TTL.
func GenerateTTL(ttl time.Duration) GenerateOption {
	return func(v url.Values) {
		v.Set("ttl", fmt.Sprintf("%d", int64(ttl.Seconds())))
	}
}

// GenerateRecipient generate a secret and send it to recipient.
func GenerateRecipient(recipient string) GenerateOption {
	return func(v url.Values) {
		v.Set("recipient", recipient)
	}
}

// SecretOption secret option.
type SecretOption func(url.Values)

// SecretPassphrase set a passphrase to receive a secret.
func SecretPassphrase(passPhrase string) SecretOption {
	return func(v url.Values) {
		v.Set("passphrase", passPhrase)
	}
}
