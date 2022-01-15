package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/dummy-sloudel.sorochinsky"
	"go-dummyapp-moysklad/internal/vendorapi"
	"log"
	"net/http"
)

type ActivateRequest struct {
	AppUID      string `json:"appUid"`
	AccountName string `json:"accountName"`
	Cause       string `json:"cause"`
	Access      []struct {
		Resource    string                 `json:"resource"`
		Scope       []string               `json:"scope"`
		Permissions map[string]interface{} `json:"permissions,omitempty"`
		AccessToken string                 `json:"access_token"`
	} `json:"access"`
}

func (s *Server) ActivateHandler(c *gin.Context) {

	var req ActivateRequest
	reqBody := c.Request.Body
	err := json.NewDecoder(reqBody).Decode(&req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	currentApp, err := s.addApp(c.Param("appId"), c.Param("accountId"), req)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": currentApp.GetBaseApp().Status})
}

func (s *Server) DeleteHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	for i, a := range s.Apps {
		if a.GetBaseApp().AppId == c.Param("appId") && a.GetBaseApp().AccountId == c.Param("accountId") {
			a.GetBaseApp().Status = app.StatusInactive
			upd := "UPDATE baseapp SET status = 'Inactive' WHERE app_id = $1 AND account_id = $2"
			_, err := s.db.Exec(upd, c.Param("appId"), c.Param("accountId"))
			if err != nil {
				log.Println(err)
			}
			s.Apps = append(s.Apps[:i], s.Apps[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func (s *Server) StatusHandler(c *gin.Context) {
	for _, a := range s.Apps {
		if a.GetBaseApp().AppId == c.Param("appId") && a.GetBaseApp().AccountId == c.Param("accountId") {
			c.JSON(http.StatusOK, gin.H{"status": a.GetBaseApp().Status})
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func (s *Server) IframeHandler(c *gin.Context) {
	userContext, err := vendorapi.GetContext(c.Param("appUid"), c.Query("contextKey"))
	if err != nil {
		log.Println(err)
	}
	var currentApp app.App
	for _, a := range s.Apps {
		if a.GetBaseApp().AppUid == c.Param("appUid") && a.GetBaseApp().AccountId == userContext.AccountId {
			currentApp = a
		}
	}

	switch c.Param("appUid") {
	case "dummy-sloudel.sorochinsky":
		htmlParams, err := currentApp.(*dummy.Dummy).RenderIframe(userContext)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.HTML(http.StatusOK, "iframe.html", htmlParams)
	default:
		log.Println("Invalid AppUid: " + c.Param("appUid"))
		c.Status(http.StatusInternalServerError)
		return
	}
}

func (s *Server) UpdateSettingsHandler(c *gin.Context) {
	switch c.Param("appUid") {
	case "dummy-sloudel.sorochinsky":
		for _, a := range s.Apps {
			if a.GetBaseApp().AppUid == c.Param("appUid") && a.GetBaseApp().AccountId == c.PostForm("accountId") {
				err := a.(*dummy.Dummy).UpdateSettings(c.PostForm("infoMessage"), c.PostForm("store"), s.db)
				if err != nil {
					log.Println(err)
					c.Status(http.StatusInternalServerError)
					return
				}
			}
		}
		c.Writer.Write([]byte("Настройки обновлены, перезагрузите приложение"))
	default:
		log.Println("Invalid AppUid: " + c.Param("appUid"))
		c.Status(http.StatusInternalServerError)
		return
	}
}