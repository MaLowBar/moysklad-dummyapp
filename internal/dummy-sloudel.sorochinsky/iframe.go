package dummy

import (
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/jsonapi"
	"go-dummyapp-moysklad/internal/vendorapi"
)

//func (d *Dummy) HtmlHandler(c *gin.Context) {
//	contextKey := c.Query("contextKey")
//	userCtx := vendorapi.GetContext(d.AppUid, contextKey)
//
//	if userCtx == nil {
//		log.Println("Something bad happened while get context :(")
//	}
//	stores := jsonapi.Stores(d.AccessToken)
//
//
//	var st []app.Store
//	st = append(st, app.Store{})
//	if stores == nil {
//		log.Println("There has no stores or something bad happened while getting it :(")
//		stores = &st
//	}
//	var appIdx int
//	for i, a := range s.Apps {
//		if a.AppId == c.Param("appId") && a.AccountId == userCtx.accountId {
//			appIdx = i
//		}
//	}
//	var storeName string
//	if s.Apps[appIdx].Store == nil {
//		storeName = ""
//	} else {
//		storeName = s.Apps[appIdx].Store.Name
//	}
//
//	c.HTML(http.StatusOK, "iframe.html", gin.H{
//		"appId":            c.Param("appId"),
//		"uid":              userCtx.uid,
//		"fio":              userCtx.fio,
//		"accountId":        userCtx.accountId,
//		"isAdmin":          userCtx.isAdmin == "ALL",
//		"settingsRequired": s.Apps[appIdx].Status == app.StatusSettingsRequired,
//		"infoMsg":          s.Apps[appIdx].InfoMessage,
//		"currentStore":     storeName,
//		"stores":           stores,
//	})
//}

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

//func (s *app.Server) GetAppStatus(appId, accountId string) (string, error){
//	url := s.Config.MoyskladVendorApiEndpointUrl + "/apps/" + appId + "/" + accountId + "/status"
//	req := s.vendorRequest("GET", url, nil)
//	client := &http.Client{Timeout: 10 * time.Second}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println(err)
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode == 200 {
//		var status map[string]interface{}
//		json.NewDecoder(resp.Body).Decode(&status)
//		if st, ok := status["status"]; ok {
//			return st.(string), nil
//		}
//		if e, ok := status["error"]; ok {
//			return "", errors.New(e.(string))
//		}
//	}
//	return "", errors.New("unsuccessful request. Request status: " + resp.Status)
//}
