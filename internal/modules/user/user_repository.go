package user

import (
	"be_entry_task/internal/modules/entities"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserRepository defines the datastore handling persisting User records.
type UserRepository interface {
	Create(entities.User) error
	Find(int64) (entities.User, error)
	SearchWithUsernameOrEmail(entities.User) ([]entities.User, error)
	Update(entities.User) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(mysql *sql.DB) UserRepository {
	return &UserRepo{db: mysql}
}

func (u *UserRepo) Create(user entities.User) error {

	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	fmt.Println(string(hashPwd))

	_, err := u.db.Exec("INSERT into users (username,password,email) values (?,?,?)", user.Username, string(hashPwd), user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("insert success!")
	return nil
}

func (u *UserRepo) Find(id int64) (entities.User, error) {

	rows := u.db.QueryRow("SELECT * from users where id = ? ", id)
	var res entities.User
	err := rows.Scan(&res.ID, &res.Username, &res.Password, &res.Email, &res.Nickname, &res.ProfilePicture, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return entities.User{}, err
	}
	return res, nil
}

func (u *UserRepo) Update(user entities.User) error {

	const timeLayout = "2006-01-02 15:04:05"
	dt := time.Now()
	updatedAt := dt.Format(timeLayout)

	_, err := u.db.Exec("UPDATE db_entry_task.users SET nickname = ?, profile_picture = ?, updated_at = ? WHERE id = ?", user.Nickname, user.ProfilePicture, updatedAt, user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("update success!")
	return nil
}

func (u *UserRepo) SearchWithUsernameOrEmail(user entities.User) ([]entities.User, error) {

	rows, err := u.db.Query("SELECT * from users where username = ? or email = ?", user.Username, user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return []entities.User{}, err
	}

	defer rows.Close()

	var result []entities.User

	for rows.Next() {
		var each = entities.User{}
		var err = rows.Scan(&each.ID, &each.Username, &each.Password, &each.Email, &each.Nickname, &each.ProfilePicture, &each.CreatedAt, &each.UpdatedAt, &each.DeletedAt)

		if err != nil {
			fmt.Println(err.Error())
			return []entities.User{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return []entities.User{}, err
	}
	return result, nil
}
