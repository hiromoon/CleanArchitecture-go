package controllers

import (
	"strconv"

	"github.com/hiromoon/CleanArchitecture/domain"
	"github.com/hiromoon/CleanArchitecture/interfaces/databases"
	"github.com/hiromoon/CleanArchitecture/usecases"
)

// UserController is ...
type UserController struct {
	Interactor usecases.UserInteractor
}

// NewUserController is ...
func NewUserController(sqlHandler databases.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecases.UserInteractor{
			UserRepository: &databases.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Create is ...
func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, u)
}

// Index is ...
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

// Show is ...
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}

// Error is ...
type Error struct {
	Message string
}

// NewError is ...
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
