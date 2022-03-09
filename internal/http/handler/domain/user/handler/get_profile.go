package handler

import (
	usrDom "be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	usrMod "be_entry_task/internal/modules/user"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

// GetProfile for GetProfile user
type GetProfile struct {
	UserSrv usrMod.UserService
}

func NewGetProfile() *GetProfile {
	return &GetProfile{}
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

	res, err := gp.UserSrv.GetProfile(usrMod.User{Username: authMeta.Username, Email: authMeta.Email})
	if err != nil {
		response.Err(w, err)
		return
	}

	if res.ProfilePicture == "" {
		res.ProfilePicture = os.Getenv("NO_IMG_URL")
	} else {
		fmt.Sprintf("https://storage.cloud.google.com/%s/%s", os.Getenv("BUCKET_NAME"), res.ProfilePicture)
	}

	profile := usrDom.Profile{
		Username:       authMeta.Username,
		Email:          authMeta.Email,
		Nickname:       res.Nickname,
		ProfilePicture: res.ProfilePicture,
	}
	response.Json(w, http.StatusOK, "Success", profile)
}
