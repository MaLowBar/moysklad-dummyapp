package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *Server) HtmlHandler(c *gin.Context) {
	contextKey := c.Query("contextKey")
	userCtx := s.getContext(contextKey)
	stores := s.jsonApiStores()

	if userCtx == nil {
		c.Status(http.StatusInternalServerError)
		log.Println("Something bad happened while get context :(")
		return
	}
	if stores == nil {
		c.Status(http.StatusInternalServerError)
		log.Println("Something bad happened while get stores :(")
		return
	}
	var appIdx int
	for i, a := range s.Apps {
		if a.AppId == c.Param("appId") && a.AccountId == userCtx.accountId {
			appIdx = i
		}
	}
	var storeName string
	if s.Apps[appIdx].Store == nil {
		storeName = ""
	} else {
		storeName = s.Apps[appIdx].Store.Name
	}

	c.HTML(http.StatusOK, "iframe.html", gin.H{
		"appId":            c.Param("appId"),
		"uid":              userCtx.uid,
		"fio":              userCtx.fio,
		"accountId":        userCtx.accountId,
		"isAdmin":          userCtx.isAdmin == "ALL",
		"settingsRequired": s.Apps[appIdx].Status == StatusSettingsRequired,
		"infoMsg":          s.Apps[appIdx].InfoMessage,
		"currentStore":     storeName,
		"stores":           *stores,
	})
}

func (s *Server) UpdateSettingsHandler(c *gin.Context) {

	var appIdx int
	for i, a := range s.Apps {
		if a.AppId == c.Param("appId") && a.AccountId == c.PostForm("accountId") {
			appIdx = i
		}
	}
	s.Apps[appIdx].InfoMessage = c.PostForm("infoMessage")
	s.Apps[appIdx].Store = &store{c.PostForm("store")}
	c.Redirect(http.StatusFound, "/iframe/:appId")
}
