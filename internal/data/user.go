package data

import (
	"context"
	"real_world/internal/biz"

	myerror "real_world/pkg/error"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string  `gorm:"column:email;type:VARCHAR(255);NOT NULL;uniqueIndex"`
	Username string  `gorm:"column:username;type:VARCHAR(255);NOT NULL;uniqueIndex"`
	Password string  `gorm:"column:password;type:VARCHAR(255);NOT NULL"`
	Bio      *string `gorm:"column:bio;type:VARCHAR(255);"`
	Image    *string `gorm:"column:image;type:VARCHAR(255);"`
}

//type Profiles struct {
//	gorm.Model
//	UserId    int64  `gorm:"column:userId;type:BIGINT;NOT NULL"`
//	Username  string `gorm:"column:username;type:VARCHAR(255);NOT NULL"`
//	Bio       string `gorm:"column:bio;type:VARCHAR(255);"`
//	Image     string `gorm:"column:image;type:VARCHAR(255);"`
//	Following int32  `gorm:"column:following;type:INT;default:0"`
//}

type UserFollow struct {
	gorm.Model
	UserName       string `gorm:"index:idx_username_followusername;type:varchar(255)"`
	FollowUserName string `gorm:"index:idx_username_followusername;type:varchar(255)"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r userRepo) Create(ctx context.Context, user *biz.User) error {
	u := bizToDataUser(user)
	result := r.data.db.Create(&u)
	log.Info("新用户Id:", u.ID, "插入条数:", result.RowsAffected)
	if result.Error != nil {
		return myerror.HttpBadRequest("user", "already exist")
	}
	return nil
}

func (r userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	var u User
	result := r.data.db.Where(&User{Email: email}).First(&u)
	if result.Error != nil {
		return nil, myerror.HttpBadRequest("user", "don't exist")
	}
	return dataToBiz(&u), nil
}

func (r userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	var u User
	result := r.data.db.Where(&User{Username: username}).First(&u)
	if result.Error != nil {
		return nil, myerror.HttpBadRequest("user", "don't exist")
	}
	return dataToBiz(&u), nil
}

func (r userRepo) Update(ctx context.Context, user *biz.User) error {
	u := bizToDataUser(user)
	var us User
	r.data.db.Select("id").Where(&User{Email: u.Email}).First(&us)

	result := r.data.db.Model(&us).Updates(u)
	if result.Error != nil {
		return myerror.HttpBadRequest("user", "update fail")
	}
	return nil
}

func dataToBiz(u *User) *biz.User {
	return &biz.User{
		Email:          u.Email,
		Username:       u.Username,
		Bio:            *u.Bio,
		Image:          *u.Image,
		PasswordHashed: u.Password,
	}
}

func bizToDataUser(u *biz.User) *User {
	return &User{
		Email:    u.Email,
		Username: u.Username,
		Password: u.PasswordHashed,
		Bio:      &u.Bio,
		Image:    &u.Image,
	}
}
