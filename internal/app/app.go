package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	StatusActivated        = "Activated"
	StatusSettingsRequired = "SettingsRequired"
	StatusActivating       = "Activating"
	StatusInactive         = "Inactive"
)

type App struct {
	AppId       string `json:"app_id"`
	AccountId   string `json:"account_id"`
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
	Store       *store `json:"store"`
	InfoMessage string `json:"info_message"`
}

func loadApp(appId, accountId, accessToken string) *App {
	filePath := fmt.Sprintf("internal/app/%s.%s.app", appId, accountId)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		a := &App{AppId: appId,
			AccountId:   accountId,
			Status:      StatusActivated,
			AccessToken: accessToken,
			Store:       nil,
			InfoMessage: ""}
		data, _ := json.MarshalIndent(*a, "", "    ")
		if err = os.WriteFile(filePath, data, 0666); err != nil {
			log.Println(err)
			return nil
		}
		return a
	}
	a := new(App)
	if err = json.Unmarshal(file, a); err != nil {
		log.Println(err)
		return nil
	}
	a.Status = StatusActivated
	return a
}
