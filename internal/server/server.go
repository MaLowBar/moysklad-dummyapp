package server

import (
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/dummy-sloudel.sorochinsky"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

type Server struct {
	Apps []app.App
}

func NewServer() *Server {
	s := new(Server)

	files, err := os.ReadDir("internal/app")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		var a dummy.Dummy

		if strings.HasSuffix(file.Name(), ".app") {
			log.Println(file.Name())
			data, err := os.ReadFile("internal/app/" + file.Name())
			if err != nil {
				log.Println(err)
				break
			}
			if err = yaml.Unmarshal(data, &a.BaseApp); err != nil {
				log.Println(err)
				break
			}

			if a.Status != app.StatusInactive {
				s.Apps = append(s.Apps, &a)
			}
		}
	}

	return s
}

func (s *Server) addApp(appId, accountId string, otherInfo ActivateRequest) (app.App, error) {
	for _, a := range s.Apps {
		if a.GetBaseApp().AppId == appId && a.GetBaseApp().AccountId == accountId {
			return a, nil
		}
	}
	a, err := loadApp(appId, accountId, otherInfo)
	if err != nil {
		return nil, err
	}

	s.Apps = append(s.Apps, a)
	return a, nil
}

func loadApp(appId, accountId string, otherInfo ActivateRequest) (app.App, error) {
	ba, err := app.LoadBaseApp(appId, otherInfo.AppUID, accountId, otherInfo.Access[0].AccessToken)
	if err != nil {
		return &app.BaseApp{}, err
	}
	switch ba.AppUid {
	case "dummy-sloudel.sorochinsky":
		da := dummy.Dummy{BaseApp: *ba}
		return &da, nil
	default:
		return ba, nil
	}
}
