package adapters

import (
	"api.turistikrota.com/place/src/adapters/memory"
	"api.turistikrota.com/place/src/adapters/mongo"
	"api.turistikrota.com/place/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
