package userusecase

import (
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/userrepository"
)

type UserUsecaseImpl struct {
	userRepo userrepository.UserRepository
}

func InitUserUsecase(userRepo userrepository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUser(data *models.User) (bool, error) {
	isValid, err := u.userRepo.Get(data)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
