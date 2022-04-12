package ots_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/ots"
)

func TestSharePassphrase(t *testing.T) {
	passphrase := "passphrase"

	expected := url.Values{
		"passphrase": []string{
			passphrase,
		},
	}

	actual := url.Values{}
	ots.SharePassphrase(passphrase)(actual)

	assert.Equal(t, expected, actual)
}

func TestShareTTL(t *testing.T) {
	ttl := 123 * time.Second

	expected := url.Values{
		"ttl": []string{
			fmt.Sprintf("%d", int64(ttl.Seconds())),
		},
	}

	actual := url.Values{}
	ots.ShareTTL(ttl)(actual)

	assert.Equal(t, expected, actual)
}

func TestShareRecipient(t *testing.T) {
	recipient := "foo@bar.com"

	expected := url.Values{
		"recipient": []string{
			recipient,
		},
	}

	actual := url.Values{}
	ots.ShareRecipient(recipient)(actual)

	assert.Equal(t, expected, actual)
}

func TestGeneratePassphrase(t *testing.T) {
	passphrase := "passphrase"

	expected := url.Values{
		"passphrase": []string{
			passphrase,
		},
	}

	actual := url.Values{}
	ots.GeneratePassphrase(passphrase)(actual)

	assert.Equal(t, expected, actual)
}

func TestGenerateTTL(t *testing.T) {
	ttl := 123 * time.Second

	expected := url.Values{
		"ttl": []string{
			fmt.Sprintf("%d", int64(ttl.Seconds())),
		},
	}

	actual := url.Values{}
	ots.GenerateTTL(ttl)(actual)

	assert.Equal(t, expected, actual)
}

func TestGenerateRecipient(t *testing.T) {
	recipient := "foo@bar.com"

	expected := url.Values{
		"recipient": []string{
			recipient,
		},
	}

	actual := url.Values{}
	ots.GenerateRecipient(recipient)(actual)

	assert.Equal(t, expected, actual)
}

func TestSecretPassphrase(t *testing.T) {
	passphrase := "passphrase"

	expected := url.Values{
		"passphrase": []string{
			passphrase,
		},
	}

	actual := url.Values{}
	ots.SecretPassphrase(passphrase)(actual)

	assert.Equal(t, expected, actual)
}
