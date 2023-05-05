package user_v1

import (
	"context"
	converter "github.com/AndrewEminov/auth/internal/converter/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	desc "github.com/AndrewEminov/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*emptypb.Empty, error) {
	log.Println("test")

	err := i.userService.Create(ctx, converter.ToInfo(req.GetInfo()))

	if err != nil {
		return nil, err
	}

	return nil, nil
}
