package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getJson(client *http.Client, url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 {
			return fmt.Errorf("Got 403 Unauthorized: Please check the API key")
		}
		return fmt.Errorf("Recived a bad status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func sendJson(client *http.Client, method, url string, payload, response interface{}) error {
	// marshal User to json
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 {
			return fmt.Errorf("Got 403 Unauthorized: Please check the API key")
		}
		return fmt.Errorf("Recived a bad status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
