package service

import (
	"context"
	"errors"
	"github.com/lhg814385644/basic-go/webook/internal/domain"
	"github.com/lhg814385644/basic-go/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail    error = errors.New("邮箱冲突")
	ErrInvalidUserOrPassword error = errors.New("无效邮箱或密码")
	ErrUserNotFound          error = errors.New("用户不存在")
)

type UserService struct {
	// TODO: repository仓储对象
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// SignUp 注册
func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	err = svc.userRepo.Create(ctx, user)
	if err == repository.ErrUserDuplicateEmail {
		return ErrUserDuplicateEmail
	}
	return err
}

// SignIn 登录
func (svc *UserService) SignIn(ctx context.Context, email, password string) (domain.User, error) {
	user, err := svc.userRepo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return *user, nil
}

// Profile 个人信息
func (svc *UserService) Profile(ctx context.Context, id int) (domain.User, error) {
	user, err := svc.userRepo.FindByID(ctx, id)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}
