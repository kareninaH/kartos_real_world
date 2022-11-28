package data

import (
	"real_world/internal/conf"

	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo, NewProfileRepo,
	NewArticleRepo, NewCommentRepo, NewTagRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	rd *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rd *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rd: rd}, cleanup, nil
}

func NewDB(c *conf.Data, logger log.Logger) *gorm.DB {
	//dsn := "root:realworld@tcp(127.0.0.1:3306)/real_world?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.NewHelper(logger).Info("数据连接失败!")
		panic(err)
	}

	initSchema(db)

	return db
}

func initSchema(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &UserFollow{}, &Article{}, &ArticleFavorited{}, &ArticleTag{},
		&Comment{}, &Tag{}); err != nil {
		panic(err)
	}
}

func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
