package ots

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

// Secret value.
type Secret struct {
	SecretKey string `json:"secret_key"`
	Value     string `json:"value"`
}

// Metadata of a secret.
type Metadata struct {
	CustID             string   `json:"custid"`
	Value              *string  `json:"value,omitempty"`
	MetadataKey        string   `json:"metadata_key"`
	SecretKey          string   `json:"secret_key"`
	TTL                Duration `json:"ttl"`
	MetadataTTL        Duration `json:"metadata_ttl,omitempty"`
	SecretTTL          Duration `json:"secret_ttl,omitempty"`
	State              State    `json:"state"`
	Recipient          []string `json:"recipient,omitempty"`
	PassphraseRequired bool     `json:"passphrase_required"`
	Created            Time     `json:"created"`
	Updated            Time     `json:"updated"`
	Received           *Time    `json:"received,omitempty"`
}
