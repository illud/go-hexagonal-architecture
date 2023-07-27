package repositories

import (
	userModel "hexagonal-architecture/app/users/domain/models"
)

type IUser interface {
	Create(user *userModel.User) string
	Get() []*userModel.User
}
