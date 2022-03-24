package handler

import (
	"be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	auth2 "be_entry_task/internal/modules/auth"
	usrMod "be_entry_task/internal/modules/user"
	"be_entry_task/internal/redis"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
)

// UpdateUser for UploadPicture user
type UpdateUser struct {
	UserSrv usrMod.UserService
	db      *sql.DB
	dbRedis redis.RedisDB
}

func NewUpdateUser(mysql *sql.DB, redis redis.RedisDB) *UpdateUser {
	return &UpdateUser{db: mysql,
		dbRedis: redis}
}

func (up *UpdateUser) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := json.NewDecoder(r.Body)

	var usr user.User

	err := req.Decode(&usr)

	if err != nil {
		response.Err(w, err)
		return
	}

	validate := validator.New()

	err = validate.Struct(usr)
	if err != nil {
		response.Err(w, err)
		return
	}

	//check if username or email exists
	//get meta data from middleware
	meta := r.Context().Value("meta")
	b, _ := json.Marshal(meta)

	var userMeta user.AuthMeta
	err = json.Unmarshal(b, &userMeta)
	if err != nil {
		response.Err(w, err)
		return
	}

	res, err := usrMod.NewUserService(up.db, usrMod.NewUserRepository(up.db), auth2.NewAuthRepo(up.db), up.dbRedis).UpdateProfile(user.User{
		ID:             userMeta.ID,
		Username:       userMeta.Username,
		Email:          userMeta.Email,
		Nickname:       usr.Nickname,
		ProfilePicture: userMeta.ProfilePicture,
		CreatedAt:      userMeta.CreatedAt,
		UpdatedAt:      userMeta.UpdatedAt,
	})

	if err != nil {
		response.Err(w, err)
		return
	}

	var picUrl string
	if *res.ProfilePicture == "" {
		picUrl = os.Getenv("NO_IMG_URL")
		res.ProfilePicture = &picUrl
	} else {
		picUrl = os.Getenv("STORAGE_URL") + *res.ProfilePicture + os.Getenv("STORAGE_MEDIA")
		res.ProfilePicture = &picUrl
	}

	result := user.Profile{
		ID:             userMeta.ID,
		Username:       res.Username,
		Email:          res.Email,
		Nickname:       res.Nickname,
		ProfilePicture: res.ProfilePicture,
	}

	response.Json(w, http.StatusOK, "Success", result)
}
