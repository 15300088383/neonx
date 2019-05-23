package redis

import (
	"strings"

	"github.com/go-redis/redis"

	"github.com/geekymedic/neonx"
	"github.com/geekymedic/neonx/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	redisList = map[string]*redis.Client{}
)

func init() {

	type RedisOptions struct {
		DSN      string
		Password string
		DB       int
	}

	neonx.AddPlugin("redis", func(status neonx.PluginStatus, viper *viper.Viper) error {
		switch status {
		case neonx.PluginLoad:

			var (
				dsnList = make(map[string]*RedisOptions)
			)

			err := viper.UnmarshalKey("redis", &dsnList)

			if err != nil {
				return errors.By(err)
			}

			if len(dsnList) == 0 {
				return errors.Format("redis plugin used, but redis config not exists.")
			}

			for name, opt := range dsnList {

				// 建立连接池
				clinet := redis.NewClient(&redis.Options{
					Addr:     opt.DSN,
					Password: opt.Password, // no password set
					DB:       opt.DB,       // use default DB
				})

				redisList[name] = clinet
			}

		}

		return nil

	})
}

func Use(name string) *redis.Client {
	return redisList[strings.ToLower(name)]
}
