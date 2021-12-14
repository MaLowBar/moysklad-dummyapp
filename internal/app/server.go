package app

import (
	"crypto/rand"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-dummyapp-moysklad/internal/config"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Server struct {
	Apps   []*App
	Config *config.Config
}

func request(method, url, bearerToken string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header = map[string][]string{
		"Authorization": {"Bearer" + bearerToken},
		"Content-Type":  {"application/json"},
	}
	return req
}

func (s *Server) vendorRequest(method, url string, body io.Reader) *http.Request {
	jti := make([]byte, 32)
	rand.Read(jti)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": s.Config.AppUid,
		"iat": time.Now().Unix(),
		"jti": jti,
	}).SignedString([]byte(s.Config.SecretKey))
	if err != nil {
		log.Printf("Creating JWT error: %s", err)
		return nil
	}
	return request(method, url, token, body)
}

func NewServer(cfg *config.Config) *Server {
	s := new(Server)
	s.Config = cfg
	files, err := os.ReadDir("internal/app")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		a := new(App)
		log.Println(file.Name())
		if strings.HasSuffix(file.Name(), ".app") {
			data, err := os.ReadFile("internal/app/" + file.Name())
			if err != nil {
				log.Println(err)
				break
			}
			if err = json.Unmarshal(data, a); err != nil {
				log.Println(err)
				break
			}
			if a.Status != StatusInactive {
				s.Apps = append(s.Apps, a)
			}
		}
	}
	return s
}

func (s *Server) addApp(appId, accountId, accessToken string) *App {
	for _, a := range s.Apps {
		if a.AppId == appId && a.AccountId == accountId {
			return a
		}
	}
	s.Apps = append(s.Apps, loadApp(appId, accountId, accessToken))
	return s.Apps[len(s.Apps)-1]
}

type userContext struct {
	uid       string
	fio       string
	accountId string
	isAdmin   string
}

func (s *Server) getContext(contextKey string) *userContext {
	client := &http.Client{Timeout: 10 * time.Second}

	url := "http://localhost:8080/ctx" //s.Config.MoyskladVendorApiEndpointUrl + "/contextKey/" + contextKey
	req := s.vendorRequest("POST", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var rawCtx map[string]interface{}
	if resp.Status == "200 OK" {
		if err = json.NewDecoder(resp.Body).Decode(&rawCtx); err != nil {
			log.Println(err)
			return nil
		}
		ctx := new(userContext)
		ctx.uid = rawCtx["uid"].(string)
		ctx.fio = rawCtx["shortFio"].(string)
		ctx.accountId = rawCtx["accountId"].(string)
		ctx.isAdmin = rawCtx["permissions"].(map[string]interface{})["admin"].(map[string]interface{})["view"].(string)
		return ctx
	}
	return nil
}

func (s *Server) ActivateHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if c.Param("appId") != s.Config.AppId {
		c.Status(http.StatusBadRequest)
		return
	}

	var req struct {
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

	reqBody := c.Request.Body
	err := json.NewDecoder(reqBody).Decode(&req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	currentApp := s.addApp(s.Config.AppId, c.Param("accountId"), req.Access[0].AccessToken)

	if currentApp == nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": currentApp.Status})
}

func (s *Server) DeleteHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	for i, a := range s.Apps {
		if a.AppId == c.Param("appId") && a.AccountId == c.Param("accountId") {
			s.Apps = append(s.Apps[:i], s.Apps[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func (s *Server) StatusHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	for _, a := range s.Apps {
		if a.AppId == c.Param("appId") && a.AccountId == c.Param("accountId") {
			c.JSON(http.StatusOK, gin.H{"status": a.Status})
			return
		}
	}
	c.Status(http.StatusNotFound)
}
