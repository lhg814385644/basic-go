package repository

import (
	"context"
	"errors"
	"github.com/lhg814385644/basic-go/webook/internal/domain"
	"github.com/lhg814385644/basic-go/webook/internal/repository/dao"
	"time"
)

var (
	ErrUserDuplicateEmail error = errors.New("邮箱冲突")
	ErrUserNotFound       error = errors.New("用户不存在")
)

type UserRepo struct {
	userDao *dao.UserDao
}

func NewUserRepo(userDao *dao.UserDao) *UserRepo {
	return &UserRepo{
		userDao: userDao,
	}
}

// Create 创建
func (up *UserRepo) Create(ctx context.Context, user domain.User) error {
	err := up.userDao.Create(ctx, dao.Users{
		Email:    user.Email,
		Password: user.Password,
		CTime:    time.Now().UnixMilli(),
		UTime:    time.Now().UnixMilli(),
	})
	if err == dao.ErrUserDuplicateEmail {
		return ErrUserDuplicateEmail
	}
	return err
}

func (up *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := up.userDao.FindByEmail(ctx, email)
	if err == nil {
		return &domain.User{
			ID:       user.ID,
			Email:    user.Email,
			Password: user.Password,
		}, nil
	}
	if err == dao.ErrUserNotFound {
		return nil, ErrUserNotFound
	}
	return nil, err
}

func (up *UserRepo) FindByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := up.userDao.FindByID(ctx, id)
	if err == nil {
		return &domain.User{
			ID:           user.ID,
			Email:        user.Email,
			Password:     user.Password,
			Birthday:     user.Birthday,
			NickName:     user.NickName,
			Introduction: user.Introduction,
		}, nil
	}
	if err == dao.ErrUserNotFound {
		return nil, ErrUserNotFound
	}
	return nil, err
}

func (up *UserRepo) Update(ctx context.Context, updateFiled map[string]interface{}, id int) error {
	if err := up.userDao.Update(ctx, updateFiled, id); err != nil {
		if err == dao.ErrUserNotFound {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
