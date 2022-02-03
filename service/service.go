package service

import (
	"conference/model"
	"conference/storage"
)

//go:generate mockgen -source service.go -destination ./mock/service.go -package mock IService
type IService interface {
	Login(req model.UserLoginReq) (*model.UserAuthResponse, error)
	RegisterUser(req model.UserSignUpReq) (*model.UserAuthResponse, error)
	LogOut(refresh string) error
}

type Service struct {
	userRepo       storage.IUserRepository
	conferenceRepo storage.IConferenceRepository
}
