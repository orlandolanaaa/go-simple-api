package user

import (
	"be_entry_task/internal/mysql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//
//type UpdateUserReq struct {
//	ID             int64  `json:"id"  db:"id"`
//	Nickname       string `json:"nickname" db:"nickname"`
//	ProfilePicture string `json:"profile_picture" db:"profile_picture"`
//}

// UserRepository defines the datastore handling persisting User records.
type UserRepository interface {
	Create(User) error
	Find(int64) (User, error)
	SearchWithUsernameOrEmail(User) ([]User, error)
	Update(User) error
	SearchWithUsernameOrEmailLogin(User) ([]User, error)
}

type UserRepo struct {
}

//func NewUserRepository() UserRepository {
//	return &User{}
//}

func (u *UserRepo) Create(user User) error {
	db, err := mysql.Conn()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer db.Close()

	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	fmt.Println(string(hashPwd))

	_, err = db.Exec("INSERT into users (username,password,email) values (?,?,?)", user.Username, string(hashPwd), user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("insert success!")
	return nil
}

func (u *UserRepo) Find(id int64) (User, error) {
	db, err := mysql.Conn()
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}

	defer db.Close()
	rows := db.QueryRow("SELECT * from users where id = ? ", id)

	var res User
	err = rows.Scan(&res.ID, &res.Username, &res.Password, &res.Email, &res.Nickname, &res.ProfilePicture, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	return res, nil
}

func (u *UserRepo) Update(user User) error {

	db, err := mysql.Conn()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer db.Close()

	const timeLayout = "2006-01-02 15:04:05"
	dt := time.Now()
	updatedAt := dt.Format(timeLayout)

	_, err = db.Exec("UPDATE db_entry_task.users SET nickname = ?, profile_picture = ?, updated_at = ? WHERE id = ?", user.Nickname, user.ProfilePicture, updatedAt, user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("update success!")
	return nil
}

func (u *UserRepo) SearchWithUsernameOrEmail(user User) ([]User, error) {

	db, err := mysql.Conn()

	if err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT * from users where username = ? or email = ?", user.Username, user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}

	defer rows.Close()

	var result []User

	for rows.Next() {
		var each = User{}
		var err = rows.Scan(&each.ID, &each.Username, &each.Password, &each.Email, &each.Nickname, &each.ProfilePicture, &each.CreatedAt, &each.UpdatedAt, &each.DeletedAt)

		if err != nil {
			fmt.Println(err.Error())
			return []User{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}
	return result, nil
}

func (u *UserRepo) SearchWithUsernameOrEmailLogin(user User) ([]User, error) {

	db, err := mysql.Conn()

	if err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT * from users where username = ? or email = ?", user.Username, user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}

	defer rows.Close()

	var result []User

	for rows.Next() {
		var each = User{}
		var err = rows.Scan(&each.ID, &each.Username, &each.Password, &each.Email, &each.Nickname, &each.ProfilePicture, &each.CreatedAt, &each.UpdatedAt, &each.DeletedAt)

		if err != nil {
			fmt.Println(err.Error())
			return []User{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return []User{}, err
	}
	return result, nil
}
