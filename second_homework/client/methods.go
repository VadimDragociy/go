package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (client *Client) GetVersion() ([]byte, error) {
	request, err := http.NewRequest("GET", client.url+"/version", nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (client *Client) PostDecode(inputString string) (string, error) {
	requestBody, err := json.Marshal(inputString)

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", client.url+"/decode", bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}
	// request.Header.Set()
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var decodedString string

	err = json.Unmarshal(body, &decodedString)

	if err != nil {
		return "", err
	}

	return decodedString, nil
}

func (client *Client) GetHardOp() (bool, int, error) {
	cntxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	request, err := http.NewRequestWithContext(cntxt, "GET", client.url+"/hard-op", nil)

	if err != nil {
		return false, 500, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		if cntxt.Err() == context.DeadlineExceeded {
			return false, 500, cntxt.Err()
		}
		return false, 500, err
	}

	defer response.Body.Close()

	return true, response.StatusCode, nil

}
