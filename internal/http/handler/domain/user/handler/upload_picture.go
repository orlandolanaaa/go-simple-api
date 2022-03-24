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
	"net/http"
	"os"
)

// UpdatePicture for UploadPicture user
type UpdatePicture struct {
	UserSrv usrMod.UserService
	db      *sql.DB
	dbRedis redis.RedisDB
}

func NewUpdatePicture(mysql *sql.DB, redis redis.RedisDB) *UpdatePicture {
	return &UpdatePicture{
		db:      mysql,
		dbRedis: redis,
	}
}
func (up *UpdatePicture) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//check if username or email exists
	//get meta data from middleware
	meta := r.Context().Value("meta")
	b, _ := json.Marshal(meta)

	var userMeta user.AuthMeta
	err := json.Unmarshal(b, &userMeta)
	if err != nil {
		response.Err(w, err)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		response.Err(w, err)
		return
	}
	defer file.Close()

	res, err := usrMod.NewUserService(up.db, usrMod.NewUserRepository(up.db), auth2.NewAuthRepo(up.db), up.dbRedis).UploadPicture(r.Context(), file, handler, user.User{
		ID:             userMeta.ID,
		Username:       userMeta.Username,
		Email:          userMeta.Username,
		Nickname:       userMeta.Nickname,
		ProfilePicture: userMeta.ProfilePicture,
	})

	var picUrl string
	if *res.ProfilePicture == "" {
		picUrl = os.Getenv("NO_IMG_URL")
		res.ProfilePicture = &picUrl
	} else {
		picUrl = os.Getenv("STORAGE_URL") + *res.ProfilePicture + os.Getenv("STORAGE_MEDIA")
		res.ProfilePicture = &picUrl
	}

	if err != nil {
		response.Err(w, err)
		return
	}

	profile := user.Profile{
		ID:             res.ID,
		Username:       res.Username,
		Email:          res.Email,
		Nickname:       res.Nickname,
		ProfilePicture: res.ProfilePicture,
	}

	response.Json(w, http.StatusOK, "Success", profile)
}
