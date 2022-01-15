package jsonapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)


const jsonBaseUrl = "https://online.moysklad.ru/api/remap/1.2"

func Request(method, url, bearerToken string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer " + bearerToken},
		"Content-Type":  {"application/json"},
	}
	return req
}

func Stores(accessToken string) ([]string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := jsonBaseUrl + "/entity/store" //"http://localhost:8080/store"
	req := Request("GET", url, accessToken, nil)

	var stores []string

	resp, err := client.Do(req)
	if err != nil {
		return stores, err
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		var rawStores map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&rawStores); err != nil {
			return stores, err
		}
		for _, st := range rawStores["rows"].([]interface{}) {
			stores = append(stores, st.(map[string]interface{})["name"].(string))
		}
		return stores, nil
	}
	return stores, errors.New(resp.Status)
}
