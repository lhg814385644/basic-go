package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail error = errors.New("邮箱冲突")
	ErrUserNotFound       error = errors.New("用户不存在")
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Create 创建用户
func (ud *UserDao) Create(ctx context.Context, user Users) error {
	err := ud.db.WithContext(ctx).Create(&user).Error
	if err == nil {
		return nil
	}
	if me, ok := err.(*mysql.MySQLError); ok {
		const uniqueErrno uint16 = 1062
		if me.Number == uniqueErrno {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

// FindByEmail 根据邮箱查询用户
func (ud *UserDao) FindByEmail(ctx context.Context, email string) (*Users, error) {
	var user Users
	err := ud.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err == nil {
		return &user, nil
	}
	if err == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	}
	return nil, err
}

func (ud *UserDao) FindByID(ctx context.Context, id int) (*Users, error) {
	var user Users
	err := ud.db.WithContext(ctx).First(&user, id).Error
	if err == nil {
		return &user, nil
	}
	if err == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	}
	return nil, err
}

// Update 更新用户
func (ud *UserDao) Update(ctx context.Context, updateFiled map[string]interface{}, id int) error {
	var user Users
	if err := ud.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserNotFound
		}
		return err
	}
	return ud.db.WithContext(ctx).Model(&user).Updates(updateFiled).Error
}

// Users 用户表
type Users struct {
	ID           int    `gorm:"primaryKey,autoIncrement"` // ID主键，自增（TODO:面试常考 自增主键的好处？？ // ）
	Email        string `gorm:"unique"`                   // 邮箱，唯一索引
	Password     string
	NickName     string `gorm:"type:varchar(50)"`  // 昵称
	Birthday     string `gorm:"type:varchar(20)"`  // 生日
	Introduction string `gorm:"type:varchar(200)"` // 简介
	CTime        int64  // 创建时间
	UTime        int64  // 更新时间

}

// InitTable 初始化表(采用GORM的AutoMigrate)
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&Users{})
}
