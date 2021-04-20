package rchatwrapr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (c Client) ValidateRequest() (err error) {
	url := c.generateAPIURL("/login")

	payload, _ := json.Marshal(map[string]string{
		"user":     c.User,
		"password": c.Password,
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Printf("failed to create http request err: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making http request err: %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading auth body body error %v", err)
		return
	}
	defer resp.Body.Close()

	trgt := RchatLoginRespone{}
	err = json.Unmarshal(body, &trgt)
	if err != nil {
		return
	}

	if trgt.Status != "success" {
		fmt.Printf("Response %v", trgt)
		err = errors.New("failed to fetch auth information check the auth parameters")
		fmt.Println(err)
		return err
	}
	fmt.Printf("Successfully validated the request!!")
	f, err := os.OpenFile(configFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	f.Chmod(0600)
	line := fmt.Sprintf("%s %s %s", c.DomainURL, encodeKey(trgt.Data.UserID), encodeKey(trgt.Data.AuthToken))
	line += "\n"

	_, err = f.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failed to add login information %v", err)
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Printf("Failed to  write auth into file %v", err)
	}
	return
}
