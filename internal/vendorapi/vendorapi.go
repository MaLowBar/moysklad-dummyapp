package vendorapi

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-dummyapp-moysklad/internal/jsonapi"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const baseUrl = "https://online.moysklad.ru/api/vendor/1.0"

type UserContext struct {
	Uid       string
	Fio       string
	AccountId string
	IsAdmin   string
}

func binToHex(bs []byte) string {
	var s string
	for _, b := range bs {
		s += fmt.Sprintf("%x", b)
	}
	return s
}

func vendorRequest(method, url, appUid string, body io.Reader) *http.Request {
	jti := make([]byte, 32)
	rand.Read(jti)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:  appUid,
		IssuedAt: time.Now().UTC().Unix(),
		Id:       binToHex(jti),
	})
	token.Header = gin.H{"alg": "HS256"}
	secretKey, _ := os.ReadFile("internal/" + appUid + "/secret.key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Creating JWT error: %s", err)
		return nil
	}
	return jsonapi.Request(method, url, tokenString, body)
}

func GetContext(appUid, contextKey string) (*UserContext, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	url := baseUrl + "/context/" + contextKey //"http://localhost:8080/ctx"
	req := vendorRequest("POST", url, appUid, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rawCtx map[string]interface{}
	if resp.Status == "200 OK" {
		if err = json.NewDecoder(resp.Body).Decode(&rawCtx); err != nil {
			return nil, err
		}
		ctx := new(UserContext)
		ctx.Uid = rawCtx["uid"].(string)
		ctx.Fio = rawCtx["shortFio"].(string)
		ctx.AccountId = rawCtx["accountId"].(string)
		ctx.IsAdmin = rawCtx["permissions"].(map[string]interface{})["admin"].(map[string]interface{})["view"].(string)
		return ctx, nil
	}
	return nil, errors.New(resp.Status)
}
