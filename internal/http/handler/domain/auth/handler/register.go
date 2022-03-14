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

// Register for registering user
type Register struct {
	UserSrv user.UserService
	db      *sql.DB
	dbRedis redis.RedisDB
}

func NewRegister(mysql *sql.DB, redis redis.RedisDB) *Register {
	return &Register{
		db:      mysql,
		dbRedis: redis,
	}
}
func (re *Register) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//validate request
	decoder := json.NewDecoder(r.Body)

	var req auth.RegisterUserRequest

	err := decoder.Decode(&req)

	if err != nil {
		response.Err(w, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		response.Err(w, err)
		return
	}

	err = user.NewUserService(re.db, user.NewUserRepository(re.db), auth2.NewAuthRepo(re.db), re.dbRedis).RegisterUser(req)
	if err != nil {
		response.Err(w, err)
		return
	}
	response.Json(w, http.StatusOK, "Success", req)
}
