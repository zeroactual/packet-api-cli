package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

const (
	GET		= "GET"
	POST 	= "POST"
	DELETE 	= "DELETE"
)

func handle(url string, method string, body interface{}, obj interface{}) error {
	var marshalled []byte
	if body != nil {
		var err error
		marshalled, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	// Construct request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(marshalled))
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-Token", viper.GetString("project.token"))
	req.Header.Set("Content-Type", "application/json")

	var client http.Client
	resp, err := client.Do(req)

	// Check if token was good
	if resp.StatusCode == 401 {
		return errors.New("provided token is not valid, please try again")
	} else if 200 > resp.StatusCode && resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("error: %d", resp.StatusCode))
	}

	if obj != nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = resp.Body.Close()
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(bodyBytes), obj)
		if err != nil {
			return err
		}
	}

	return nil
}

func HandleErrs(errors ...error) bool {
	var errored bool
	for _, err := range errors {
		if err != nil {
			fmt.Printf("API call failed: %v\n", err)
			errored = true
		}
	}
	return errored
}