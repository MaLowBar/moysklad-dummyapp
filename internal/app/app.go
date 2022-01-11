package app

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	StatusActivated        = "Activated"
	StatusSettingsRequired = "SettingsRequired"
	StatusActivating       = "Activating"
	StatusInactive         = "Inactive"
)

type App interface {
	GetBaseApp() *BaseApp
}

type BaseApp struct {
	AppId       string `json:"app_id"`
	AppUid      string `json:"app_uid"`
	AccountId   string `json:"account_id"`
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
	SecretKey   string `json:"secret_key"`
}

func (ba *BaseApp) GetBaseApp() *BaseApp {
	return ba
}

func LoadBaseApp(appId, appUid, accountId, accessToken string) (*BaseApp, error) {
	filePath := fmt.Sprintf("internal/app/%s.%s.app", appUid, accountId)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		secretKey, _ := os.ReadFile("internal/" + appUid + "/secret.key")
		a := BaseApp{AppId: appId,
			AppUid:      appUid,
			AccountId:   accountId,
			Status:      StatusActivated,
			AccessToken: accessToken,
			SecretKey:   string(secretKey)}
		data, _ := yaml.Marshal(a)
		if err = os.WriteFile(filePath, data, 0666); err != nil {
			return &BaseApp{}, err
		}
		return &a, nil
	}
	a := BaseApp{}
	if err = json.Unmarshal(file, &a); err != nil {
		return &BaseApp{}, err
	}
	a.Status = StatusActivated
	return &a, nil
}
