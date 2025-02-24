package config

import "time"

//TODO:: Mark all of them in viper config
const (
	BindPort    = "8081"
	AlgoAddress = "https://testnet-algorand.api.purestake.io/ps2"
	PsTokenKey  = "X-API-Key"
	PsToken     = "qAMLbrOhmT9ewbvFUkUwD8kOOJ6ifFCz1boJoXyb"
	//TODO::Place it in a separate const directory temporay placed here
	KnownTLSError = "Post \"https://testnet-algorand.api.purestake.io/ps2/v2/transactions\": net/http: TLS handshake timeout"
	//db
	DbFolder   = "data"
	DbFileName = "sqlite.db"
	//cron job
	CronJobSpec = "@every 1m"
	//REDIS config
	RedisUrl      = "redis://localhost:6379/0"
	RedisPrefix   = "ms-multisig"
	RedisMaxRetry = 3
	//Limiter rate
	// Simplified format "<limit>-<period>"
	// * limit: number
	// * period:
	// * "S": second
	// * "M": minute
	// * "H": hour
	// * "D": day
	//
	// Examples:
	//
	// * 5 reqs/second: "5-S"
	// * 10 reqs/minute: "10-M"
	// * 1000 reqs/hour: "1000-H"
	// * 2000 reqs/day: "2000-D"
	LimiterRate = "1000-H"
	CsrfSecret  = "32-byte-long-csrf-secret"
	//JWT
	JwtRealm         = "ms-multisign"
	JwtSecret        = "Use a super secret key here"
	JwtTokenLookup   = "header: Authorization"
	JwtTokenHeadName = "Bearer"
	JwtTimeout       = 12 * time.Hour
	JwtMaxRefresh    = 12 * time.Hour
	//Swagger
	SwaggerHost = "localhost:8081/"
	//CORS
	CorsAllowOrigin = "http://localhost:3000"
)
