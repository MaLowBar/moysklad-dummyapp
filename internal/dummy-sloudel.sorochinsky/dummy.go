package dummy

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/jsonapi"
	"go-dummyapp-moysklad/vendorapi"
)

type Dummy struct {
	app.BaseApp
	InfoMessage string `yaml:"info_message"`
	Store       string `yaml:"store"`
}

func (d *Dummy) GetBaseApp() *app.BaseApp {
	return &d.BaseApp
}

func LoadDummyApp(baseApp app.BaseApp, db *sql.DB) (*Dummy, error) {
	da := Dummy{BaseApp: baseApp}
	query := "SELECT info_message, store FROM dummyapp WHERE baseapp_id = (SELECT id FROM baseapp WHERE app_id = $1 AND account_id = $2)"
	row := db.QueryRow(query, baseApp.AppId, baseApp.AccountId)
	err := row.Scan(&da.InfoMessage, &da.Store)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		ins := "INSERT INTO dummyapp VALUES ((SELECT id FROM baseapp WHERE app_id = $1 AND account_id = $2), $3, $4)"
		_, err = db.Exec(ins, baseApp.AppId, baseApp.AccountId, "", "")
		if err != nil {
			return nil, err
		}
		return &da, nil
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return &da, nil
}

func (d *Dummy) RenderIframe(userCtx *vendorapi.UserContext) (map[string]interface{}, error) {
	stores, err := jsonapi.Stores(d.AccessToken)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"appUid":           d.AppUid,
		"uid":              userCtx.Uid,
		"fio":              userCtx.Fio,
		"accountId":        userCtx.AccountId,
		"isAdmin":          userCtx.IsAdmin == "ALL",
		"settingsRequired": d.Status == app.StatusSettingsRequired,
		"infoMsg":          d.InfoMessage,
		"currentStore":     d.Store,
		"stores":           stores,
	}, nil
}

func (d *Dummy) UpdateSettings(infoMessage, store string, db *sql.DB) error {
	if d.Status == app.StatusSettingsRequired {
		d.Status = app.StatusActivated
		upd := "UPDATE baseapp SET status = 'Activated' WHERE app_id = $1 AND account_id = $2"
		_, err := db.Exec(upd, d.AppId, d.AccountId)
		if err != nil {
			return err
		}
	}
	upd := "UPDATE dummyapp SET info_message = $1, store = $2 WHERE baseapp_id = (SELECT id FROM baseapp WHERE app_uid = $3 AND account_id = $4)"
	_, err := db.Exec(upd, infoMessage, store, d.AppUid, d.AccountId)
	if err != nil {
		return err
	}
	d.InfoMessage = infoMessage
	d.Store = store
	return nil
}
