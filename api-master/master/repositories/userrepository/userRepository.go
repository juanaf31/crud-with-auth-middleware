package userrepository

import "liveCodeAPI/api-master/master/models"

type UserRepository interface {
	Get(*models.User) (bool, error)
}
