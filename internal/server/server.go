package server

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/dummy-sloudel.sorochinsky"
	"log"
)

type Server struct {
	Apps []app.App
	db *sql.DB
}

func NewServer() (*Server, error) {
	s := new(Server)
	db, err := sql.Open("pgx", "postgres://user:pswd@host/db")
	if err != nil {
		return nil, err
	}
	s.db = db

	rows, err := db.Query("SELECT * FROM baseapp")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			a app.BaseApp
			aId int
		)
		err = rows.Scan(&aId, &a.AppId, &a.AppUid, &a.AccountId, &a.Status, &a.AccessToken, &a.SecretKey)
		if err != nil {
			return nil, err
		}

		switch a.AppUid {
		case "dummy-sloudel.sorochinsky":
			da := dummy.Dummy{BaseApp: a}

			row := db.QueryRow("SELECT info_message, store FROM dummyapp WHERE baseapp_id = $1", aId)
			err = row.Scan(&da.InfoMessage, &da.Store)
			if err != nil {
				log.Println("here")
				return nil, err
			}
			if da.Status != app.StatusInactive {
				s.Apps = append(s.Apps, &da)
			}
		default:
			return nil, errors.New("wrong AppUid: " + a.AppUid)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//files, err := os.ReadDir("internal/app")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//for _, file := range files {
	//	var a dummy.Dummy
	//
	//	if strings.HasSuffix(file.Name(), ".app") {
	//		log.Println(file.Name())
	//		data, err := os.ReadFile("internal/app/" + file.Name())
	//		if err != nil {
	//			log.Println(err)
	//			break
	//		}
	//		if err = yaml.Unmarshal(data, &a.BaseApp); err != nil {
	//			log.Println(err)
	//			break
	//		}
	//
	//		if a.Status != app.StatusInactive {
	//			s.Apps = append(s.Apps, &a)
	//		}
	//	}
	//}

	return s, nil
}

func (s *Server) addApp(appId, accountId string, otherInfo ActivateRequest) (app.App, error) {
	for _, a := range s.Apps {
		if a.GetBaseApp().AppId == appId && a.GetBaseApp().AccountId == accountId {
			return a, nil
		}
	}
	a, err := s.loadApp(appId, accountId, otherInfo)
	if err != nil {
		return nil, err
	}

	s.Apps = append(s.Apps, a)
	return a, nil
}

func (s *Server) loadApp(appId, accountId string, otherInfo ActivateRequest) (app.App, error) {
	ba, err := app.LoadBaseApp(appId, otherInfo.AppUID, accountId, otherInfo.Access[0].AccessToken, s.db)
	if err != nil {
		return &app.BaseApp{}, err
	}
	switch ba.AppUid {
	case "dummy-sloudel.sorochinsky":
		da, err := dummy.LoadDummyApp(*ba, s.db)
		if err != nil {
			return nil, err
		}
		return da, nil
	default:
		return ba, nil
	}
}
