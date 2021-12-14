package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type store struct {
	Name string
}

func (s *Server) jsonApiStores() *[]store {
	client := &http.Client{Timeout: 10 * time.Second}

	url := "http://localhost:8080/store" //s.Config.MoyskladJsonApiEndpointUrl + "/entity/store"
	req := request("GET", url, s.Apps[0].AccessToken, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	var rawStores map[string]interface{}
	if resp.Status == "200 OK" {
		if err = json.NewDecoder(resp.Body).Decode(&rawStores); err != nil {
			log.Println(err)
			return nil
		}
		stores := new([]store)
		for _, st := range rawStores["rows"].([]interface{}) {
			*stores = append(*stores, store{st.(map[string]interface{})["name"].(string)})
		}
		return stores
	}
	return nil
}
