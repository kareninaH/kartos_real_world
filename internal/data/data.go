package data

import (
	"real_world/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo,
	NewArticleRepo, NewCommentRepo, NewTagRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	//dsn := "root:realworld@tcp(127.0.0.1:3306)/real_world?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Info("数据连接失败!")
		panic(err)
	}

	if err = db.AutoMigrate(); err != nil {
		panic(err)
	}

	return db
}
