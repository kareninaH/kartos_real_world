package model

type Profiles struct {
	Id        int64  `gorm:"column:id;type:BIGINT;AUTO_INCREMENT;NOT NULL"`
	UserId    int64  `gorm:"column:userId;type:BIGINT;NOT NULL"`
	Username  string `gorm:"column:username;type:VARCHAR(255);NOT NULL"`
	Bio       string `gorm:"column:bio;type:VARCHAR(255);"`
	Image     string `gorm:"column:image;type:VARCHAR(255);"`
	Following int32  `gorm:"column:following;type:INT;"`
}
