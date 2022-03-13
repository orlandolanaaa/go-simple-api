package server

import (
	middleware "be_entry_task/internal/http/handler"
	"be_entry_task/internal/http/handler/domain/auth/handler"
	user "be_entry_task/internal/http/handler/domain/user/handler"
	"be_entry_task/internal/redis"
	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	"context"
	"database/sql"
	errors "errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Server struct {
	srv             *http.Server
	Ctx             context.Context
	Storage         *cloud.Client
	FireStoreClient *firestore.Client
	db              *sql.DB
	redis           redis.RedisDB
}

func Get(mysql *sql.DB) *Server {
	return &Server{
		srv:   &http.Server{},
		db:    mysql,
		redis: redis.NewRedis(),
	}
}

func (s *Server) WithAddr(addr string) *Server {
	s.srv.Addr = addr
	return s
}

func (s *Server) WithErrLogger(l *log.Logger) *Server {
	s.srv.ErrorLog = l
	return s
}

func (s *Server) WithRouter() *Server {
	router := httprouter.New()

	//AUTH
	router.POST("/register", handler.NewRegister(s.db).Handle)
	router.POST("/login", handler.NewLogin(s.db).Handle)

	//FEATURE
	router.PUT("/users/profile", middleware.Auth(user.NewUpdateUser(s.db).Handle, s.db, s.redis))
	router.PUT("/users/profile-picture", middleware.Auth(user.NewUpdatePicture(s.db).Handle, s.db, s.redis))
	router.GET("/users/profile", middleware.Auth(user.NewGetProfile(s.db).Handle, s.db, s.redis))

	s.srv.Handler = router
	return s
}

func (s *Server) Start() error {
	if len(s.srv.Addr) == 0 {
		return errors.New("server missing address")
	}

	if s.srv.Handler == nil {
		return errors.New("server missing handler")
	}

	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}
