package usecases

import "github.com/hiromoon/CleanArchitecture/domain"

// UserRepository is ...
type UserRepository interface {
	Store(domain.User) (int, error)
	FindByID(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
