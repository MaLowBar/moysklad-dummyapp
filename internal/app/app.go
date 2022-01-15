package app

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
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

func LoadBaseApp(appId, appUid, accountId, accessToken string, db *sql.DB) (*BaseApp, error) {
	var (
		a BaseApp
		aId int
	)
	row := db.QueryRow(`SELECT * FROM baseapp WHERE app_id = $1 AND account_id = $2`, appId, accountId)
	err := row.Scan(&aId, &a.AppId, &a.AppUid, &a.AccountId, &a.Status, &a.AccessToken, &a.SecretKey)

	if err != nil && err != sql.ErrNoRows{
		return nil, err
	} else if err == sql.ErrNoRows {
		secretKey, _ := os.ReadFile("internal/" + appUid + "/secret.key")
		ins := "INSERT INTO baseapp (app_id, app_uid, account_id, status, access_token, secret_key) VALUES ($1, $2, $3, $4, $5, $6)"
		_, err = db.Exec(ins, appId, appUid, accountId, StatusSettingsRequired, accessToken, string(secretKey))
		if err != nil {
			return nil, err
		}

		a.AppId = appId
		a.AppUid = appUid
		a.AccountId = accountId
		a.Status = StatusSettingsRequired
		a.AccessToken = accessToken
		a.SecretKey = string(secretKey)
		return &a, nil
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	if a.Status == StatusInactive {
		a.Status = StatusSettingsRequired
		a.AccessToken = accessToken
		_, err := db.Exec("UPDATE baseapp SET status = 'SettingsRequired', access_token = $1 WHERE id = $2", accessToken, aId)
		if err != nil {
			return &a, err
		}
	}
	return &a, nil
}
