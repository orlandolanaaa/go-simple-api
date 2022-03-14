package handler

import (
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/response"
	auth2 "be_entry_task/internal/modules/auth"
	"be_entry_task/internal/modules/user"
	"be_entry_task/internal/redis"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Login for Login user
type Login struct {
	UserSrv user.UserService
	db      *sql.DB
	dbRedis redis.RedisDB
}

func NewLogin(mysql *sql.DB, redis redis.RedisDB) *Login {
	return &Login{
		db:      mysql,
		dbRedis: redis,
	}
}

func (l *Login) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := json.NewDecoder(r.Body)

	var logReq auth.LoginRequest

	err := req.Decode(&logReq)

	if err != nil {
		response.Err(w, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(logReq)
	if err != nil {
		response.Err(w, err)
		return
	}

	usr, err := user.NewUserService(l.db, user.NewUserRepository(l.db), auth2.NewAuthRepo(l.db), l.dbRedis).Login(logReq)
	if err != nil {
		response.Err(w, err)
		return
	}

	response.Json(w, http.StatusOK, "Success", usr)

}
