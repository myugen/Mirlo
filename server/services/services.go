package services

import "github.com/alephshahor/Mirlo/server/repositories"

type IServices interface {
	User() IUserService
	SignUp() ISignUpService
}

type services struct {
	userService   IUserService
	signUpService ISignUpService
}

var servicesInstance *services

func Services() IServices {
	if servicesInstance == nil {
		servicesInstance = &services{}
	}
	return servicesInstance
}

func (s *services) User() IUserService {
	if s.userService == nil {
		s.userService = NewUserService(repositories.Repositories())
	}
	return s.userService
}

func (s *services) SignUp() ISignUpService {
	if s.signUpService == nil {
		s.signUpService = NewSignUpService(repositories.Repositories(), s)
	}
	return s.signUpService
}
