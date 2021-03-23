package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/alcalbg/gotdd/middleware"
	"github.com/alcalbg/gotdd/render"
	"github.com/alcalbg/gotdd/session"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	SID          string
	Email        string
	PasswordHash string
}

type UserRepository interface {
	GetUserByEmail(email string) (User, error)
}

type Server struct {
	Router         *mux.Router
	session        *session.Session
	userRepository UserRepository
}

func NewServer(logger *log.Logger, s *session.Session, userRepository UserRepository) *Server {
	srv := &Server{}
	srv.session = s
	srv.userRepository = userRepository

	srv.Router = mux.NewRouter()
	srv.Router.NotFoundHandler = srv.notFound()

	srv.Router.Handle("/", srv.home()).Methods(http.MethodGet)
	srv.Router.Handle("/login", srv.login()).Methods(http.MethodGet)
	srv.Router.Handle("/login", srv.loginSubmit()).Methods(http.MethodPost)
	srv.Router.Handle("/register", srv.login()).Methods(http.MethodGet, http.MethodPost)

	//bad := func() http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		var x map[string]int
	//		x["y"] = 1 // will produce nil map panic
	//	})
	//}
	//s.Router.Handle("/bad", bad())

	srv.Router.Use(middleware.Logger(logger))
	srv.Router.Use(middleware.AuthRedirector(srv.session))

	return srv
}

func (srv Server) home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (srv Server) login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := render.NewTemplate("login.html")
		t.Render(w, r, http.StatusOK)
	})
}

func (srv Server) loginSubmit() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
		password := r.FormValue("password")

		user, err := srv.userRepository.GetUserByEmail(email)
		if err != nil {
			// TODO
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			// TODO - login failed
			return
		}

		// user is ok, save to session
		srv.session.SetUserSID(w, r, user.SID)

		http.Redirect(w, r, "/", http.StatusFound)
		return
	})
}

func (srv Server) notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
}