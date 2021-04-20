package rchatwrapr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (c Client) SetUserStatus(message, status string) (err error) {
	err = c.SetUserToken()
	if err != nil {
		return
	}
	bd := map[string]string{
		"message": message,
	}
	if status != "" {
		bd["status"] = status
	}
	payload, _ := json.Marshal(bd)

	url := c.generateAPIURL("/users.setStatus")
	resp, err := c.makeRequest("POST", url, payload)
	if err != nil {
		fmt.Printf("Failed to set user status error: err %v", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	trgt := map[string]interface{}{}
	err = json.Unmarshal(body, &trgt)
	if err != nil {
		return
	}
	if s, ok := trgt["success"].(bool); !ok || !s {
		err = fmt.Errorf("rchat response not successful response %+v", trgt)
	}
	return
}
