package handler

import (
	"be_entry_task/internal/http/handler/domain/user"
	"be_entry_task/internal/http/response"
	auth2 "be_entry_task/internal/modules/auth"
	user2 "be_entry_task/internal/modules/user"
	"be_entry_task/internal/redis"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
	"time"
)

func Auth(n httprouter.Handle, db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)
		authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]

		ctx := r.Context()

		meta, err := ValidateToken(ctx, authToken, db)
		if err != nil {

			response.Err(w, err)
			return
		}

		ctx = context.WithValue(ctx, "meta", meta)
		r = r.WithContext(ctx)

		n(w, r, ps)
	}
}

func ValidateToken(ctx context.Context, usrToken string, db *sql.DB) (user.AuthMeta, error) {
	re := redis.NewRedis()
	//check if token
	redisKey := "User-Auth:" + usrToken
	var usrMeta user.AuthMeta
	u, _ := re.GetBytes(ctx, redisKey)
	err := json.Unmarshal(u, &usrMeta)

	if usrMeta.ID != 0 {
		return usrMeta, nil
	}

	tokenObj, err := auth2.NewAuthRepo(db).SearchWithToken(usrToken)
	if err != nil {
		log.Print(err)
		return user.AuthMeta{}, err
	}

	if tokenObj.ID == 0 {
		return user.AuthMeta{}, errors.New("user not authorize")
	}

	const timeLayout = "2006-01-02 15:04:05"
	layout := "2006-01-02T15:04:05Z"

	expiryTime, _ := time.Parse(layout, tokenObj.ExpiredAt.String())
	currentTime, _ := time.Parse(layout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return user.AuthMeta{}, errors.New("The token is expired.\r\n")
	}

	usr, err := user2.NewUserRepository(db).Find(tokenObj.UserID)

	if err != nil {
		return user.AuthMeta{}, err
	}

	tokenDetails := user.AuthMeta{
		ID:             usr.ID,
		Username:       usr.Username,
		Email:          usr.Email,
		Nickname:       usr.Nickname,
		ProfilePicture: usr.ProfilePicture,
		CreatedAt:      usr.CreatedAt,
		UpdatedAt:      usr.UpdatedAt,
	}

	b, err := json.Marshal(&tokenDetails)
	if err != nil {
		return user.AuthMeta{}, err
	}

	err = re.Set(ctx, redisKey, b, 5*time.Minute)
	if err != nil {
		return user.AuthMeta{}, err
	}
	return tokenDetails, nil

}
