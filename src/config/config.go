package config

type MongoPlace struct {
	Host       string `env:"MONGO_PLACE_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_PLACE_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_PLACE_USERNAME" envDefault:""`
	Password   string `env:"MONGO_PLACE_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_PLACE_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_PLACE_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_PLACE_QUERY" envDefault:""`
}

type MongoFeature struct {
	Collection string `env:"MONGO_FEATURE_COLLECTION" envDefault:"accounts"`
}

type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type Topics struct {
	Account  AccountEvents
	Platform PlatformEvents
}

type AccountEvents struct {
	Deleted       string `env:"STREAMING_TOPIC_ACCOUNT_DELETED"`
	Created       string `env:"STREAMING_TOPIC_ACCOUNT_CREATED"`
	Updated       string `env:"STREAMING_TOPIC_ACCOUNT_UPDATED"`
	Disabled      string `env:"STREAMING_TOPIC_ACCOUNT_DISABLED"`
	Enabled       string `env:"STREAMING_TOPIC_ACCOUNT_ENABLED"`
	SocialAdded   string `env:"STREAMING_TOPIC_ACCOUNT_SOCIAL_ADDED"`
	SocialUpdated string `env:"STREAMING_TOPIC_ACCOUNT_SOCIAL_UPDATED"`
	SocialRemoved string `env:"STREAMING_TOPIC_ACCOUNT_SOCIAL_REMOVED"`
}

type PlatformEvents struct {
	Created            string `env:"STREAMING_TOPIC_PLATFORM_CREATED"`
	Updated            string `env:"STREAMING_TOPIC_PLATFORM_UPDATED"`
	Deleted            string `env:"STREAMING_TOPIC_PLATFORM_DELETED"`
	Disabled           string `env:"STREAMING_TOPIC_PLATFORM_DISABLED"`
	Enabled            string `env:"STREAMING_TOPIC_PLATFORM_ENABLED"`
	TranslationCreated string `env:"STREAMING_TOPIC_PLATFORM_TRANSLATION_CREATED"`
	TranslationUpdated string `env:"STREAMING_TOPIC_PLATFORM_TRANSLATION_UPDATED"`
	TranslationRemoved string `env:"STREAMING_TOPIC_PLATFORM_TRANSLATION_REMOVED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type CDN struct {
	Url string `env:"CDN_URL" envDefault:"http://localhost:3000"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Feature MongoFeature
		Place   MongoPlace
	}
	RSA         RSA
	HttpHeaders HttpHeaders
	Server      Server
	Session     Session
	I18n        I18n
	Topics      Topics
	Nats        Nats
	Redis       Redis
	CacheRedis  CacheRedis
	TokenSrv    TokenSrv
	CDN         CDN
}
