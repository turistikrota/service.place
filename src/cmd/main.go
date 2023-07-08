package main

import (
	"context"

	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.shared/auth/session"
	"github.com/turistikrota/service.shared/auth/token"
	"github.com/turistikrota/service.shared/db/mongo"
	"github.com/turistikrota/service.shared/db/redis"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/env"
	"github.com/mixarchitecture/microp/events/nats"
	"github.com/mixarchitecture/microp/logs"
	"github.com/mixarchitecture/mredis"
	"github.com/turistikrota/service.place/src/config"
	"github.com/turistikrota/service.place/src/delivery"
	"github.com/turistikrota/service.place/src/service"
)

func main() {
	logs.Init()
	ctx := context.Background()
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	valid := validator.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadMongo(cnf)
	cache := mredis.New(&mredis.Config{
		Host:     cnf.CacheRedis.Host,
		Port:     cnf.CacheRedis.Port,
		Password: cnf.CacheRedis.Pw,
		DB:       cnf.CacheRedis.Db,
	})
	app := service.NewApplication(service.Config{
		App:         cnf,
		EventEngine: eventEngine,
		Mongo:       mongo,
		Validator:   valid,
		CacheSrv:    cache,
	})
	r := redis.New(&redis.Config{
		Host:     cnf.Redis.Host,
		Port:     cnf.Redis.Port,
		Password: cnf.Redis.Pw,
		DB:       cnf.Redis.Db,
	})
	tknSrv := token.New(token.Config{
		Expiration:     cnf.TokenSrv.Expiration,
		PublicKeyFile:  cnf.RSA.PublicKeyFile,
		PrivateKeyFile: cnf.RSA.PrivateKeyFile,
	})
	session := session.NewSessionApp(session.Config{
		Redis:       r,
		EventEngine: eventEngine,
		TokenSrv:    tknSrv,
		Topic:       cnf.Session.Topic,
		Project:     cnf.TokenSrv.Project,
	})
	del := delivery.New(delivery.Config{
		App:         app,
		Config:      cnf,
		I18n:        i18n,
		Validator:   valid,
		Ctx:         ctx,
		EventEngine: eventEngine,
		SessionSrv:  session.Service,
		TokenSrv:    tknSrv,
	})
	del.Load()
}

func loadMongo(cnf config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  cnf.DB.Place.Host,
		Port:  cnf.DB.Place.Port,
		User:  cnf.DB.Place.Username,
		Pass:  cnf.DB.Place.Password,
		Db:    cnf.DB.Place.Database,
		Query: cnf.DB.Place.Query,
	})
	d, err := mongo.New(uri, cnf.DB.Place.Database)
	if err != nil {
		panic(err)
	}
	return d
}
