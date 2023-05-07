package service

import (
	"context"
	"fmt"
	"kratos-gorm-git/models"
	"kratos-gorm-git/pkg/helper"

	pb "kratos-gorm-git/api/git"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	fmt.Println(req.Username, req.Password)
	ub := new(models.UserBasic)
	err := models.DB.Where("username = ? AND password = ?", req.Username, helper.GetMd5(req.Password)).Find(ub).Error
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(ub.Identity, ub.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: token,
	}, nil
}
