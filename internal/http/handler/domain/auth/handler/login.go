package handler

import (
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/response"
	"be_entry_task/internal/modules/user"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Login for Login user
type Login struct {
	UserSrv user.UserService
}

func NewLogin() *Login {
	return &Login{}
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

	usr, err := l.UserSrv.Login(logReq)
	if err != nil {
		response.Err(w, err)
		return
	}

	response.Json(w, http.StatusOK, "Success", usr)

}
