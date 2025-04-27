package config

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	env "github.com/lpernett/godotenv"
	"github.com/meldtechnology/merchant-inventory-go/pkg/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

const (
	defaultServerPort         = 8080
	defaultJWTExpirationHours = 72
)

// Config represents an application configuration.
type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN  string `yaml:"dsn" env:"DSN,secret"` // the data source name (DSN) for connecting to the database. required.
	MODE string `yaml:"mode" env:"ENV_MODE"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DSN, validation.Required),
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{
		ServerPort:    defaultServerPort,
		JWTExpiration: defaultJWTExpirationHours,
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.Load(".env"); err != nil {
		return nil, err
	}

	logger.Infof("App will run in %s mode on port %s", c.MODE, strconv.Itoa(c.ServerPort))

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}

// TODO: implement in revision 3
// Helpers
//func getEnv(key, fallback string) string {
//	if v, ok := os.LookupEnv(key); ok {
//		return v
//	}
//	return fallback
//}
//
//func getEnvAsBool(key string, fallback bool) bool {
//	val := getEnv(key, "")
//	if b, err := strconv.ParseBool(val); err == nil {
//		return b
//	}
//	return fallback
//}
//
//func getEnvAsInt(key string, fallback int) int {
//	val := getEnv(key, "")
//	if b, err := strconv.ParseInt(val, 10, 0); err == nil {
//		return int(b)
//	}
//	return fallback
//}
