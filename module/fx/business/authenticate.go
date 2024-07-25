package fx_business

import "golang-server/module/fx/repository"

type IAuthenticateBiz interface {
}

type authenticateBiz struct {
	userRepo repository.IUserRepository
}

func NewAuthenticateBiz(
	userRepo repository.IUserRepository,
) IAuthenticateBiz {
	return authenticateBiz{
		userRepo: userRepo,
	}
}
