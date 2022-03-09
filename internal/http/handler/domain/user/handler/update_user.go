package handler

import (
	"be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	usrMod "be_entry_task/internal/modules/user"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
)

// UpdateUser for UploadPicture user
type UpdateUser struct {
	UserSrv usrMod.UserService
}

func NewUpdateUser() *UpdateUser {
	return &UpdateUser{}
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

	if userMeta.ProfilePicture == "" {
		userMeta.ProfilePicture = os.Getenv("NO_IMG_URL")
	} else {
		userMeta.ProfilePicture = fmt.Sprintf("https://storage.cloud.google.com/%s/%s", os.Getenv("BUCKET_NAME"), userMeta.ProfilePicture)
	}

	res, err := up.UserSrv.UploadProfile(user.User{
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
	fmt.Println(res.ProfilePicture)
	
	if res.ProfilePicture == "" {
		res.ProfilePicture = os.Getenv("NO_IMG_URL")
	} else {
		fmt.Sprintf("https://storage.cloud.google.com/%s/%s", os.Getenv("BUCKET_NAME"), res.ProfilePicture)
	}

	result := user.Profile{
		Username:       res.Username,
		Email:          res.Email,
		Nickname:       res.Nickname,
		ProfilePicture: fmt.Sprintf("https://storage.cloud.google.com/%s/%s", os.Getenv("BUCKET_NAME"), res.ProfilePicture),
	}

	response.Json(w, http.StatusOK, "Success", result)
}
