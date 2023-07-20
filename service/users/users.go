package users

import (
	"archive/dbconection"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Id        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Users []User
)

func (u *User) Rest() echo.Map {

	resp := echo.Map{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
	return resp
}

func (u *Users) Rest() []echo.Map {
	resp := make([]echo.Map, 0)

	for _, r := range *u {
		resp = append(resp, r.Rest())
	}

	return resp
}

func LoadUserByName(name string) (*User, error) {

	user := new(User)
	if err := dbconection.DB.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func SaveUser(id int64, userRegister *User) (*User, error) {

	if id == 0 {
		dbconection.DB.Create(userRegister)

		return userRegister, nil
	}

	dbconection.DB.Save(userRegister)

	return userRegister, nil
}

func LoadUserById(id int64) (*User, error) {

	user := new(User)
	if err := dbconection.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil

}

func DeleteUser(id int64) error {

	user := new(User)
	if err := dbconection.DB.Where("id = ?", id).First(user).Error; err != nil {
		return errors.New("user not found")
	}

	dbconection.DB.Delete(user)

	return nil
}
