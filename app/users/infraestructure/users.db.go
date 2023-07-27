package infraestructure

import (
	userModel "hexagonal-architecture/app/users/domain/models"
	db "hexagonal-architecture/data"
)

type UserDb struct {
	// Add any dependencies or configurations related to the UserRepository here, if needed.
}

func NewUserDb() *UserDb {
	// Initialize any dependencies and configurations for the UserRepository here, if needed.
	return &UserDb{}
}

func (u *UserDb) Create(user *userModel.User) string {
	db.Client().Create(&userModel.User{Name: user.Name, Age: user.Age})

	return "User created"
}

func (u *UserDb) Get() []*userModel.User {
	var users []*userModel.User
	db.Client().Order("updated_at desc").Find(&users)

	return users
}
