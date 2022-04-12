package ots

const (
	// Offline means API is offline.
	Offline Status = iota
	// Nominal means API works fine.
	Nominal
)

const (
	// StateNew means that secret is not received or viewed.
	StateNew = "new"
	// StateViewed means that secret metadata is already viewed.
	StateViewed = "viewed"
	// StateReceived means that secret is received.
	StateReceived = "received"
)

const endpoint = "https://onetimesecret.com/api/v1"
