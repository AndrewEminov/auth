package user_v1

import (
	user "github.com/AndrewEminov/auth/internal/service/user"
	desc "github.com/AndrewEminov/auth/pkg/user_v1"
)

//import (
//	desc "github.com/AndrewEminov/auth/pkg/user_v1"
//)
//
//type Implementation struct {
//	desc.UnimplementedUserV1Server
//}
//
//func NewImplementation() *Implementation {
//	return &Implementation{}
//}

type Implementation struct {
	desc.UnimplementedUserV1Server

	userService user.Service
}

func NewImplementation(userService user.Service) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
