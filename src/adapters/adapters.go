package adapters

import (
	"api.turistikrota.com/account/src/adapters/memory"
	"api.turistikrota.com/account/src/adapters/mongo"
	"api.turistikrota.com/account/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
