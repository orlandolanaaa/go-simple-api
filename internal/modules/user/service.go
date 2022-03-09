package user

import (
	"be_entry_task/internal/firebase"
	"be_entry_task/internal/http/handler/domain/auth"
	"be_entry_task/internal/http/handler/domain/user"
	auth2 "be_entry_task/internal/modules/auth"
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
	UserRepo UserRepo
	AuthRepo auth2.AuthRepo
}

//func NewUserService(userRepository repository.User) *UserService {
//	return &UserService{
//		UserRepo: userRepository,
//	}
//}
//

//RegisterUser is business logic to register user
func (re *UserService) RegisterUser(req auth.RegisterUserRequest) error {
	//check if username or email exists

	userEx, err := re.UserRepo.SearchWithUsernameOrEmailLogin(User{
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

	err = re.UserRepo.Create(User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

//GetProfile is business logic to get profile user
func (re *UserService) GetProfile(usr User) (user.User, error) {
	//check if username or email exists
	userEx, err := re.UserRepo.SearchWithUsernameOrEmailLogin(usr)
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
		Nickname:       userEx[0].Nickname.String,
		ProfilePicture: userEx[0].ProfilePicture.String,
		CreatedAt:      userEx[0].CreatedAt.Time.String(),
		UpdatedAt:      userEx[0].UpdatedAt.Time.String(),
	}, nil
}

func (re *UserService) Login(usr auth.LoginRequest) (auth2.UserToken, error) {

	//check if username or email exists
	userEx, err := re.UserRepo.SearchWithUsernameOrEmailLogin(User{Username: usr.Username})
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
	expiresAt := expireTime.Format(timeLayout)

	var userTokenEn auth2.UserToken
	userTokenEn.Token = authToken
	userTokenEn.UserID = userEx[0].ID
	userTokenEn.ExpiredAt = expiresAt
	userTokenEn.CreatedAt = struct {
		Time  time.Time
		Valid bool
	}{Time: dt, Valid: true}

	id, err := re.AuthRepo.Create(userTokenEn)

	if err != nil {
		return auth2.UserToken{}, err
	}

	userTokenEn.ID = id

	return userTokenEn, err
}

//UploadProfile is business logic to upload profile user
func (re *UserService) UploadProfile(usr user.User) (user.User, error) {
	//check if username or email exists
	userEx, err := re.UserRepo.Find(usr.ID)
	if err != nil {
		return user.User{}, err
	}

	if userEx.ID == 0 {
		return user.User{}, err
	}

	err = re.UserRepo.Update(User{
		ID: usr.ID,
		Nickname: sql.NullString{
			String: usr.Nickname,
			Valid:  true,
		},
		ProfilePicture: sql.NullString{
			String: usr.ProfilePicture,
			Valid:  true,
		},
	})

	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

//UploadPicture is business logic to upload picture user
func (re *UserService) UploadPicture(ctx context.Context, file multipart.File, handler *multipart.FileHeader, usr user.User) (user.User, error) {
	//check if username or email exists
	userEx, err := re.UserRepo.Find(usr.ID)

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

	err = re.UserRepo.Update(User{
		ID: usr.ID,
		Nickname: sql.NullString{
			String: usr.Nickname,
			Valid:  true,
		},
		ProfilePicture: sql.NullString{
			String: fileName,
			Valid:  true,
		},
	})

	uploadedImageUrl := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucketName, fileName)
	usr.ProfilePicture = uploadedImageUrl

	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}
