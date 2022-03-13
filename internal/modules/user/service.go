package user

import (
	"be_entry_task/internal/firebase"
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/handler/domain/user"
	auth2 "be_entry_task/internal/modules/auth"
	"be_entry_task/internal/redis"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
	redis redis.RedisDB
	db    *sql.DB
}

func NewUserService(mysql *sql.DB) *UserService {
	return &UserService{
		redis: redis.NewRedis(),
		db:    mysql,
	}
}

//RegisterUser is business logic to register user
func (re *UserService) RegisterUser(req auth.RegisterUserRequest) error {
	//check if username or email exists
	userEx, err := NewUserRepository(re.db).SearchWithUsernameOrEmail(User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	if len(userEx) > 0 {
		return errors.New("user exists")
	}

	err = NewUserRepository(re.db).Create(User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (re *UserService) Login(usr auth.LoginRequest) (auth2.UserToken, error) {

	//check if username or email exists
	userEx, err := NewUserRepository(re.db).SearchWithUsernameOrEmail(User{Username: usr.Username})
	if err != nil {
		return auth2.UserToken{}, err
	}

	if len(userEx) == 0 {
		return auth2.UserToken{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEx[0].Password), []byte(usr.Password))

	if err != nil {
		return auth2.UserToken{}, err
	}

	randomToken := make([]byte, 32)

	_, err = rand.Read(randomToken)

	if err != nil {
		return auth2.UserToken{}, err
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	const timeLayout = "2006-01-02 15:04:05"

	dt := time.Now()
	//generatedAt := dt.Format(timeLayout)
	expireTime := time.Now().Add(time.Minute * 60)
	//expiresAt := expireTime.Format(timeLayout)

	var userTokenEn auth2.UserToken
	userTokenEn.Token = authToken
	userTokenEn.UserID = userEx[0].ID
	userTokenEn.ExpiredAt = &expireTime
	userTokenEn.CreatedAt = &dt

	id, err := auth2.NewAuthRepo(re.db).Create(userTokenEn)

	if err != nil {
		return auth2.UserToken{}, err
	}

	userTokenEn.ID = id

	return userTokenEn, err
}

//GetProfile is business logic to get profile user
func (re *UserService) GetProfile(usr User) (user.User, error) {
	//check if username or email exists
	userEx, err := NewUserRepository(re.db).SearchWithUsernameOrEmail(usr)
	if err != nil {
		return user.User{}, err
	}

	if len(userEx) == 0 {
		return user.User{}, err
	}

	return user.User{
		ID:             userEx[0].ID,
		Username:       userEx[0].Username,
		Email:          userEx[0].Email,
		Password:       userEx[0].Password,
		Nickname:       userEx[0].Nickname,
		ProfilePicture: userEx[0].ProfilePicture,
		CreatedAt:      userEx[0].CreatedAt,
		UpdatedAt:      userEx[0].UpdatedAt,
	}, nil
}

//UpdateProfile is business logic to upload profile user
func (re *UserService) UpdateProfile(usr user.User) (user.User, error) {
	//check if username or email exists
	userEx, err := NewUserRepository(re.db).Find(usr.ID)
	if err != nil {
		return user.User{}, err
	}

	if userEx.ID == 0 {
		return user.User{}, err
	}

	err = NewUserRepository(re.db).Update(User{
		ID:             usr.ID,
		Nickname:       usr.Nickname,
		ProfilePicture: usr.ProfilePicture,
	})

	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

//UploadPicture is business logic to upload picture user
func (re *UserService) UploadPicture(ctx context.Context, file multipart.File, handler *multipart.FileHeader, usr user.User) (user.User, error) {
	//check if username or email exists
	userEx, err := NewUserRepository(re.db).Find(usr.ID)

	if err != nil {
		return user.User{}, err
	}

	if userEx.ID == 0 {
		return user.User{}, err
	}

	//setup & upload image

	fileName := strings.Join(strings.Fields(handler.Filename+strconv.FormatInt(usr.ID, 10)), "")

	bucketName := os.Getenv("BUCKET_NAME") //ToDo: Replace with your bucket url
	fmt.Println(bucketName, fileName)

	fb := firebase.Firebase{}

	fb.NewService(ctx)

	writer := fb.Storage.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	defer writer.Close()

	byteSize, err := io.Copy(writer, file)
	if err != nil {
		return user.User{}, err
	}

	fmt.Printf("File size uploaded: %v\n", byteSize)

	err = NewUserRepository(re.db).Update(User{
		ID:             usr.ID,
		Nickname:       usr.Nickname,
		ProfilePicture: &fileName,
	})

	if err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:             userEx.ID,
		Username:       userEx.Username,
		Email:          userEx.Email,
		Nickname:       userEx.Nickname,
		ProfilePicture: &fileName,
		CreatedAt:      userEx.CreatedAt,
		UpdatedAt:      userEx.UpdatedAt,
	}, nil
}
