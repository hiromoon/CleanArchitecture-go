package usecases

import "github.com/hiromoon/CleanArchitecture/domain"

// UserInteractor is ...
type UserInteractor struct {
	UserRepository UserRepository
}

// Add is ...
func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err = interactor.UserRepository.Store(u)
	return err
}

// Users is ...
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

// UserByID is ...
func (interactor *UserInteractor) UserByID(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(identifier)
	return
}
