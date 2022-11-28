package model

type Users struct {
	Id       int64  `gorm:"column:id;type:BIGINT;AUTO_INCREMENT;NOT NULL"`
	Email    string `gorm:"column:email;type:VARCHAR(255);NOT NULL"`
	Username string `gorm:"column:username;type:VARCHAR(255);NOT NULL"`
	Password string `gorm:"column:password;type:VARCHAR(255);NOT NULL"`
	Token    string `gorm:"column:token;type:VARCHAR(255);NOT NULL"`
	Bio      string `gorm:"column:bio;type:VARCHAR(255);"`
	Image    string `gorm:"column:image;type:VARCHAR(255);"`
}
