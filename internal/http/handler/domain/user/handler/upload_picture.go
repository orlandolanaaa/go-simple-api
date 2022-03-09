package handler

import (
	"be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	usrMod "be_entry_task/internal/modules/user"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// UpdatePicture for UploadPicture user
type UpdatePicture struct {
	UserSrv usrMod.UserService
}

func NewUpdatePicture() *UpdatePicture {
	return &UpdatePicture{}
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

	res, err := up.UserSrv.UploadPicture(r.Context(), file, handler, user.User{
		ID:             userMeta.ID,
		Username:       userMeta.Username,
		Email:          userMeta.Username,
		Nickname:       userMeta.Nickname,
		ProfilePicture: userMeta.ProfilePicture,
	})
	if err != nil {
		response.Err(w, err)
		return
	}

	response.Json(w, http.StatusOK, "Success", res)
}
