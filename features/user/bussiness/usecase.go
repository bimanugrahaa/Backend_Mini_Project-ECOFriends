package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/user"

	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userData user.Data
	validate *validator.Validate
}

func NewUserBussiness(userData user.Data) user.Bussiness {
	return &userUsecase{
		userData: userData,
		validate: validator.New(),
	}
}

func (uu *userUsecase) CreateUser(data user.UserCore) (resp user.UserCore, err error) {
	return user.UserCore{}, nil
}

func (uu *userUsecase) GetAllUser() (resp []user.UserCore) {
	resp = uu.userData.SelectData()
	return
}
