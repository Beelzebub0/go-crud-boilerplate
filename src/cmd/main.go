package main

import (
	"log"

	"github.com/Beelzebub0/go-crud-boilerplate/src/business/domain"
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/usecase"
	config "github.com/Beelzebub0/go-crud-boilerplate/src/conf"
	restserver "github.com/Beelzebub0/go-crud-boilerplate/src/handler/rest"
	"github.com/Beelzebub0/go-crud-boilerplate/src/lib/database"
)

var (
	dom domain.Domain
	uc  usecase.Usecase
	db  database.SQL

	// Configuration
	conf       config.Config
	serverConf config.ServerConfig
	sqlConf    config.SQLConfig
	redisConf  config.RedisConfig
)

func init() {

	log.SetFlags(log.Llongfile | log.Ltime)

	// Configuration
	conf = conf.GetConfig()
	serverConf = conf.Server
	sqlConf = conf.SQL
	redisConf = conf.Redis

	// Library
	db = database.InitSQL(sqlConf)

	// Infrastructure
	dom = domain.Init(db)
	uc = usecase.Init(dom)
}

func main() {
	restserver.Init(uc, serverConf)
}
