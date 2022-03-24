package handler

import (
	usrDom "be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	auth2 "be_entry_task/internal/modules/auth"
	"be_entry_task/internal/modules/entities"
	usrMod "be_entry_task/internal/modules/user"
	"be_entry_task/internal/redis"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

// GetProfile for GetProfile user
type GetProfile struct {
	UserSrv usrMod.UserService
	db      *sql.DB
	dbRedis redis.RedisDB
}

func NewGetProfile(mysql *sql.DB, redis redis.RedisDB) *GetProfile {
	return &GetProfile{
		db:      mysql,
		dbRedis: redis,
	}
}

func (gp *GetProfile) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	meta := r.Context().Value("meta")
	b, _ := json.Marshal(meta)

	var authMeta usrDom.AuthMeta
	err := json.Unmarshal(b, &authMeta)
	if err != nil {
		response.Err(w, err)
		return
	}

	res, err := usrMod.NewUserService(gp.db, usrMod.NewUserRepository(gp.db), auth2.NewAuthRepo(gp.db), gp.dbRedis).GetProfile(entities.User{Username: authMeta.Username, Email: authMeta.Email})
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

	profile := usrDom.Profile{
		ID:             res.ID,
		Username:       res.Username,
		Email:          res.Email,
		Nickname:       res.Nickname,
		ProfilePicture: res.ProfilePicture,
	}

	response.Json(w, http.StatusOK, "Success", profile)
}
