package userusecase

import "liveCodeAPI/api-master/master/models"

type UserUsecase interface {
	GetUser(*models.User) (bool, error)
}
