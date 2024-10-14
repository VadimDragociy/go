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
	request, request_err := http.NewRequest("GET", client.url+"/version", nil)

	if request_err != nil {
		return nil, request_err
	}

	response, response_err := http.DefaultClient.Do(request)

	if response_err != nil {
		return nil, response_err
	}

	defer response.Body.Close()

	body, read_err := io.ReadAll(response.Body)
	if read_err != nil {
		return nil, read_err
	}

	return body, nil

}

func (client *Client) PostDecode(inputString string) (string, error) {
	requestBody, encoding_err := json.Marshal(inputString)

	if encoding_err != nil {
		return "", encoding_err
	}

	request, request_err := http.NewRequest("POST", client.url+"/decode", bytes.NewBuffer(requestBody))

	if request_err != nil {
		return "", request_err
	}
	// request.Header.Set()
	response, response_err := http.DefaultClient.Do(request)

	if response_err != nil {
		return "", response_err
	}
	defer response.Body.Close()

	body, read_err := io.ReadAll(response.Body)
	if read_err != nil {
		return "", read_err
	}

	var decoded_str string

	decode_err := json.Unmarshal(body, &decoded_str)

	if decode_err != nil {
		return "", decode_err
	}

	return decoded_str, nil
}

func (client *Client) GetHardOp() (bool, int, error) {
	cntxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	request, request_err := http.NewRequestWithContext(cntxt, "GET", client.url+"/hard-op", nil)

	if request_err != nil {
		return false, 500, request_err
	}

	response, response_err := http.DefaultClient.Do(request)

	if response_err != nil {
		if cntxt.Err() == context.DeadlineExceeded {
			return false, 500, cntxt.Err()
		}
		return false, 500, response_err
	}

	defer response.Body.Close()

	return true, response.StatusCode, nil

}
