package connectwise

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient   *http.Client
	encodedCreds string
	clientId     string
}

type Creds struct {
	PublicKey  string
	PrivateKey string
	ClientId   string
	CompanyId  string // The company name you enter when you log in to the PSA
}

func NewClient(creds Creds, httpClient *http.Client) *Client {
	username := fmt.Sprintf("%s+%s", creds.CompanyId, creds.PublicKey)
	return &Client{
		httpClient:   httpClient,
		encodedCreds: basicAuth(username, creds.PrivateKey),
		clientId:     creds.ClientId,
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
