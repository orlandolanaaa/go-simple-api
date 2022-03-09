package handler

import (
	"be_entry_task/internal/http/response"
	auth2 "be_entry_task/internal/modules/auth"
	user2 "be_entry_task/internal/modules/user"
	"context"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
	"time"
)

func Auth(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)
		authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

		meta, err := ValidateToken(authToken)
		if err != nil {

			response.Err(w, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "meta", meta)
		r = r.WithContext(ctx)

		n(w, r, ps)
	}
}

func ValidateToken(usrToken string) (map[string]interface{}, error) {

	//check if username or email exists
	repo := auth2.AuthRepo{}
	tokenObj, err := repo.SearchWithToken(usrToken)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if tokenObj.ID == 0 {
		return nil, errors.New("user not authorize")
	}

	const timeLayout = "2006-01-02 15:04:05"
	layout := "2006-01-02T15:04:05Z"
	expiryTime, _ := time.Parse(layout, tokenObj.ExpiredAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The token is expired.\r\n")
	}
	userREpo := user2.UserRepo{}
	usr, err := userREpo.Find(tokenObj.UserID)

	if err != nil {
		return nil, err
	}

	tokenDetails := map[string]interface{}{
		"id":              usr.ID,
		"username":        usr.Username,
		"email":           usr.Email,
		"nickname":        usr.Nickname.String,
		"profile_picture": usr.ProfilePicture.String,
		"token":           tokenObj.Token,
		"expires_at":      tokenObj.ExpiredAt,
	}

	return tokenDetails, nil

}
