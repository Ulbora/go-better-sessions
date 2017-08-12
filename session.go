package services

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var sessionKey string

//Session a sesstion
type Session struct {
	store      *sessions.CookieStore
	MaxAge     int
	Name       string
	SessionKey string
}

//InitSessionStore initialize session store
func (s *Session) InitSessionStore(res http.ResponseWriter, req *http.Request) {
	if s.store == nil {
		s.createSessionStore(res, req)
	}
	if s.MaxAge == 0 {
		fmt.Println("using defalut max age")
		s.MaxAge = 3600 //default 3600 seconds --  1 hour
	}
	if s.Name == "" {
		fmt.Println("using defalut name")
		s.Name = "user-session"
	}
	if s.SessionKey == "" {
		fmt.Println("using defalut sesstion key")
		s.SessionKey = "554dfgdffdd11dfgf1ff1f" // default key
	}

}

// CreateSessionStore creates a sesstion
func (s *Session) createSessionStore(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating Session Store")
	s.store = sessions.NewCookieStore([]byte(s.SessionKey))
	s.store.Options = &sessions.Options{
		MaxAge:   s.MaxAge,
		HttpOnly: true,
	}
}

//GetSession get session
func (s *Session) GetSession(req *http.Request) (*sessions.Session, error) {
	return s.store.Get(req, s.Name)
	//return session, err
}
