package data

import (
	"study-kratis/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger,db *gorm.DB,rdb. *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db,rdb: rdb}, cleanup, nil
}

func NewDB(c *conf.Data)*gorm.DB{
	newLogger := logger.New(
		slog.New(os.Srdout,"\r\n",slog.LstdFlags),
		logger.Config{
			SlowThreshold:time.Second,
			Colorful:true,
			LogLevel:logger.Info
		},
	)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
        Logger:                                   newLogger,
        DisableForeignKeyConstraintWhenMigrating: true,
        NamingStrategy:                           schema.NamingStrategy{
            //SingularTable: true, // 表名是否加 s
        },
    })

    if err != nil {
        log.Errorf("failed opening connection to sqlite: %v", err)
        panic("failed to connect database")
    }

    return db
}

func NewRedis(c *conf.Data) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:         c.Redis.Addr,
        Password:     c.Redis.Password,
        DB:           int(c.Redis.Db),
        DialTimeout:  c.Redis.DialTimeout.AsDuration(),
        WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
        ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
    })
    rdb.AddHook(redisotel.TracingHook{})
    if err := rdb.Close(); err != nil {
        log.Error(err)
    }
    return rdb
}
