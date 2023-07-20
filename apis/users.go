package apis

import (
	"archive/service/users"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// POST /login: User login with username and password.
func Login(c echo.Context) error {

	var (
		err error
	)

	u := new(users.User)

	l := new(users.LoginForm)
	err = c.Bind(l)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while binding")})
	}

	if l.Username != "" && l.Password != "" {
		u, err = users.LoadUserByName(l.Username)
		if err != nil {

			return c.JSON(400, echo.Map{"error": errors.New("user not found or password is incorrect")})
		}
	}

	if !CheckPasswordHash(l.Password, u.Password) {

		return c.JSON(400, echo.Map{"error": errors.New("user not found or password is incorrect")})
	}

	token, err := CreateJWTToken(u.Id)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while creating JWT token")})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token, "massage": "logined"})
}

// POST /register: User registration with username, email, and password.
func Register(c echo.Context) error {

	var (
		err error
	)

	u := new(users.User)

	r := new(users.RegisterForm)
	err = c.Bind(r)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while binding")})
	}

	err = r.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("a problem occurred while validating")})
	}

	u.Name = r.Username
	u.Email = r.Email
	u.Password, _ = HashPassword(r.Password)
	u.CreatedAt = time.Now()

	u, err = users.SaveUser(0, u)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while saving user")})
	}

	return c.JSON(http.StatusOK, echo.Map{"user": u.Rest(), "message": "user created successfully"})

}

// GET /users/{id}: Get details of a specific user by their ID.
func GetUser(c echo.Context) error {

	var (
		err error
	)

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	u := new(users.User)

	u, err = users.LoadUserById(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("user not found: " + id)})
	}

	return c.JSON(http.StatusOK, echo.Map{"user": u.Rest(), "message": "loaded successfully"})
}

// PUT /users/{id}: Update the details of a specific user by their ID.
func UpdateUser(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(400, echo.Map{"error": errors.New("invalid id")})
	}

	newUser := new(users.RegisterForm)
	err = c.Bind(newUser)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while binding")})
	}

	err = newUser.Validate()
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("a problem occurred while validating")})
	}

	u := new(users.User)
	u, err = users.LoadUserById(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("user not found: " + id)})
	}

	u.Name = newUser.Username
	u.Email = newUser.Email
	u.Password, _ = HashPassword(newUser.Password)
	u.UpdatedAt = time.Now()

	u, err = users.SaveUser(intId, u)
	if err != nil {

		return c.JSON(500, echo.Map{"error": errors.New("a problem occurred while updating user")})
	}

	return c.JSON(http.StatusOK, echo.Map{"user": u.Rest(), "message": "user updated successfully"})
}

// DELETE /users/{id}: Delete a specific user by their ID.
func DeleteUser(c echo.Context) error {

	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.JSON(400, echo.Map{"error": errors.New("invalid id ")})
	}

	err = users.DeleteUser(intId)
	if err != nil {

		return c.JSON(400, echo.Map{"error": errors.New("a problem occurred while deleting user")})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "user deleted successfully"})
}
