package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"errors"

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

	email, _ := uu.userData.SelectUserEmail(data)

	if email.Email == data.Email {
		err = errors.New("data is available")
		return user.UserCore{}, err
	}

	resp, err = uu.userData.InsertUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUsecase) UpdateUser(data user.UserCore) (resp user.UserCore, err error) {

	userID, _ := uu.userData.SelectUserById(data.ID)

	resp, err = uu.userData.EditUser(data)
	resp.ID = userID.ID

	if err != nil {
		return user.UserCore{}, err
	}

	return resp, nil
}

func (uu *userUsecase) DeleteUser(id int) (err error) {

	err = uu.userData.RemoveUser(id)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetAllUser() (resp []user.UserCore) {
	resp = uu.userData.SelectAllUser()
	return
}

func (uu *userUsecase) GetUserById(id int) (resp user.UserCore, err error) {
	resp, err = uu.userData.SelectUserById(id)
	return
}

func (uu *userUsecase) Login(data user.UserCore) (resp user.UserCore, err error) {
	resp, err = uu.userData.Login(data)

	if err != nil {
		err = errors.New("something went wrong")
		return user.UserCore{}, err
	}

	return resp, nil
}
