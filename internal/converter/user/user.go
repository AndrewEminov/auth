package user

import (
	"github.com/AndrewEminov/auth/internal/model"
	desc "github.com/AndrewEminov/auth/pkg/user_v1"
)

func ToInfo(info *desc.UserInfo) *model.User {
	return &model.User{
		Username: info.GetUsername(),
		Password: info.GetPassword(),
		Email:    info.GetEmail(),
		RoleId:   info.GetRoleId(),
	}
}
