package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode  indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	ServiceHost string
	ServicePort string

	Environment string // debug, test, release
	Version     string
	SecretKey   string

	AuthServiceHost string
	AuthGRPCPort    string

	BookServiceHost string
	BookGRPCPort    string

	DefaultOffset string
	DefaultLimit  string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("ErrEnvNodFound")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "api_gateways"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "3.123.20.220"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":7000"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_HOST", "3.123.20.220"))
	config.AuthGRPCPort = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_PORT", ":9102"))

	config.BookServiceHost = cast.ToString(getOrReturnDefaultValue("BOOK_SERVICE_HOST", "3.123.20.220"))
	config.BookGRPCPort = cast.ToString(getOrReturnDefaultValue("BOOK_SERVICE_PORT", ":9101"))

	config.SecretKey = "hello"

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "3.123.20.220"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "abdurahmon"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "aus1003"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.ServiceName))

	// config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
