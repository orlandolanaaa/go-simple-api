package handler

import (
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/response"
	"be_entry_task/internal/modules/user"
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
}

func NewRegister(mysql *sql.DB) *Register {
	return &Register{
		db: mysql,
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

	err = user.NewUserService(re.db).RegisterUser(req)
	if err != nil {
		response.Err(w, err)
		return
	}
	response.Json(w, http.StatusOK, "Success", req)
}
