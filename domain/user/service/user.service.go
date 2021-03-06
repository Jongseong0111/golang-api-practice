package service

import (
	"context"
	"errors"
	"tutorial.sqlc.dev/app/domain/user/dto"
	"tutorial.sqlc.dev/app/model"
)

type UserService struct{}

func (receiver UserService) CreateUser(req dto.User) (res model.ScmAccount, err error) {
	duplicateAccount, err := dao.CheckDuplicateAccount(context.Background(), req.UserAccount)
	if err != nil {
		return
	}

	if len(duplicateAccount) > 0 {
		err = errors.New("duplicate Account")
		return
	}

	duplicateEmail, err := dao.CheckDuplicateEmail(context.Background(), req.UserEmail)
	if err != nil {
		return
	}

	if len(duplicateEmail) > 0 {
		err = errors.New("duplicate Email")
		return
	}

	params := model.CreateUserParams{
		UserName:  req.UserName,
		UserEmail: req.UserEmail,
		UserAccount:  req.UserAccount,
		UserPassword: req.UserPassword,
	}

	result, err := dao.CreateUser(context.Background(), params)

	userID, err := result.LastInsertId()
	if err != nil {
		return
	}

	fetchedUser, err := dao.GetUser(context.Background(), int32(userID))
	if err != nil {
		return
	}

	return fetchedUser, err
}

func (receiver UserService) GetUserList() (res []model.ListUserRow, err error) {
	users, err := dao.ListUser(context.Background())
	if err != nil {
		return
	}

	return users, err
}

func (receiver UserService) UpdateUserAccount(userInfo model.UpdateUserParams) (err error) {
	err = dao.UpdateUser(context.Background(), userInfo)
	if err != nil {
		return
	}
	return
}

func (receiver UserService) GetSignedID(req string) (res int32, err error) {
	userId, err := dao.GetSignedId(context.Background(), req)
	if err != nil {
		return
	}
	return userId, err
}