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
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": currentApp.GetBaseApp().Status})
}

func (s *Server) DeleteHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	for i, a := range s.Apps {
		if a.GetBaseApp().AppId == c.Param("appId") && a.GetBaseApp().AccountId == c.Param("accountId") {
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
	userContext := vendorapi.GetContext(c.Param("appUid"), c.Query("contextKey"))
	if userContext == nil {
		log.Println("Something bad happened while get context :(")
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
			log.Println(err.Error())
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
				a.(*dummy.Dummy).UpdateSettings(c.PostForm("infoMessage"), c.PostForm("store"))
			}
		}
		c.Redirect(http.StatusFound, "/echo/iframe/"+c.Param("appUid"))
	default:
		log.Println("Invalid AppUid: " + c.Param("appUid"))
		c.Status(http.StatusInternalServerError)
		return
	}
}
