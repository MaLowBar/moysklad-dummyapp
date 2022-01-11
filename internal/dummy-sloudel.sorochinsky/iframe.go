package dummy

import (
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/jsonapi"
	"go-dummyapp-moysklad/internal/vendorapi"
)



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

func (d *Dummy) UpdateSettings(infoMessage, store string) {
	d.InfoMessage = infoMessage
	d.Store = store
}
