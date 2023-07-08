package adapters

import (
	"github.com/turistikrota/service.place/src/adapters/memory"
	"github.com/turistikrota/service.place/src/adapters/mongo"
	"github.com/turistikrota/service.place/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
