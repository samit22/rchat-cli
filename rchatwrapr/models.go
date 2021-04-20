package rchatwrapr

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var configFile = filepath.Join(os.Getenv("HOME"), ".rchat-config")

// RchatClient is used to communicate with rocket chat APIs
type Client struct {
	DomainURL string
	User      string
	Password  string
	userID    string
	authToken string
}

// RchatLoginRespone holds login information
type RchatLoginRespone struct {
	Status string `json:"status"`
	Data   struct {
		AuthToken string                 `json:"authToken"`
		UserID    string                 `json:"userId"`
		MyData    map[string]interface{} `json:"me"`
	} `json:"data"`
}

// SetAPIURL generates the API url
func (c Client) generateAPIURL(path string) string {
	return c.DomainURL + "/api/v1" + path
}

// SetUserToken sets the token for the request
func (c *Client) SetUserToken() (err error) {
	input, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Authentication failed please run init command error: %v", err)
		return
	}

	lines := strings.SplitAfter(string(input), "\n")
	totalLines := len(lines)

	lastLine := lines[totalLines-2]

	spltLines := strings.Split(lastLine, " ")
	if len(spltLines) < 3 {
		err = fmt.Errorf("token malformed err %v", err)
		fmt.Printf("Authentication failed please run init command error: token malformed")
		return
	}
	domainURL, rUserID, rAuthToken := spltLines[0], spltLines[1], spltLines[2]

	userID, err := decodeKey(string(rUserID))
	if err != nil {
		fmt.Printf("Authentication failed please run init command error: %v", err)
		return
	}
	authToken, err := decodeKey(string(rAuthToken))
	if err != nil {
		fmt.Printf("Authentication failed please run init command error: %v", err)
		return
	}
	c.DomainURL = string(domainURL)
	c.authToken = string(authToken)
	c.userID = string(userID)
	return
}

func decodeKey(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(key)
}
func encodeKey(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}
