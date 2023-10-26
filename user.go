package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users = make([]*User, 11)

func init() {
	for i := 1; i <= 10; i++ {
		users[i] = &User{
			FirstName: fmt.Sprint("FirstName", i),
			LastName:  fmt.Sprint("LastName", i),
			Age:       i,
		}
	}
}

type User struct {
	FirstName string
	LastName  string
	Age       int
}

type UserServiceInterface interface {
	Register(*User) error
	Update(*User) error
	Delete(int) error
	Find(int) (*User, error)
	List() ([]*User, error)
}

type userService struct{}

func NewUserService() UserServiceInterface {
	return &userService{}
}

func (u *userService) Register(user *User) error {
	// register process here
	fmt.Println("Registring user:", user)
	return nil
}

func (u *userService) Update(user *User) error {
	if user.FirstName == "" {
		return errors.New("First name is empty")
	}
	fmt.Println("Updating User")
	return nil
}

func (u *userService) Delete(id int) error {
	fmt.Println("Deleting user id:", id)
	return nil
}


func (u *userService) Find(id int) (*User, error) {
	return users[id], nil
}

func (u *userService) List() ([]*User, error) {
	return users, nil
}

func findUserHandler(service UserServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var err error
		var user *User
		user, err = service.Find(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func listUsersHandler(service UserServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := service.List()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}

func registerUser(service UserServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		service.Register(user)
		return c.JSON(http.StatusOK, user)
	}
}

type Message struct {
	Message string
}

func updateUser(service UserServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		if err := service.Update(user); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on creating user"})
		}
		return c.JSON(http.StatusOK, user)
	}
}

func deleteHandler(service UserServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		service.Delete(id)
		return c.JSON(http.StatusOK, Message{Message: "User deleted successfully"})
	}
}

func UserRoutes(server *echo.Echo) {
	userService := NewUserService()
	server.POST("/user", registerUser(userService))
	server.PUT("/user", updateUser(userService))
	server.DELETE("/user/:id", deleteHandler(userService))
	server.GET("/user/:id", findUserHandler(userService))
	server.GET("/user", listUsersHandler(userService))
}
